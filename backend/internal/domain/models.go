// Package domain holds the wire types.
//
// These mirror `api/openapi.yaml` exactly, including the JSON tags. The
// contract is the source of truth; if a field here disagrees with the spec,
// the spec wins and this file is wrong.
package domain

import (
	"time"

	"github.com/google/uuid"
)

type Term string

const (
	TermSummer Term = "summer"
	TermWinter Term = "winter"
)

type DocumentKind string

type IngestStatus string

const (
	StatusQueued      IngestStatus = "queued"
	StatusParsing     IngestStatus = "parsing"
	StatusSegmenting  IngestStatus = "segmenting"
	StatusClassifying IngestStatus = "classifying"
	StatusEmbedding   IngestStatus = "embedding"
	StatusIndexed     IngestStatus = "indexed"
	StatusFailed      IngestStatus = "failed"
)

// QueryIntent is the router's decision. Only Explain and Hybrid reach a model.
type QueryIntent string

const (
	IntentEnumerate QueryIntent = "enumerate"
	IntentExplain   QueryIntent = "explain"
	IntentSimilar   QueryIntent = "similar"
	IntentAnalyse   QueryIntent = "analyse"
	IntentHybrid    QueryIntent = "hybrid"
)

type PageMeta struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

// Paged is the envelope every list endpoint returns.
type Paged[T any] struct {
	Data []T      `json:"data"`
	Meta PageMeta `json:"meta"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email,omitempty"`
}

type ProviderInfo struct {
	Provider   string `json:"provider"`
	Model      string `json:"model"`
	Local      bool   `json:"local"`
	Dimensions int    `json:"dimensions,omitempty"`
}

type Providers struct {
	LLM       ProviderInfo `json:"llm"`
	Embedding ProviderInfo `json:"embedding"`
}

type Course struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Code          string     `json:"code,omitempty"`
	Institution   string     `json:"institution,omitempty"`
	Language      string     `json:"language,omitempty"`
	ExamCount     int        `json:"examCount"`
	QuestionCount int        `json:"questionCount"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
}

type YearRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type CourseDetail struct {
	Course
	ChapterCount  int        `json:"chapterCount"`
	DocumentCount int        `json:"documentCount"`
	YearRange     *YearRange `json:"yearRange,omitempty"`
}

type Chapter struct {
	ID            uuid.UUID `json:"id"`
	Number        int       `json:"number"`
	Title         string    `json:"title"`
	Topics        []string  `json:"topics,omitempty"`
	QuestionCount int       `json:"questionCount"`
}

// ChapterRef is nullable on a Question and carries confidence, so the UI can
// flag a low-confidence assignment instead of presenting a guess as fact.
type ChapterRef struct {
	ID         uuid.UUID `json:"id"`
	Number     int       `json:"number"`
	Title      string    `json:"title"`
	Confidence *float64  `json:"confidence,omitempty"`
}

type Exam struct {
	ID            uuid.UUID  `json:"id"`
	CourseID      uuid.UUID  `json:"courseId,omitempty"`
	Year          int        `json:"year"`
	Term          Term       `json:"term"`
	Title         string     `json:"title,omitempty"`
	DocumentID    *uuid.UUID `json:"documentId,omitempty"`
	HasSolutions  bool       `json:"hasSolutions"`
	QuestionCount int        `json:"questionCount"`
}

type ExamDetail struct {
	Exam
	Questions []Question `json:"questions"`
}

type Question struct {
	ID          uuid.UUID   `json:"id"`
	ExamID      *uuid.UUID  `json:"examId,omitempty"`
	Number      string      `json:"number"`
	Text        string      `json:"text"`
	Marks       *float64    `json:"marks,omitempty"`
	SourcePage  int         `json:"sourcePage"`
	Chapter     *ChapterRef `json:"chapter,omitempty"`
	Topic       string      `json:"topic,omitempty"`
	Year        *int        `json:"year,omitempty"`
	Term        Term        `json:"term,omitempty"`
	HasSolution bool        `json:"hasSolution"`
	// Type is the question's format (true/false, proof, derivation ...),
	// independent of its chapter.
	Type string `json:"type,omitempty"`
}

type QuestionDetail struct {
	Question
	DocumentID         *uuid.UUID `json:"documentId,omitempty"`
	SolutionText       string     `json:"solutionText,omitempty"`
	SolutionSourcePage *int       `json:"solutionSourcePage,omitempty"`
}

type SimilarQuestion struct {
	Question Question `json:"question"`
	Score    float64  `json:"score"`
}

type Citation struct {
	DocumentID     uuid.UUID    `json:"documentId"`
	DocumentTitle  string       `json:"documentTitle"`
	DocumentKind   DocumentKind `json:"documentKind,omitempty"`
	Page           int          `json:"page"`
	QuestionID     *uuid.UUID   `json:"questionId,omitempty"`
	QuestionNumber string       `json:"questionNumber,omitempty"`
	Label          string       `json:"label,omitempty"`
}

type Usage struct {
	PromptTokens     int    `json:"promptTokens,omitempty"`
	CompletionTokens int    `json:"completionTokens,omitempty"`
	RetrievedChunks  int    `json:"retrievedChunks"`
	LatencyMs        int64  `json:"latencyMs"`
	Provider         string `json:"provider,omitempty"`
	Model            string `json:"model,omitempty"`
}

type Answer struct {
	ConversationID uuid.UUID   `json:"conversationId"`
	MessageID      uuid.UUID   `json:"messageId"`
	Intent         QueryIntent `json:"intent"`
	Answer         string      `json:"answer"`
	Sources        []Citation  `json:"sources"`
	Questions      []Question  `json:"questions,omitempty"`
	Usage          *Usage      `json:"usage,omitempty"`
}

type RetrievalFilters struct {
	ChapterIDs         []uuid.UUID    `json:"chapterIds,omitempty"`
	ChapterNumbers     []int          `json:"chapterNumbers,omitempty"`
	YearFrom           *int           `json:"yearFrom,omitempty"`
	YearTo             *int           `json:"yearTo,omitempty"`
	DocumentKinds      []DocumentKind `json:"documentKinds,omitempty"`
	PreferOwnDocuments bool           `json:"preferOwnDocuments,omitempty"`
}

type SearchHit struct {
	ChunkID     uuid.UUID `json:"chunkId"`
	Text        string    `json:"text"`
	Score       float64   `json:"score"`
	DenseScore  float64   `json:"denseScore,omitempty"`
	SparseScore float64   `json:"sparseScore,omitempty"`
	Citation    Citation  `json:"citation"`
}

type Document struct {
	ID          uuid.UUID    `json:"id"`
	CourseID    uuid.UUID    `json:"courseId,omitempty"`
	Filename    string       `json:"filename"`
	Kind        DocumentKind `json:"kind"`
	Status      IngestStatus `json:"status"`
	Year        *int         `json:"year,omitempty"`
	Term        Term         `json:"term,omitempty"`
	PageCount   *int         `json:"pageCount,omitempty"`
	SizeBytes   int64        `json:"sizeBytes,omitempty"`
	ContentHash string       `json:"contentHash,omitempty"`
	UploadedAt  *time.Time   `json:"uploadedAt,omitempty"`
}

type DocumentStatus struct {
	DocumentID         uuid.UUID    `json:"documentId"`
	Status             IngestStatus `json:"status"`
	Progress           float64      `json:"progress"`
	StageDetail        string       `json:"stageDetail,omitempty"`
	QuestionsExtracted int          `json:"questionsExtracted"`
	Error              string       `json:"error,omitempty"`
}

type YearCount struct {
	Year          int `json:"year"`
	QuestionCount int `json:"questionCount"`
}

type ChapterFrequency struct {
	Chapter       ChapterRef  `json:"chapter"`
	QuestionCount int         `json:"questionCount"`
	Share         float64     `json:"share"`
	ExamCount     int         `json:"examCount"`
	ByYear        []YearCount `json:"byYear,omitempty"`
}

type ChapterFrequencyResponse struct {
	TotalQuestions int                `json:"totalQuestions"`
	Data           []ChapterFrequency `json:"data"`
}

// QuestionFilters carries the structured filters for the enumerate path.
type QuestionFilters struct {
	ChapterID     *uuid.UUID
	ChapterNumber *int
	YearFrom      *int
	YearTo        *int
	Term          *Term
	QuestionType  *string
	Query         string
	Sort          string
	Page          int
	PageSize      int
}
