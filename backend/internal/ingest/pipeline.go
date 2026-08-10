package ingest

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/provider"
)

// Store is the persistence the pipeline needs.
//
// Defined here, in the consumer, rather than in the db package. The pipeline
// states what it requires; the database happens to satisfy it. That keeps this
// package testable with a fake and free of SQL.
type Store interface {
	SetStatus(ctx context.Context, documentID uuid.UUID, status string, progress float64, detail string) error
	FailDocument(ctx context.Context, documentID uuid.UUID, reason string) error
	ChaptersFor(ctx context.Context, courseID uuid.UUID) ([]Chapter, error)
	ReplaceQuestions(ctx context.Context, documentID uuid.UUID, questions []StoredQuestion) ([]uuid.UUID, error)
	ReplaceChunks(ctx context.Context, documentID uuid.UUID, chunks []StoredChunk) error
	SetQuestionEmbeddings(ctx context.Context, courseID uuid.UUID, ids []uuid.UUID, vectors [][]float32, model string) error
	FinishDocument(ctx context.Context, documentID uuid.UUID, pageCount, questionCount int) error
}

type StoredQuestion struct {
	Number        string
	Ordinal       int
	Text          string
	Marks         *float64
	SourcePage    int
	ChapterNumber int
	Confidence    *float64
	Source        string
	Topic         string
	Type          QuestionType
}

type StoredChunk struct {
	QuestionIndex int // -1 when the chunk is page text with no owning question
	ChapterNumber int
	Ordinal       int
	Text          string
	Page          int
	Embedding     []float32
}

// Job is one ingestion request.
type Job struct {
	DocumentID uuid.UUID
	CourseID   uuid.UUID
	Filename   string
	Content    []byte
}

// Pipeline turns an uploaded document into indexed, citable questions.
type Pipeline struct {
	parsers  *Chain
	store    Store
	embedder provider.Embedder
	llm      provider.LLM
	log      *slog.Logger
}

func NewPipeline(store Store, embedder provider.Embedder, llm provider.LLM, log *slog.Logger) *Pipeline {
	chain := &Chain{}
	chain.Register(PDFParser{})
	return &Pipeline{parsers: chain, store: store, embedder: embedder, llm: llm, log: log}
}

// RegisterParser adds a parser implementation. This is the seam a Python OCR
// sidecar plugs into: register it and scanned documents start routing to it,
// with no other change anywhere.
func (p *Pipeline) RegisterParser(parser DocumentParser) { p.parsers.Register(parser) }

// Run executes the full pipeline.
//
// Progress is reported at each stage because ingesting a long paper takes real
// time and a silent spinner is indistinguishable from a hang.
func (p *Pipeline) Run(ctx context.Context, job Job) error {
	fail := func(stage string, err error) error {
		p.log.Error("ingest failed", "document", job.DocumentID, "stage", stage, "err", err)
		// Best-effort: the original error matters more than a failure to
		// record it.
		_ = p.store.FailDocument(ctx, job.DocumentID, fmt.Sprintf("%s: %v", stage, err))
		return fmt.Errorf("%s: %w", stage, err)
	}

	// --- parse ---------------------------------------------------------
	_ = p.store.SetStatus(ctx, job.DocumentID, "parsing", 0.1, "extracting text")

	probe := ProbePDF(job.Content)
	doc, err := p.parsers.Parse(ctx, bytes.NewReader(job.Content), job.Filename, probe)
	if err != nil {
		return fail("parsing", err)
	}

	// --- segment -------------------------------------------------------
	_ = p.store.SetStatus(ctx, job.DocumentID, "segmenting", 0.3, "splitting into questions")

	questions := SegmentQuestions(doc.Pages)
	if len(questions) == 0 {
		// Not an error. Notes and textbooks have no questions; they are still
		// worth indexing as searchable passages.
		p.log.Info("no questions found, indexing as passages",
			"document", job.DocumentID, "pages", doc.PageCount)
		return p.indexPassagesOnly(ctx, job, doc)
	}

	// --- classify ------------------------------------------------------
	_ = p.store.SetStatus(ctx, job.DocumentID, "classifying", 0.45,
		fmt.Sprintf("classifying %d questions", len(questions)))

	chapters, err := p.store.ChaptersFor(ctx, job.CourseID)
	if err != nil {
		return fail("classifying", err)
	}
	classifier := NewClassifier(chapters, p.llm)

	stored := make([]StoredQuestion, len(questions))
	for i, q := range questions {
		c := classifier.Classify(ctx, q.Text)
		sq := StoredQuestion{
			Number: q.Number, Ordinal: q.Ordinal, Text: q.Text,
			Marks: q.Marks, SourcePage: q.SourcePage,
			ChapterNumber: c.ChapterNumber, Source: c.Source, Topic: c.Topic,
			// Format is independent of topic, so it is derived from the
			// wording rather than from the chapter taxonomy.
			Type: ClassifyType(q.Text),
		}
		if c.ChapterNumber != 0 {
			conf := c.Confidence
			sq.Confidence = &conf
		}
		stored[i] = sq

		if i%5 == 0 {
			_ = p.store.SetStatus(ctx, job.DocumentID, "classifying",
				0.45+0.2*float64(i)/float64(len(questions)),
				fmt.Sprintf("classifying question %d of %d", i+1, len(questions)))
		}
	}

	questionIDs, err := p.store.ReplaceQuestions(ctx, job.DocumentID, stored)
	if err != nil {
		return fail("storing questions", err)
	}

	// --- embed ---------------------------------------------------------
	_ = p.store.SetStatus(ctx, job.DocumentID, "embedding", 0.7, "computing embeddings")

	texts := make([]string, len(questions))
	for i, q := range questions {
		texts[i] = q.Text
	}
	vectors, err := p.embedBatched(ctx, texts)
	if err != nil {
		return fail("embedding", err)
	}

	chunks := make([]StoredChunk, len(questions))
	for i, q := range questions {
		chunks[i] = StoredChunk{
			QuestionIndex: i,
			ChapterNumber: stored[i].ChapterNumber,
			Ordinal:       1,
			Text:          q.Text,
			Page:          q.SourcePage,
			Embedding:     vectors[i],
		}
	}
	if err := p.store.ReplaceChunks(ctx, job.DocumentID, chunks); err != nil {
		return fail("storing chunks", err)
	}

	// A question is embedded twice on purpose: once as a retrievable chunk and
	// once as a whole-question vector. "Similar questions" compares complete
	// questions, which is a different unit from a passage window.
	if err := p.store.SetQuestionEmbeddings(ctx, job.CourseID, questionIDs, vectors,
		p.embedder.Info().Model); err != nil {
		return fail("storing question embeddings", err)
	}

	if err := p.store.FinishDocument(ctx, job.DocumentID, doc.PageCount, len(questions)); err != nil {
		return fail("finalising", err)
	}

	p.log.Info("ingested",
		"document", job.DocumentID, "pages", doc.PageCount,
		"questions", len(questions), "parser", doc.ParsedBy)
	return nil
}

// indexPassagesOnly handles notes and textbooks: no questions, but the text is
// still worth retrieving over.
func (p *Pipeline) indexPassagesOnly(ctx context.Context, job Job, doc *ParsedDocument) error {
	_ = p.store.SetStatus(ctx, job.DocumentID, "embedding", 0.7, "indexing passages")

	var chunks []StoredChunk
	var texts []string
	for _, page := range doc.Pages {
		for i, part := range chunkText(page.Text, 1200, 150) {
			chunks = append(chunks, StoredChunk{
				QuestionIndex: -1, Ordinal: i + 1, Text: part, Page: page.Number,
			})
			texts = append(texts, part)
		}
	}
	if len(chunks) == 0 {
		return p.store.FinishDocument(ctx, job.DocumentID, doc.PageCount, 0)
	}

	vectors, err := p.embedBatched(ctx, texts)
	if err != nil {
		_ = p.store.FailDocument(ctx, job.DocumentID, "embedding: "+err.Error())
		return err
	}
	for i := range chunks {
		chunks[i].Embedding = vectors[i]
	}
	if err := p.store.ReplaceChunks(ctx, job.DocumentID, chunks); err != nil {
		return err
	}
	return p.store.FinishDocument(ctx, job.DocumentID, doc.PageCount, 0)
}

// embedBatched keeps request sizes bounded. A 200-page textbook is thousands of
// chunks, and sending them in one call either times out or exhausts the
// provider's request limit.
func (p *Pipeline) embedBatched(ctx context.Context, texts []string) ([][]float32, error) {
	const batch = 32
	out := make([][]float32, 0, len(texts))

	for start := 0; start < len(texts); start += batch {
		end := start + batch
		if end > len(texts) {
			end = len(texts)
		}
		vectors, err := p.embedder.Embed(ctx, texts[start:end])
		if err != nil {
			return nil, err
		}
		out = append(out, vectors...)
	}
	return out, nil
}

// chunkText splits prose into overlapping windows on sentence boundaries.
//
// The overlap exists so a claim spanning a window edge is still retrievable
// intact from at least one chunk.
func chunkText(text string, size, overlap int) []string {
	text = collapseWhitespace(text)
	if len(text) <= size {
		if strings.TrimSpace(text) == "" {
			return nil
		}
		return []string{text}
	}

	var out []string
	for start := 0; start < len(text); {
		end := start + size
		if end >= len(text) {
			out = append(out, strings.TrimSpace(text[start:]))
			break
		}
		// Prefer to cut at a sentence end inside the last quarter of the window.
		cut := end
		for i := end; i > start+size*3/4; i-- {
			if text[i] == '.' || text[i] == '?' || text[i] == '!' {
				cut = i + 1
				break
			}
		}
		out = append(out, strings.TrimSpace(text[start:cut]))
		start = cut - overlap
		if start < 0 {
			start = 0
		}
	}
	return out
}
