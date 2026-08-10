// Package config loads all runtime configuration from the environment.
//
// Environment-based, always. Locally the values come from .env.local; in
// production they come from real environment variables and secrets. The code
// path is identical, the difference between a laptop and a cluster is
// configuration and infrastructure, never a branch in the application.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

// Config is the fully resolved application configuration.
type Config struct {
	Env             string
	HTTPAddr        string
	ShutdownTimeout time.Duration

	DatabaseURL string
	NATSURL     string
	RedisURL    string
	ParserAddr  string

	StorageDir  string
	MaxUploadMB int

	LLM       provider.Settings
	Embedding provider.Settings

	OTelEndpoint string
	LogLevel     string
}

// Load reads configuration from the environment and validates it.
//
// Validation happens once, at startup, and fails loudly. A misconfigured
// provider should stop the process immediately, not surface as a confusing
// 500 on the first user question.
func Load() (*Config, error) {
	c := &Config{
		Env:             env("APP_ENV", "development"),
		HTTPAddr:        env("HTTP_ADDR", ":8080"),
		ShutdownTimeout: envDuration("SHUTDOWN_TIMEOUT", 15*time.Second),

		DatabaseURL: env("DATABASE_URL", ""),
		NATSURL:     env("NATS_URL", "nats://localhost:4222"),
		RedisURL:    env("REDIS_URL", "redis://localhost:6379"),
		ParserAddr:  env("PARSER_ADDR", "localhost:50051"),

		StorageDir:  env("STORAGE_DIR", "./data/uploads"),
		MaxUploadMB: envInt("MAX_UPLOAD_MB", 50),

		LLM: provider.Settings{
			Provider: env("LLM_PROVIDER", "ollama"),
			Model:    env("LLM_MODEL", "llama3.1:8b"),
			BaseURL:  env("LLM_BASE_URL", "http://localhost:11434"),
			APIKey:   env("LLM_API_KEY", ""),
		},
		Embedding: provider.Settings{
			Provider:   env("EMBEDDING_PROVIDER", "ollama"),
			Model:      env("EMBEDDING_MODEL", "bge-m3"),
			BaseURL:    env("EMBEDDING_BASE_URL", "http://localhost:11434"),
			APIKey:     env("EMBEDDING_API_KEY", ""),
			Dimensions: envInt("EMBEDDING_DIMENSIONS", 768),
		},

		OTelEndpoint: env("OTEL_EXPORTER_OTLP_ENDPOINT", ""),
		LogLevel:     env("LOG_LEVEL", "info"),
	}

	return c, c.validate()
}

func (c *Config) validate() error {
	var problems []string

	if c.DatabaseURL == "" {
		problems = append(problems, "DATABASE_URL is required")
	}
	if c.LLM.Provider == "" {
		problems = append(problems, "LLM_PROVIDER is required")
	}
	if c.Embedding.Provider == "" {
		problems = append(problems, "EMBEDDING_PROVIDER is required")
	}

	// The schema pins vector columns to a fixed width. Booting with a
	// mismatched embedding model would write vectors that are silently
	// incomparable to every existing row, so refuse to start instead.
	const schemaDimensions = 768
	if c.Embedding.Dimensions != schemaDimensions {
		problems = append(problems, fmt.Sprintf(
			"EMBEDDING_DIMENSIONS is %d but the schema expects %d, "+
				"changing the embedding model requires a migration and a re-embed",
			c.Embedding.Dimensions, schemaDimensions))
	}

	// Hosted providers need a key; local ones must not silently require one.
	if requiresAPIKey(c.LLM.Provider) && c.LLM.APIKey == "" {
		problems = append(problems, fmt.Sprintf("LLM_API_KEY is required for provider %q", c.LLM.Provider))
	}
	if requiresAPIKey(c.Embedding.Provider) && c.Embedding.APIKey == "" {
		problems = append(problems, fmt.Sprintf("EMBEDDING_API_KEY is required for provider %q", c.Embedding.Provider))
	}

	if len(problems) > 0 {
		return fmt.Errorf("invalid configuration:\n  - %s", strings.Join(problems, "\n  - "))
	}
	return nil
}

// requiresAPIKey reports whether a provider name denotes a hosted service.
// Local runtimes (Ollama, llama.cpp, a self-hosted TGI) need no credentials.
func requiresAPIKey(name string) bool {
	switch name {
	case "ollama", "llamacpp", "tgi", "local":
		return false
	default:
		return true
	}
}

// IsLocal reports whether inference stays on this machine, surfaced in the UI
// so a user can see at a glance that nothing left the device.
func (c *Config) IsLocal() bool { return !requiresAPIKey(c.LLM.Provider) }

func env(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func envDuration(key string, def time.Duration) time.Duration {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
