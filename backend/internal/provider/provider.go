// Package provider defines the boundary between Skriptra and any model vendor.
//
// The application depends on these interfaces and nothing else. Which
// implementation is constructed is decided once, at startup, from
// configuration:
//
//	configuration -> dependency initialization -> interfaces -> application
//
// There is deliberately no `if production { hosted } else { local }` anywhere in
// this codebase. Moving from a local Ollama model to a hosted provider is an
// environment variable and a restart, with no change to business logic. That
// property is the reason this package exists.
package provider

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
)

// ---------------------------------------------------------------------------
// LLM
// ---------------------------------------------------------------------------

// Role identifies the author of a message in a generation request.
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// Message is one turn in a conversation.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// GenerateRequest is vendor-neutral on purpose: it carries only options that
// every supported backend can honour. Anything vendor-specific belongs in that
// vendor's adapter configuration, not here.
type GenerateRequest struct {
	Messages    []Message
	MaxTokens   int
	Temperature float32
	Stop        []string
}

// Usage is reported back for observability and cost accounting. Local providers
// report token counts but no cost, which is exactly the point of measuring it.
type Usage struct {
	PromptTokens     int    `json:"promptTokens"`
	CompletionTokens int    `json:"completionTokens"`
	Provider         string `json:"provider"`
	Model            string `json:"model"`
}

// GenerateResponse is a completed, non-streamed generation.
type GenerateResponse struct {
	Text  string `json:"text"`
	Usage Usage  `json:"usage"`
}

// Chunk is one streamed fragment. Exactly one of Text or Err is meaningful;
// Done marks the final chunk and is the only one carrying Usage.
type Chunk struct {
	Text  string
	Done  bool
	Usage Usage
	Err   error
}

// LLM generates text. Implementations must be safe for concurrent use.
type LLM interface {
	// Generate returns a complete response.
	Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error)

	// Stream returns a channel closed when generation ends. Implementations
	// must respect ctx cancellation and must always close the channel.
	Stream(ctx context.Context, req GenerateRequest) (<-chan Chunk, error)

	// Info describes what is actually running, for /api/v1/providers.
	Info() Info
}

// ---------------------------------------------------------------------------
// Embedder
// ---------------------------------------------------------------------------

// Embedder turns text into vectors.
//
// Dimensions() is part of the interface rather than a detail because changing
// the embedding model changes the vector dimension, which invalidates every
// stored vector in the database. The schema records the model and dimension per
// row so a mismatch fails loudly at startup instead of silently returning
// meaningless similarity scores.
type Embedder interface {
	// Embed returns one vector per input, in input order.
	Embed(ctx context.Context, texts []string) ([][]float32, error)

	// Dimensions is the fixed output width of this model.
	Dimensions() int

	Info() Info
}

// Info describes a configured provider for the /providers endpoint.
type Info struct {
	Provider   string `json:"provider"`
	Model      string `json:"model"`
	Local      bool   `json:"local"`
	Dimensions int    `json:"dimensions,omitempty"`
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

var (
	// ErrUnavailable means the provider could not be reached. It maps to
	// HTTP 503 so a stopped Ollama is distinguishable from a bug.
	ErrUnavailable = errors.New("provider unavailable")

	// ErrContextTooLong means the prompt exceeded the model's window. The
	// caller should retrieve fewer chunks rather than fail the request.
	ErrContextTooLong = errors.New("context too long")

	// ErrRateLimited means back off and retry.
	ErrRateLimited = errors.New("rate limited")
)

// ---------------------------------------------------------------------------
// Registry
// ---------------------------------------------------------------------------

// Factories are registered by adapter packages in their init(), so adding a
// provider means adding a file — never editing a switch statement in the
// application.
type (
	LLMFactory      func(cfg Settings) (LLM, error)
	EmbedderFactory func(cfg Settings) (Embedder, error)
)

// Settings is the resolved configuration for one provider instance.
type Settings struct {
	Provider   string
	Model      string
	BaseURL    string
	APIKey     string
	Dimensions int
	Extra      map[string]string
}

var (
	mu             sync.RWMutex
	llmFactories   = map[string]LLMFactory{}
	embedFactories = map[string]EmbedderFactory{}
)

// RegisterLLM registers an LLM adapter under a provider name.
func RegisterLLM(name string, f LLMFactory) {
	mu.Lock()
	defer mu.Unlock()
	llmFactories[name] = f
}

// RegisterEmbedder registers an embedding adapter under a provider name.
func RegisterEmbedder(name string, f EmbedderFactory) {
	mu.Lock()
	defer mu.Unlock()
	embedFactories[name] = f
}

// NewLLM constructs the configured LLM adapter.
func NewLLM(cfg Settings) (LLM, error) {
	mu.RLock()
	f, ok := llmFactories[cfg.Provider]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown LLM provider %q (available: %v)", cfg.Provider, LLMProviders())
	}
	return f(cfg)
}

// NewEmbedder constructs the configured embedding adapter.
func NewEmbedder(cfg Settings) (Embedder, error) {
	mu.RLock()
	f, ok := embedFactories[cfg.Provider]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown embedding provider %q (available: %v)", cfg.Provider, EmbedderProviders())
	}
	return f(cfg)
}

// LLMProviders lists registered LLM adapter names.
func LLMProviders() []string { return keys(llmFactories) }

// EmbedderProviders lists registered embedding adapter names.
func EmbedderProviders() []string { return keys(embedFactories) }

func keys[V any](m map[string]V) []string {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
