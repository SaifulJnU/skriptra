// Package ollama adapts a local Ollama server to the provider interfaces.
//
// Registered in init(), so enabling it is an import — the application never
// contains a switch over provider names.
package ollama

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

func init() {
	provider.RegisterLLM("ollama", func(s provider.Settings) (provider.LLM, error) {
		return &client{settings: s, http: newHTTPClient()}, nil
	})
	provider.RegisterEmbedder("ollama", func(s provider.Settings) (provider.Embedder, error) {
		return &client{settings: s, http: newHTTPClient()}, nil
	})
}

func newHTTPClient() *http.Client {
	// No overall timeout: generation legitimately runs for minutes on a local
	// model. Cancellation is the caller's context, which is the right lever.
	return &http.Client{
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
			MaxIdleConnsPerHost: 4,
		},
	}
}

type client struct {
	settings provider.Settings
	http     *http.Client
}

func (c *client) Info() provider.Info {
	return provider.Info{
		Provider:   "ollama",
		Model:      c.settings.Model,
		Local:      true,
		Dimensions: c.settings.Dimensions,
	}
}

func (c *client) Dimensions() int { return c.settings.Dimensions }

type chatRequest struct {
	Model    string            `json:"model"`
	Messages []provider.Message `json:"messages"`
	Stream   bool              `json:"stream"`
	Options  map[string]any    `json:"options,omitempty"`
}

type chatResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
	Done            bool   `json:"done"`
	PromptEvalCount int    `json:"prompt_eval_count"`
	EvalCount       int    `json:"eval_count"`
	Error           string `json:"error"`
}

func (c *client) options(req provider.GenerateRequest) map[string]any {
	o := map[string]any{}
	if req.Temperature > 0 {
		o["temperature"] = req.Temperature
	}
	if req.MaxTokens > 0 {
		o["num_predict"] = req.MaxTokens
	}
	if len(req.Stop) > 0 {
		o["stop"] = req.Stop
	}
	return o
}

func (c *client) post(ctx context.Context, path string, body any) (*http.Response, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.settings.BaseURL+path, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(httpReq)
	if err != nil {
		// A stopped Ollama is a deployment state, not a bug — surface it as
		// ErrUnavailable so the API can answer 503 rather than 500.
		return nil, fmt.Errorf("%w: %v", provider.ErrUnavailable, err)
	}
	if res.StatusCode >= 400 {
		res.Body.Close()
		return nil, fmt.Errorf("%w: ollama returned %d", provider.ErrUnavailable, res.StatusCode)
	}
	return res, nil
}

func (c *client) Generate(ctx context.Context, req provider.GenerateRequest) (*provider.GenerateResponse, error) {
	res, err := c.post(ctx, "/api/chat", chatRequest{
		Model: c.settings.Model, Messages: req.Messages, Stream: false, Options: c.options(req),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var out chatResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	if out.Error != "" {
		return nil, fmt.Errorf("ollama: %s", out.Error)
	}
	return &provider.GenerateResponse{
		Text: out.Message.Content,
		Usage: provider.Usage{
			PromptTokens:     out.PromptEvalCount,
			CompletionTokens: out.EvalCount,
			Provider:         "ollama",
			Model:            c.settings.Model,
		},
	}, nil
}

func (c *client) Stream(ctx context.Context, req provider.GenerateRequest) (<-chan provider.Chunk, error) {
	res, err := c.post(ctx, "/api/chat", chatRequest{
		Model: c.settings.Model, Messages: req.Messages, Stream: true, Options: c.options(req),
	})
	if err != nil {
		return nil, err
	}

	ch := make(chan provider.Chunk)
	go func() {
		defer close(ch)
		defer res.Body.Close()

		scanner := bufio.NewScanner(res.Body)
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

		for scanner.Scan() {
			line := bytes.TrimSpace(scanner.Bytes())
			if len(line) == 0 {
				continue
			}
			var frame chatResponse
			if err := json.Unmarshal(line, &frame); err != nil {
				continue // a malformed frame should not kill the stream
			}
			if frame.Error != "" {
				send(ctx, ch, provider.Chunk{Err: fmt.Errorf("ollama: %s", frame.Error)})
				return
			}
			if frame.Message.Content != "" {
				if !send(ctx, ch, provider.Chunk{Text: frame.Message.Content}) {
					return
				}
			}
			if frame.Done {
				send(ctx, ch, provider.Chunk{Done: true, Usage: provider.Usage{
					PromptTokens:     frame.PromptEvalCount,
					CompletionTokens: frame.EvalCount,
					Provider:         "ollama",
					Model:            c.settings.Model,
				}})
				return
			}
		}
		if err := scanner.Err(); err != nil {
			send(ctx, ch, provider.Chunk{Err: err})
		}
	}()
	return ch, nil
}

// send respects cancellation so an abandoned request cannot leak a goroutine
// blocked forever on an unread channel.
func send(ctx context.Context, ch chan<- provider.Chunk, c provider.Chunk) bool {
	select {
	case ch <- c:
		return true
	case <-ctx.Done():
		return false
	}
}

type embedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type embedResponse struct {
	Embeddings [][]float32 `json:"embeddings"`
	Error      string      `json:"error"`
}

func (c *client) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}
	res, err := c.post(ctx, "/api/embed", embedRequest{Model: c.settings.Model, Input: texts})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var out embedResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	if out.Error != "" {
		return nil, fmt.Errorf("ollama: %s", out.Error)
	}
	if len(out.Embeddings) != len(texts) {
		return nil, fmt.Errorf("ollama returned %d embeddings for %d inputs", len(out.Embeddings), len(texts))
	}
	// Guard the schema invariant here rather than letting Postgres reject the
	// insert with a less obvious error.
	if d := c.settings.Dimensions; d > 0 && len(out.Embeddings[0]) != d {
		return nil, fmt.Errorf("model %q returned %d dimensions, expected %d — the schema and EMBEDDING_DIMENSIONS must agree",
			c.settings.Model, len(out.Embeddings[0]), d)
	}
	return out.Embeddings, nil
}
