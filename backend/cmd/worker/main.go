// Command worker consumes ingestion jobs and runs the pipeline.
//
// Separate from the API process on purpose. Ingestion is CPU and model bound
// and takes tens of seconds per document; running it in the API would make
// request latency depend on how many people happen to be uploading. Split, they
// scale independently.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/skriptra/skriptra/backend/internal/cache"
	"github.com/skriptra/skriptra/backend/internal/config"
	"github.com/skriptra/skriptra/backend/internal/db"
	"github.com/skriptra/skriptra/backend/internal/ingest"
	"github.com/skriptra/skriptra/backend/internal/provider"
	"github.com/skriptra/skriptra/backend/internal/queue"

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

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	store, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer store.Close()

	embedder, err := provider.NewEmbedder(cfg.Embedding)
	if err != nil {
		return err
	}
	llm, err := provider.NewLLM(cfg.LLM)
	if err != nil {
		return err
	}

	cached, err := cache.Connect(ctx, cfg.RedisURL, log)
	if err != nil {
		return err
	}
	// Re-ingesting a document after a parser fix re-embeds mostly unchanged
	// questions, which is exactly what this saves.
	embedder = cache.NewEmbedder(embedder, cached)

	q, err := queue.Connect(ctx, cfg.NATSURL)
	if err != nil {
		return err
	}
	defer q.Close()

	pipeline := ingest.NewPipeline(store, embedder, llm, log)

	// Where a Python OCR sidecar would be registered, once one exists:
	//   pipeline.RegisterParser(grpcparser.New(cfg.ParserAddr))
	// Nothing else in this file, or in the pipeline, would change.

	log.Info("worker started", "nats", cfg.NATSURL, "storage", cfg.StorageDir)

	return q.Consume(ctx, func(ctx context.Context, job queue.IngestMessage) error {
		log.Info("ingesting", "document", job.DocumentID, "file", job.Filename)

		content, err := os.ReadFile(filepath.Join(cfg.StorageDir, job.StorageKey))
		if err != nil {
			_ = store.FailDocument(ctx, job.DocumentID, "reading stored file: "+err.Error())
			// Returning nil acknowledges the message: a missing file will not
			// appear on retry, so redelivering it only wastes capacity.
			return nil
		}

		if err := pipeline.Run(ctx, ingest.Job{
			DocumentID: job.DocumentID,
			CourseID:   job.CourseID,
			Filename:   job.Filename,
			Content:    content,
		}); err != nil {
			return err
		}

		// New questions change every aggregate for this course, so the derived
		// caches are dropped rather than left to expire. Stale analytics after
		// an upload would be a visible, confusing wrong answer.
		cached.DeletePrefix(ctx, cache.CourseScope(job.CourseID.String()))
		return nil
	})
}
