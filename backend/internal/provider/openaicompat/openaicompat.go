// Package openaicompat adapts any OpenAI-compatible endpoint to the provider
// interfaces.
//
// One adapter covers a large share of hosted and self-hosted options, vLLM,
// TGI, LM Studio, llama.cpp's server, Groq, Together, and OpenAI itself, // because they all speak the same wire format. Only the base URL and key
// change, which is the whole argument for the interface.
package openaicompat

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

func init() {
	for _, name := range []string{"openai-compatible", "openai", "vllm", "tgi", "groq", "together"} {
		provider.RegisterLLM(name, func(s provider.Settings) (provider.LLM, error) {
			return &client{settings: s, http: newHTTPClient()}, nil
		})
		provider.RegisterEmbedder(name, func(s provider.Settings) (provider.Embedder, error) {
			return &client{settings: s, http: newHTTPClient()}, nil
		})
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
			MaxIdleConnsPerHost: 8,
		},
	}
}

type client struct {
	settings provider.Settings
	http     *http.Client
}

func (c *client) Info() provider.Info {
	return provider.Info{
		Provider:   c.settings.Provider,
		Model:      c.settings.Model,
		Local:      false,
		Dimensions: c.settings.Dimensions,
	}
}

func (c *client) Dimensions() int { return c.settings.Dimensions }

type chatRequest struct {
	Model       string             `json:"model"`
	Messages    []provider.Message `json:"messages"`
	Stream      bool               `json:"stream"`
	MaxTokens   int                `json:"max_tokens,omitempty"`
	Temperature float32            `json:"temperature,omitempty"`
	Stop        []string           `json:"stop,omitempty"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason *string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

func (c *client) post(ctx context.Context, path string, body any) (*http.Response, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		strings.TrimRight(c.settings.BaseURL, "/")+path, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.settings.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.settings.APIKey)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", provider.ErrUnavailable, err)
	}

	switch {
	case res.StatusCode == http.StatusTooManyRequests:
		res.Body.Close()
		return nil, provider.ErrRateLimited
	case res.StatusCode == http.StatusRequestEntityTooLarge:
		res.Body.Close()
		return nil, provider.ErrContextTooLong
	case res.StatusCode >= 400:
		defer res.Body.Close()
		var e chatResponse
		_ = json.NewDecoder(res.Body).Decode(&e)
		if e.Error != nil {
			// Providers signal an over-long prompt in the message rather than
			// the status, and the caller's remedy (retrieve fewer chunks)
			// differs from a generic failure.
			if strings.Contains(strings.ToLower(e.Error.Message), "context length") ||
				strings.Contains(strings.ToLower(e.Error.Message), "too many tokens") {
				return nil, provider.ErrContextTooLong
			}
			return nil, fmt.Errorf("%s: %s", c.settings.Provider, e.Error.Message)
		}
		return nil, fmt.Errorf("%s returned %d", c.settings.Provider, res.StatusCode)
	}
	return res, nil
}

func (c *client) Generate(ctx context.Context, req provider.GenerateRequest) (*provider.GenerateResponse, error) {
	res, err := c.post(ctx, "/chat/completions", chatRequest{
		Model: c.settings.Model, Messages: req.Messages, Stream: false,
		MaxTokens: req.MaxTokens, Temperature: req.Temperature, Stop: req.Stop,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var out chatResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	if len(out.Choices) == 0 {
		return nil, fmt.Errorf("%s returned no choices", c.settings.Provider)
	}
	return &provider.GenerateResponse{
		Text: out.Choices[0].Message.Content,
		Usage: provider.Usage{
			PromptTokens:     out.Usage.PromptTokens,
			CompletionTokens: out.Usage.CompletionTokens,
			Provider:         c.settings.Provider,
			Model:            c.settings.Model,
		},
	}, nil
}

func (c *client) Stream(ctx context.Context, req provider.GenerateRequest) (<-chan provider.Chunk, error) {
	res, err := c.post(ctx, "/chat/completions", chatRequest{
		Model: c.settings.Model, Messages: req.Messages, Stream: true,
		MaxTokens: req.MaxTokens, Temperature: req.Temperature, Stop: req.Stop,
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
		completion := 0

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if !strings.HasPrefix(line, "data:") {
				continue
			}
			payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			if payload == "[DONE]" {
				break
			}
			var frame chatResponse
			if err := json.Unmarshal([]byte(payload), &frame); err != nil {
				continue
			}
			if len(frame.Choices) == 0 {
				continue
			}
			if text := frame.Choices[0].Delta.Content; text != "" {
				completion++
				select {
				case ch <- provider.Chunk{Text: text}:
				case <-ctx.Done():
					return
				}
			}
		}

		select {
		case ch <- provider.Chunk{Done: true, Usage: provider.Usage{
			CompletionTokens: completion,
			Provider:         c.settings.Provider,
			Model:            c.settings.Model,
		}}:
		case <-ctx.Done():
		}
	}()
	return ch, nil
}

type embedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`

	// Matryoshka models (OpenAI text-embedding-3-*, Gemini) are trained so a
	// truncated prefix of the vector is still a usable embedding, and expose
	// that as a request parameter. The schema pins vectors to a fixed width,
	// so asking for it is the difference between a model being usable here and
	// not.
	//
	// Sent whenever a width is configured. A server that does not understand
	// the parameter rejects the request, which is the same outcome as the
	// width check below rejecting a differently-sized vector, but with a
	// clearer error and before anything is stored.
	Dimensions int `json:"dimensions,omitempty"`
}

type embedResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
}

func (c *client) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}
	res, err := c.post(ctx, "/embeddings", embedRequest{
		Model:      c.settings.Model,
		Input:      texts,
		Dimensions: c.settings.Dimensions,
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var out embedResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	if len(out.Data) != len(texts) {
		return nil, fmt.Errorf("%s returned %d embeddings for %d inputs", c.settings.Provider, len(out.Data), len(texts))
	}

	// The API does not guarantee input order, and silently mismatched vectors
	// would be nearly impossible to debug later.
	vectors := make([][]float32, len(texts))
	for _, d := range out.Data {
		if d.Index < 0 || d.Index >= len(vectors) {
			return nil, fmt.Errorf("%s returned out-of-range index %d", c.settings.Provider, d.Index)
		}
		vectors[d.Index] = d.Embedding
	}
	if d := c.settings.Dimensions; d > 0 && len(vectors[0]) != d {
		return nil, fmt.Errorf("model %q returned %d dimensions, expected %d", c.settings.Model, len(vectors[0]), d)
	}
	return vectors, nil
}
