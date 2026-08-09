package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
	"github.com/skriptra/skriptra/backend/internal/provider"
	"github.com/skriptra/skriptra/backend/internal/router"
)

func isErr(err, target error) bool { return errors.Is(err, target) }

type askRequest struct {
	CourseID       uuid.UUID               `json:"courseId" binding:"required"`
	Question       string                  `json:"question" binding:"required"`
	ConversationID *uuid.UUID              `json:"conversationId"`
	Filters        domain.RetrievalFilters `json:"filters"`
	Stream         bool                    `json:"stream"`
}

// systemPrompt is deliberately strict about two things: cite what you used, and
// refuse when the passages do not contain the answer.
//
// An exam-revision tool that invents a plausible answer is worse than one that
// says nothing, because the student cannot tell the difference and finds out in
// the exam.
const systemPrompt = `You are a study assistant for a university course. You answer strictly from the retrieved passages supplied below.

Rules:
- Use only the passages. Do not add material from general knowledge.
- If the passages do not contain the answer, say so plainly and stop. Do not speculate.
- Refer to sources as [1], [2] matching the numbered passages.
- Be precise and concise. Prefer the notation the course itself uses.
- Mathematics may be written in plain text or LaTeX, whichever is clearer.`

func (s *Server) ask(c *gin.Context) {
	var req askRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}
	if len(req.Question) > 4000 {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "Question is too long.")
		return
	}

	started := time.Now()

	// The taxonomy is what lets "the maximum likelihood chapter" resolve to a
	// filter rather than a similarity search.
	chapters, err := s.store.ListChapters(c, req.CourseID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	rc := make([]router.Chapter, len(chapters))
	for i, ch := range chapters {
		rc[i] = router.Chapter{Number: ch.Number, Title: ch.Title, Topics: ch.Topics}
	}

	decision := router.Route(req.Question, rc, time.Now().Year())
	if decision.ChapterNumber != nil && len(req.Filters.ChapterNumbers) == 0 {
		req.Filters.ChapterNumbers = []int{*decision.ChapterNumber}
	}
	if decision.YearFrom != nil && req.Filters.YearFrom == nil {
		req.Filters.YearFrom = decision.YearFrom
	}
	if decision.YearTo != nil && req.Filters.YearTo == nil {
		req.Filters.YearTo = decision.YearTo
	}

	w := newWriter(c, req.Stream)
	defer w.close()
	w.event("intent", gin.H{"intent": decision.Intent})

	switch decision.Intent {
	case domain.IntentAnalyse:
		s.answerAnalyse(c, w, req, started)
	case domain.IntentEnumerate:
		s.answerEnumerate(c, w, req, decision, started)
	default:
		s.answerExplain(c, w, req, decision, started)
	}
}

// answerAnalyse serves the aggregate path. No model is invoked at all, so the
// figures are exact and the response is immediate.
func (s *Server) answerAnalyse(c *gin.Context, w *writer, req askRequest, started time.Time) {
	freq, err := s.store.ChapterFrequency(c, req.CourseID, req.Filters.YearFrom, req.Filters.YearTo)
	if err != nil {
		w.fail(err)
		return
	}

	top := domain.ChapterFrequency{}
	for _, f := range freq.Data {
		if f.QuestionCount > top.QuestionCount {
			top = f
		}
	}

	var text string
	if top.QuestionCount == 0 {
		text = "No questions have been classified into chapters for this course yet."
	} else {
		text = fmt.Sprintf(
			"**Chapter %d — %s** is the most frequently tested material, accounting for %.0f%% of all %d indexed questions and appearing in %d separate exams.\n\nThe full distribution is on the Analytics tab. These figures come from the question index directly, so they are exact.",
			top.Chapter.Number, top.Chapter.Title, top.Share*100, freq.TotalQuestions, top.ExamCount)
	}

	w.event("sources", gin.H{"sources": []domain.Citation{}})
	w.stream(text)
	w.done(domain.Answer{
		ConversationID: uuid.New(),
		MessageID:      uuid.New(),
		Intent:         domain.IntentAnalyse,
		Answer:         text,
		Sources:        []domain.Citation{},
		Usage:          &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
	})
}

// answerEnumerate serves the exhaustive path — SQL, not retrieval. Top-k has no
// notion of completeness, so it is not asked for one.
func (s *Server) answerEnumerate(c *gin.Context, w *writer, req askRequest, d router.Decision, started time.Time) {
	f := domain.QuestionFilters{
		ChapterNumber: d.ChapterNumber,
		YearFrom:      req.Filters.YearFrom,
		YearTo:        req.Filters.YearTo,
		Sort:          "newest",
		Page:          1,
		PageSize:      100,
	}
	questions, total, err := s.store.ListQuestions(c, req.CourseID, f, nil)
	if err != nil {
		w.fail(err)
		return
	}

	years := map[int]struct{}{}
	sources := make([]domain.Citation, 0, 5)
	for i, q := range questions {
		if q.Year != nil {
			years[*q.Year] = struct{}{}
		}
		if i < 5 {
			sources = append(sources, citationFor(q))
		}
	}

	scope := ""
	if d.ChapterNumber != nil {
		scope = fmt.Sprintf(" in Chapter %d", *d.ChapterNumber)
	}
	text := fmt.Sprintf(
		"Found **%d question%s**%s across %d exam year%s.\n\nThey are listed below, newest first. This is an exhaustive result from the question index, not a sample.",
		total, plural(total), scope, len(years), plural(len(years)))
	if total == 0 {
		text = "No questions match those filters." + scope
	}

	w.event("sources", gin.H{"sources": sources, "questions": questions})
	w.stream(text)
	w.done(domain.Answer{
		ConversationID: uuid.New(),
		MessageID:      uuid.New(),
		Intent:         domain.IntentEnumerate,
		Answer:         text,
		Sources:        sources,
		Questions:      questions,
		Usage:          &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
	})
}

// answerExplain is the only path that reaches a model, and it does so with
// retrieved passages and a refusal instruction.
func (s *Server) answerExplain(c *gin.Context, w *writer, req askRequest, d router.Decision, started time.Time) {
	vectors, err := s.embedder.Embed(c, []string{req.Question})
	if err != nil {
		w.fail(err)
		return
	}

	hits, err := s.store.HybridSearch(c, req.CourseID, req.Question, vectors[0], req.Filters, 8)
	if err != nil {
		w.fail(err)
		return
	}

	// Nothing retrieved means nothing to ground an answer in. Refuse here
	// rather than letting the model fill the gap from general knowledge.
	if len(hits) == 0 {
		const text = "**No indexed passage covers that.**\n\nNothing in the uploaded material for this course matches the question. Try rephrasing it, or upload the relevant paper or notes."
		w.event("sources", gin.H{"sources": []domain.Citation{}})
		w.stream(text)
		w.done(domain.Answer{
			ConversationID: uuid.New(),
			MessageID:      uuid.New(),
			Intent:         d.Intent,
			Answer:         text,
			Sources:        []domain.Citation{},
			Usage:          &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
		})
		return
	}

	sources := make([]domain.Citation, len(hits))
	var ctxBuf strings.Builder
	for i, h := range hits {
		sources[i] = h.Citation
		fmt.Fprintf(&ctxBuf, "[%d] (%s)\n%s\n\n", i+1, h.Citation.Label, h.Text)
	}
	w.event("sources", gin.H{"sources": sources})

	genReq := provider.GenerateRequest{
		Messages: []provider.Message{
			{Role: provider.RoleSystem, Content: systemPrompt},
			{Role: provider.RoleUser, Content: fmt.Sprintf(
				"Retrieved passages:\n\n%s\nQuestion: %s", ctxBuf.String(), req.Question)},
		},
		Temperature: 0.2,
		MaxTokens:   800,
	}

	stream, err := s.llm.Stream(c, genReq)
	if err != nil {
		w.fail(err)
		return
	}

	var full strings.Builder
	usage := domain.Usage{RetrievedChunks: len(hits)}
	for chunk := range stream {
		if chunk.Err != nil {
			w.fail(chunk.Err)
			return
		}
		if chunk.Text != "" {
			full.WriteString(chunk.Text)
			w.event("token", gin.H{"text": chunk.Text})
		}
		if chunk.Done {
			usage.PromptTokens = chunk.Usage.PromptTokens
			usage.CompletionTokens = chunk.Usage.CompletionTokens
			usage.Provider = chunk.Usage.Provider
			usage.Model = chunk.Usage.Model
		}
	}
	usage.LatencyMs = time.Since(started).Milliseconds()

	w.done(domain.Answer{
		ConversationID: uuid.New(),
		MessageID:      uuid.New(),
		Intent:         d.Intent,
		Answer:         full.String(),
		Sources:        sources,
		Usage:          &usage,
	})
}

func citationFor(q domain.Question) domain.Citation {
	title := "Exam"
	if q.Year != nil {
		season := ""
		switch q.Term {
		case domain.TermSummer:
			season = " Summer"
		case domain.TermWinter:
			season = " Winter"
		}
		title = fmt.Sprintf("%d%s Exam", *q.Year, season)
	}
	id := q.ID
	return domain.Citation{
		DocumentTitle:  title,
		Page:           q.SourcePage,
		QuestionID:     &id,
		QuestionNumber: q.Number,
		Label:          fmt.Sprintf("%s · Q%s · Page %d", title, q.Number, q.SourcePage),
	}
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

// ---------------------------------------------------------------- writer --

// writer renders either SSE frames or a single JSON body, so each answer path
// is written once instead of twice.
type writer struct {
	c        *gin.Context
	streaming bool
	flusher  http.Flusher
}

func newWriter(c *gin.Context, streaming bool) *writer {
	w := &writer{c: c, streaming: streaming}
	if streaming {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		// Without this, nginx buffers the whole response and the client sees
		// the answer appear at once after a long pause.
		c.Header("X-Accel-Buffering", "no")
		w.flusher, _ = c.Writer.(http.Flusher)
	}
	return w
}

func (w *writer) event(name string, payload any) {
	if !w.streaming {
		return
	}
	w.c.SSEvent(name, payload)
	if w.flusher != nil {
		w.flusher.Flush()
	}
}

// stream emits pre-computed text as tokens so a SQL-backed answer arrives the
// same way a generated one does and the client needs only one code path.
func (w *writer) stream(text string) {
	if !w.streaming {
		return
	}
	for _, word := range splitKeepingSpace(text) {
		select {
		case <-w.c.Request.Context().Done():
			return
		default:
		}
		w.event("token", gin.H{"text": word})
	}
}

func (w *writer) done(a domain.Answer) {
	if !w.streaming {
		w.c.JSON(http.StatusOK, a)
		return
	}
	w.event("done", a)
}

func (w *writer) fail(err error) {
	if !w.streaming {
		if e := mapProviderError(err); e != nil {
			w.c.JSON(e.status, gin.H{"error": errBody{Code: e.code, Message: e.message}})
			return
		}
		w.c.JSON(http.StatusInternalServerError, gin.H{
			"error": errBody{Code: "internal", Message: "Something went wrong."},
		})
		return
	}
	// The status line is already sent, so a mid-stream failure has to be an
	// event the client can render rather than an HTTP status.
	msg := "Something went wrong."
	if e := mapProviderError(err); e != nil {
		msg = e.message
	}
	w.event("error", gin.H{"message": msg})
}

func (w *writer) close() {
	if w.streaming && w.flusher != nil {
		w.flusher.Flush()
	}
}

type mappedError struct {
	status  int
	code    string
	message string
}

func mapProviderError(err error) *mappedError {
	switch {
	case err == nil:
		return nil
	case isErr(err, provider.ErrUnavailable):
		return &mappedError{http.StatusServiceUnavailable, "provider_unavailable",
			"The model provider is not reachable. Is it running?"}
	case isErr(err, provider.ErrRateLimited):
		return &mappedError{http.StatusTooManyRequests, "rate_limited",
			"The model provider is rate limiting requests."}
	case isErr(err, provider.ErrContextTooLong):
		return &mappedError{http.StatusUnprocessableEntity, "context_too_long",
			"Too much context for this model. Narrow the question with a chapter or year filter."}
	case isErr(err, context.Canceled), isErr(err, io.EOF):
		return &mappedError{499, "client_closed", "Request cancelled."}
	}
	return nil
}

func splitKeepingSpace(s string) []string {
	fields := strings.SplitAfter(s, " ")
	out := make([]string, 0, len(fields))
	for _, f := range fields {
		if f != "" {
			out = append(out, f)
		}
	}
	return out
}
