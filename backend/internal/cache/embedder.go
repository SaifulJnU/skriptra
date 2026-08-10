package cache

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

// EmbedTTL is long because an embedding is a pure function of its text and
// model. The only thing that invalidates it is changing the model, and that is
// already part of the key.
const EmbedTTL = 30 * 24 * time.Hour

// CachedEmbedder wraps an Embedder with a read-through cache.
//
// A decorator rather than a change inside each adapter: it implements the same
// interface, so every provider gains caching and no provider knows about it.
//
// This is where a cache pays for itself. Embedding is deterministic, and the
// same text is embedded repeatedly in practice: a student asking a question
// twice, the eval harness replaying a golden set on every run, and re-ingesting
// a document after a parser fix, where nearly every question is unchanged.
type CachedEmbedder struct {
	inner provider.Embedder
	cache Cache
	model string
}

func NewEmbedder(inner provider.Embedder, c Cache) provider.Embedder {
	if c == nil || !c.Enabled() {
		// No wrapper at all rather than one that always misses. Keeps the call
		// stack honest when reading a stack trace.
		return inner
	}
	return &CachedEmbedder{inner: inner, cache: c, model: inner.Info().Model}
}

func (c *CachedEmbedder) Dimensions() int      { return c.inner.Dimensions() }
func (c *CachedEmbedder) Info() provider.Info  { return c.inner.Info() }

// Embed returns cached vectors where available and computes only the misses.
//
// Partial hits matter: a batch of 32 chunks from a re-ingested document may have
// 30 unchanged. Sending only the two new ones is most of the saving, and the
// results still come back in the caller's original order.
func (c *CachedEmbedder) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	out := make([][]float32, len(texts))
	var missIdx []int
	var missTexts []string

	for i, t := range texts {
		if raw, ok := c.cache.Get(ctx, c.key(t)); ok {
			if v, err := decodeVector(raw); err == nil && len(v) == c.inner.Dimensions() {
				out[i] = v
				continue
			}
			// A corrupt or wrong-width entry is treated as a miss rather than
			// an error: a stale encoding must not be able to break embedding.
		}
		missIdx = append(missIdx, i)
		missTexts = append(missTexts, t)
	}

	if len(missTexts) == 0 {
		return out, nil
	}

	fresh, err := c.inner.Embed(ctx, missTexts)
	if err != nil {
		return nil, err
	}
	if len(fresh) != len(missTexts) {
		return nil, fmt.Errorf("embedder returned %d vectors for %d inputs", len(fresh), len(missTexts))
	}

	for j, idx := range missIdx {
		out[idx] = fresh[j]
		c.cache.Set(ctx, c.key(missTexts[j]), encodeVector(fresh[j]), EmbedTTL)
	}
	return out, nil
}

// key includes the model, so switching models cannot return vectors from the
// previous one. That would be silent and near-impossible to diagnose: the
// numbers are the right shape and mean nothing.
func (c *CachedEmbedder) key(text string) string {
	sum := sha256.Sum256([]byte(text))
	return fmt.Sprintf("emb:%s:%s", c.model, hex.EncodeToString(sum[:16]))
}

// Vectors are stored as raw little-endian float32 rather than JSON: 768 floats
// are 3 KB packed against roughly 9 KB as text, and this runs on every chunk.
func encodeVector(v []float32) []byte {
	buf := make([]byte, 4*len(v))
	for i, f := range v {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(f))
	}
	return buf
}

func decodeVector(b []byte) ([]float32, error) {
	if len(b)%4 != 0 {
		return nil, fmt.Errorf("cached vector has length %d, not a multiple of 4", len(b))
	}
	v := make([]float32, len(b)/4)
	for i := range v {
		v[i] = math.Float32frombits(binary.LittleEndian.Uint32(b[i*4:]))
	}
	return v, nil
}
