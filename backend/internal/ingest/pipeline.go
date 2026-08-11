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
	// ClearChunks and AppendChunks let a long document be written in batches,
	// so memory stays flat rather than scaling with page count.
	ClearChunks(ctx context.Context, documentID uuid.UUID) error
	AppendChunks(ctx context.Context, documentID uuid.UUID, chunks []StoredChunk) error
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

func NewPipeline(store Store, embedder provider.Embedder, llm provider.LLM, ocrURL string, log *slog.Logger) *Pipeline {
	chain := &Chain{}
	chain.Register(PDFParser{})
	chain.Register(DOCXParser{})
	// Registering the OCR adapter is the entire integration. With OCR_URL
	// unset this is skipped and a photograph is refused with a message naming
	// the missing capability, rather than being indexed as nothing.
	if ocr := NewOCRParser(ocrURL); ocr != nil {
		chain.Register(ocr)
		// The same service also extracts an existing text layer, which it does
		// far better than the pure-Go reader, so it is registered separately
		// and outranks the fallback on quality.
		chain.Register(NewPDFTextParser(ocrURL))
		log.Info("document service registered", "url", ocrURL,
			"parsers", "pdftotext, tesseract-ocr")
	}
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

	format, probe := ProbeDocument(job.Content)
	if format == FormatUnknown {
		return fail("parsing", fmt.Errorf("unrecognised file type, supported: %v", SupportedFormats()))
	}
	// A DOCX is never routed to OCR: it either has text or it is empty, and
	// running OCR over an empty one would waste minutes to produce nothing.
	if format == FormatDOCX {
		probe.LikelyScanned = false
	}

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

// Limits for the passage path.
//
// A 16 MB textbook took the whole Docker VM down: several hundred pages became
// thousands of chunks, every chunk and every 768-float vector held in memory at
// once, next to Postgres and a local model in the same VM. Exams are bounded by
// their question count; a textbook is bounded by nothing.
const (
	// maxPassagePages caps how much of a very long document is indexed. Beyond
	// this the document is rejected rather than half-indexed, because a corpus
	// silently containing the first third of a textbook answers questions from
	// it and stays quiet about the rest.
	maxPassagePages = 600

	// passageBatch is how many chunks are embedded and written at a time.
	// Memory stays flat regardless of document length.
	passageBatch = 32
)

// indexPassagesOnly handles notes and textbooks: no questions, but the text is
// still worth retrieving over.
//
// Streams in batches rather than building the whole document in memory. The
// previous version collected every chunk, then every vector, then wrote once,
// so peak memory scaled with document size and a large enough upload killed the
// process rather than failing.
func (p *Pipeline) indexPassagesOnly(ctx context.Context, job Job, doc *ParsedDocument) error {
	if len(doc.Pages) > maxPassagePages {
		err := fmt.Errorf("%d pages exceeds the %d page limit for reference material; split the document or index the relevant chapters separately",
			len(doc.Pages), maxPassagePages)
		_ = p.store.FailDocument(ctx, job.DocumentID, err.Error())
		return err
	}

	_ = p.store.SetStatus(ctx, job.DocumentID, "embedding", 0.7, "indexing passages")

	// Re-ingestion is idempotent: clear once, then append batch by batch.
	if err := p.store.ClearChunks(ctx, job.DocumentID); err != nil {
		return err
	}

	var batch []StoredChunk
	total := 0

	flush := func() error {
		if len(batch) == 0 {
			return nil
		}
		texts := make([]string, len(batch))
		for i, c := range batch {
			texts[i] = c.Text
		}
		vectors, err := p.embedder.Embed(ctx, texts)
		if err != nil {
			return err
		}
		for i := range batch {
			batch[i].Embedding = vectors[i]
		}
		if err := p.store.AppendChunks(ctx, job.DocumentID, batch); err != nil {
			return err
		}
		total += len(batch)
		// Reuse the slice: the vectors are now the database's problem, not
		// this process's.
		batch = batch[:0]
		return nil
	}

	for pageIdx, page := range doc.Pages {
		for i, part := range chunkText(page.Text, 1200, 150) {
			batch = append(batch, StoredChunk{
				QuestionIndex: -1, Ordinal: i + 1, Text: part, Page: page.Number,
			})
			if len(batch) >= passageBatch {
				if err := flush(); err != nil {
					_ = p.store.FailDocument(ctx, job.DocumentID, "embedding: "+err.Error())
					return err
				}
			}
		}

		// A long document spends minutes here, so report movement. Without it
		// the dialog sits at one number and looks hung.
		if pageIdx%10 == 0 {
			_ = p.store.SetStatus(ctx, job.DocumentID, "embedding",
				0.7+0.25*float64(pageIdx)/float64(len(doc.Pages)),
				fmt.Sprintf("indexing page %d of %d", pageIdx+1, len(doc.Pages)))
		}
	}

	if err := flush(); err != nil {
		_ = p.store.FailDocument(ctx, job.DocumentID, "embedding: "+err.Error())
		return err
	}

	p.log.Info("indexed as passages",
		"document", job.DocumentID, "pages", doc.PageCount, "chunks", total)
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
