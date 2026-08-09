// Package store defines persistence boundaries.
package store

import (
	"context"

	"github.com/google/uuid"
)

// Chunk is an indexable passage with everything needed to cite it.
type Chunk struct {
	ID         uuid.UUID
	CourseID   uuid.UUID
	DocumentID uuid.UUID
	QuestionID *uuid.UUID
	ChapterID  *uuid.UUID

	Ordinal int
	Text    string
	Page    int

	Embedding      []float32
	EmbeddingModel string

	// SourceWeight implements source-priority ranking: a user's own notes can
	// be weighted above shared material when both match a query.
	SourceWeight float32
}

// Filters are applied as WHERE clauses before ranking, never as post-filtering.
// Nil fields mean "no constraint".
type Filters struct {
	ChapterIDs     []uuid.UUID
	ChapterNumbers []int
	YearFrom       *int
	YearTo         *int
	DocumentKinds  []string
}

// Query is one hybrid retrieval request.
type Query struct {
	CourseID uuid.UUID
	Text     string    // drives the sparse/full-text side
	Vector   []float32 // drives the dense side
	Filters  Filters
	Limit    int
}

// Match is a ranked passage with its citation payload.
type Match struct {
	ChunkID     uuid.UUID
	Text        string
	Page        int
	Score       float64
	DenseScore  float64
	SparseScore float64

	DocumentID     uuid.UUID
	DocumentTitle  string
	DocumentKind   string
	QuestionID     *uuid.UUID
	QuestionNumber string
}

// VectorStore abstracts the vector index.
//
// The MVP implementation is PostgreSQL + pgvector, which lets hybrid ranking,
// structured filters, joins and aggregates run in a single statement and a
// single ACID transaction alongside the relational data.
//
// This interface exists so that decision stays reversible. The trigger to add a
// Qdrant implementation is a filtered candidate set consistently above ~100k
// rows, or adopting late-interaction (ColBERT) reranking — the two places where
// a dedicated vector database genuinely wins. Until then, a separate store
// would add a dual-write consistency problem in exchange for nothing.
type VectorStore interface {
	Upsert(ctx context.Context, chunks []Chunk) error
	Search(ctx context.Context, q Query) ([]Match, error)
	DeleteByDocument(ctx context.Context, documentID uuid.UUID) error
}
