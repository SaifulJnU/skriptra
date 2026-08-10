package cache

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

// memCache is a Cache for tests, so caching behaviour is verified without a
// Redis server.
type memCache struct {
	mu   sync.Mutex
	data map[string][]byte
}

func newMem() *memCache { return &memCache{data: map[string][]byte{}} }

func (m *memCache) Enabled() bool { return true }
func (m *memCache) Get(_ context.Context, k string) ([]byte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[k]
	return v, ok
}
func (m *memCache) Set(_ context.Context, k string, v []byte, _ time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}
func (m *memCache) DeletePrefix(_ context.Context, p string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k := range m.data {
		if len(k) >= len(p) && k[:len(p)] == p {
			delete(m.data, k)
		}
	}
}

// countingEmbedder records exactly which texts reached the real provider.
type countingEmbedder struct {
	calls    int
	embedded []string
	dims     int
	model    string
	err      error
}

func (c *countingEmbedder) Dimensions() int { return c.dims }
func (c *countingEmbedder) Info() provider.Info {
	return provider.Info{Provider: "stub", Model: c.model, Dimensions: c.dims}
}
func (c *countingEmbedder) Embed(_ context.Context, texts []string) ([][]float32, error) {
	if c.err != nil {
		return nil, c.err
	}
	c.calls++
	c.embedded = append(c.embedded, texts...)
	out := make([][]float32, len(texts))
	for i, t := range texts {
		v := make([]float32, c.dims)
		// Deterministic and text-dependent, so a wrong cache hit is visible.
		v[0] = float32(len(t))
		v[1] = float32(t[0])
		out[i] = v
	}
	return out, nil
}

func TestEmbedCachesRepeatedText(t *testing.T) {
	inner := &countingEmbedder{dims: 4, model: "m1"}
	e := NewEmbedder(inner, newMem())
	ctx := context.Background()

	first, err := e.Embed(ctx, []string{"alpha", "beta"})
	if err != nil {
		t.Fatalf("Embed() error = %v", err)
	}
	second, err := e.Embed(ctx, []string{"alpha", "beta"})
	if err != nil {
		t.Fatalf("Embed() error = %v", err)
	}

	if inner.calls != 1 {
		t.Errorf("provider called %d times, want 1 (second call should be all hits)", inner.calls)
	}
	for i := range first {
		if first[i][0] != second[i][0] || first[i][1] != second[i][1] {
			t.Errorf("cached vector %d differs from the original", i)
		}
	}
}

// The case that matters when re-ingesting a document after a parser fix: most
// questions are unchanged, so only the new ones should reach the provider.
func TestEmbedOnlySendsMisses(t *testing.T) {
	inner := &countingEmbedder{dims: 4, model: "m1"}
	e := NewEmbedder(inner, newMem())
	ctx := context.Background()

	if _, err := e.Embed(ctx, []string{"a", "b", "c"}); err != nil {
		t.Fatal(err)
	}
	inner.embedded = nil

	got, err := e.Embed(ctx, []string{"a", "NEW", "c"})
	if err != nil {
		t.Fatal(err)
	}

	if len(inner.embedded) != 1 || inner.embedded[0] != "NEW" {
		t.Errorf("provider received %v, want only [NEW]", inner.embedded)
	}
	// Order must survive the partial fill.
	if got[1][1] != float32('N') {
		t.Errorf("results came back out of order: index 1 is not the new text")
	}
	if got[0][1] != float32('a') || got[2][1] != float32('c') {
		t.Errorf("cached entries landed at the wrong indices")
	}
}

// Changing the embedding model must never return vectors from the previous one.
// They are the right shape and mean nothing, which is close to undiagnosable.
func TestEmbedKeysIncludeModel(t *testing.T) {
	shared := newMem()
	ctx := context.Background()

	old := &countingEmbedder{dims: 4, model: "old-model"}
	if _, err := NewEmbedder(old, shared).Embed(ctx, []string{"alpha"}); err != nil {
		t.Fatal(err)
	}

	fresh := &countingEmbedder{dims: 4, model: "new-model"}
	if _, err := NewEmbedder(fresh, shared).Embed(ctx, []string{"alpha"}); err != nil {
		t.Fatal(err)
	}

	if fresh.calls != 1 {
		t.Error("a different model reused the previous model's cached vector")
	}
}

// A corrupt or wrong-width entry must degrade to a miss, never to an error.
func TestEmbedTreatsBadEntryAsMiss(t *testing.T) {
	mem := newMem()
	inner := &countingEmbedder{dims: 4, model: "m1"}
	e := NewEmbedder(inner, mem)
	ctx := context.Background()

	if _, err := e.Embed(ctx, []string{"alpha"}); err != nil {
		t.Fatal(err)
	}
	for k := range mem.data {
		mem.data[k] = []byte{1, 2, 3} // not a multiple of 4, and wrong width
	}

	got, err := e.Embed(ctx, []string{"alpha"})
	if err != nil {
		t.Fatalf("a corrupt cache entry broke embedding: %v", err)
	}
	if len(got[0]) != 4 {
		t.Errorf("got %d dimensions, want 4 recomputed", len(got[0]))
	}
}

func TestEmbedPropagatesProviderError(t *testing.T) {
	inner := &countingEmbedder{dims: 4, model: "m1", err: errors.New("provider down")}
	e := NewEmbedder(inner, newMem())

	if _, err := e.Embed(context.Background(), []string{"alpha"}); err == nil {
		t.Error("provider error was swallowed")
	}
}

// With no cache configured the wrapper is not applied at all, so the call stack
// stays honest.
func TestNoOpReturnsInnerUnwrapped(t *testing.T) {
	inner := &countingEmbedder{dims: 4, model: "m1"}
	if got := NewEmbedder(inner, NoOp{}); got != provider.Embedder(inner) {
		t.Error("embedder was wrapped despite caching being disabled")
	}
}

func TestVectorRoundTrip(t *testing.T) {
	in := []float32{0, 1, -1.5, 3.14159, 1e-8}
	out, err := decodeVector(encodeVector(in))
	if err != nil {
		t.Fatal(err)
	}
	for i := range in {
		if in[i] != out[i] {
			t.Errorf("index %d: got %v, want %v", i, out[i], in[i])
		}
	}
}
