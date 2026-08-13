// Command api serves the Skriptra HTTP API.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/skriptra/skriptra/backend/internal/api"
	"github.com/skriptra/skriptra/backend/internal/auth"
	"github.com/skriptra/skriptra/backend/internal/cache"
	"github.com/skriptra/skriptra/backend/internal/config"
	"github.com/skriptra/skriptra/backend/internal/db"
	"github.com/skriptra/skriptra/backend/internal/provider"
	"github.com/skriptra/skriptra/backend/internal/queue"

	// Adapters register themselves in init(). Adding a provider means adding an
	// import here, never editing a switch statement in the application.
	_ "github.com/skriptra/skriptra/backend/internal/provider/ollama"
	_ "github.com/skriptra/skriptra/backend/internal/provider/openaicompat"
)

func main() {
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	level := slog.LevelInfo
	if cfg.LogLevel == "debug" {
		level = slog.LevelDebug
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(log)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	store, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer store.Close()

	// Fails here on a weak or missing signing secret rather than serving
	// forgeable tokens.
	issuer, err := auth.NewIssuer(cfg.JWTSecret)
	if err != nil {
		return err
	}

	// Construct providers once, at startup. A bad provider name or a missing
	// key should stop the process here, not surface as a confusing 500 on the
	// first question a user asks.
	llm, err := provider.NewLLM(cfg.LLM)
	if err != nil {
		return err
	}
	embedder, err := provider.NewEmbedder(cfg.Embedding)
	if err != nil {
		return err
	}

	// Optional. With REDIS_URL unset this is a no-op and the application
	// behaves identically, just without the savings.
	cached, err := cache.Connect(ctx, cfg.RedisURL, log)
	if err != nil {
		return err
	}
	embedder = cache.NewEmbedder(embedder, cached)

	q, err := queue.Connect(ctx, cfg.NATSURL)
	if err != nil {
		return err
	}
	defer q.Close()

	log.Info("providers configured",
		"llm", cfg.LLM.Provider, "llm_model", cfg.LLM.Model,
		"embedding", cfg.Embedding.Provider, "embedding_model", cfg.Embedding.Model,
		"local_inference", cfg.IsLocal(), "cache", cached.Enabled())

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           api.New(cfg, store, llm, embedder, q, cached, issuer, log).Routes(),
		ReadHeaderTimeout: 10 * time.Second,
		// No WriteTimeout: /ask streams, and a write deadline would cut a long
		// local generation off mid-answer. Cancellation is the request context.
		IdleTimeout: 60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Info("listening", "addr", cfg.HTTPAddr, "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Info("shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
