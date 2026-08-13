package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
	"github.com/skriptra/skriptra/backend/internal/ingest"
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
- If the passages do not contain the answer, reply with exactly INSUFFICIENT_CONTEXT on the first line and nothing else. Do not speculate and do not apologise at length.
- Refer to sources as [1], [2] matching the numbered passages.
- Be precise and concise. Prefer the notation the course itself uses.
- Mathematics may be written in plain text or LaTeX, whichever is clearer.`

// refusalToken is the sentinel the prompt asks for when the corpus cannot
// support an answer.
//
// An explicit signal rather than inferring refusal from the text. The earlier
// attempt treated "no [1] markers" as proof of an ungrounded answer, which was
// behaviourally sound but wrong in practice: a small model often produces a
// perfectly grounded answer without emitting markers, and its citations were
// then stripped incorrectly. Asking for one token is something even a 3B model
// follows reliably, and it does not depend on refusal phrasing or language.
const refusalToken = "INSUFFICIENT_CONTEXT"

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

	// Checked before any work: the course arrives in the body, so the path
	// guard cannot see it, and answering from a corpus the caller has no claim
	// on is the exact failure membership exists to prevent.
	if !s.mayAccessCourse(c, req.CourseID) {
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

	// The thread is resolved before any work is done, so an id that does not
	// belong to this user fails with 404 rather than after an answer has been
	// generated and paid for.
	user := currentUser(c)
	conversationID, err := s.store.StartOrContinueConversation(
		c, req.CourseID, user, req.ConversationID, req.Question)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	// History is read before the question is appended, so the current turn does
	// not appear in its own context.
	var history []domain.Message
	if req.ConversationID != nil {
		if history, err = s.store.RecentMessages(c, conversationID, historyTurns); err != nil {
			s.log.Warn("read conversation history", "conversation", conversationID, "error", err)
		}
	}

	// Recorded before generation. A question that failed to produce an answer is
	// still something the student asked, and losing it would make the thread
	// read as if they never asked.
	if _, err := s.store.AppendMessage(c, conversationID, domain.Message{
		Role:    domain.RoleUser,
		Content: req.Question,
	}); err != nil {
		s.log.Error("persist question", "conversation", conversationID, "error", err)
	}

	t := turn{conversationID: conversationID, history: history}

	w := newWriter(c, req.Stream)
	defer w.close()
	w.event("intent", gin.H{"intent": decision.Intent})

	switch decision.Intent {
	case domain.IntentAnalyse:
		s.answerAnalyse(c, w, t, req, started)
	case domain.IntentEnumerate:
		s.answerEnumerate(c, w, t, req, decision, started)
	default:
		s.answerExplain(c, w, t, req, decision, started)
	}
}

// historyTurns caps how much of a thread is replayed to the model.
//
// Six messages is three exchanges, which covers the follow-ups students
// actually ask ("why?", "show me an example of that") without letting an old
// thread crowd out the retrieved passages, which are what the answer must
// actually come from.
const historyTurns = 6

// turn is the conversation a request belongs to, threaded through the answer
// paths so each one persists its result the same way.
type turn struct {
	conversationID uuid.UUID
	history        []domain.Message
}

// finish records the assistant's turn and sends the terminal payload.
//
// Persistence failure is logged, not surfaced: the student has their answer,
// and replacing it with an error because a history row could not be written
// would be a worse outcome than a gap in the thread.
func (s *Server) finish(c *gin.Context, w *writer, t turn, a domain.Answer) {
	a.ConversationID = t.conversationID
	if a.Sources == nil {
		a.Sources = []domain.Citation{}
	}

	id, err := s.store.AppendMessage(c, t.conversationID, domain.Message{
		Role:    domain.RoleAssistant,
		Content: a.Answer,
		Intent:  a.Intent,
		Sources: a.Sources,
		Usage:   a.Usage,
	})
	if err != nil {
		s.log.Error("persist answer", "conversation", t.conversationID, "error", err)
		id = uuid.New()
	}
	a.MessageID = id

	w.done(a)
}

// answerAnalyse serves the aggregate path. No model is invoked at all, so the
// figures are exact and the response is immediate.
func (s *Server) answerAnalyse(c *gin.Context, w *writer, t turn, req askRequest, started time.Time) {
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
			"**Chapter %d, %s** is the most frequently tested material, accounting for %.0f%% of all %d indexed questions and appearing in %d separate exams.\n\nThe full distribution is on the Analytics tab. These figures come from the question index directly, so they are exact.",
			top.Chapter.Number, top.Chapter.Title, top.Share*100, freq.TotalQuestions, top.ExamCount)
	}

	w.event("sources", gin.H{"sources": []domain.Citation{}})
	w.stream(text)
	s.finish(c, w, t, domain.Answer{
		Intent:  domain.IntentAnalyse,
		Answer:  text,
		Sources: []domain.Citation{},
		Usage:   &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
	})
}

// answerEnumerate serves the exhaustive path, SQL, not retrieval. Top-k has no
// notion of completeness, so it is not asked for one.
func (s *Server) answerEnumerate(c *gin.Context, w *writer, t turn, req askRequest, d router.Decision, started time.Time) {
	f := domain.QuestionFilters{
		ChapterNumber:  d.ChapterNumber,
		ChapterNumbers: d.ChapterNumbers,
		YearFrom:       req.Filters.YearFrom,
		YearTo:         req.Filters.YearTo,
		Sort:           "newest",
		Page:           1,
		PageSize:       100,
	}
	if d.QuestionType != "" {
		qt := d.QuestionType
		f.QuestionType = &qt
	}
	if d.Marks != nil {
		f.MarksMin, f.MarksMax = d.Marks.Min, d.Marks.Max
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

	// Every filter that was applied is named back to the user. Silently
	// dropping one and answering a broader question is how "true/false
	// questions from chapter 2" came back as all 22 chapter 2 questions.
	var applied []string
	if d.QuestionType != "" {
		applied = append(applied, ingest.QuestionType(d.QuestionType).Label()+" questions")
	}
	if d.Marks != nil {
		applied = append(applied, d.Marks.Label)
	}
	switch len(d.ChapterNumbers) {
	case 0:
	case 1:
		applied = append(applied, fmt.Sprintf("Chapter %d", d.ChapterNumbers[0]))
	default:
		parts := make([]string, len(d.ChapterNumbers))
		for i, n := range d.ChapterNumbers {
			parts[i] = fmt.Sprintf("%d", n)
		}
		applied = append(applied, "Chapters "+strings.Join(parts, " and "))
	}
	if d.YearFrom != nil && d.YearTo != nil {
		if *d.YearFrom == *d.YearTo {
			applied = append(applied, fmt.Sprintf("%d", *d.YearFrom))
		} else {
			applied = append(applied, fmt.Sprintf("%d to %d", *d.YearFrom, *d.YearTo))
		}
	}
	scope := ""
	if len(applied) > 0 {
		scope = " matching " + strings.Join(applied, ", ")
	}

	text := fmt.Sprintf(
		"Found **%d question%s**%s across %d exam year%s.\n\nThey are listed below, newest first. This is an exhaustive result from the question index, not a sample.",
		total, plural(total), scope, len(years), plural(len(years)))

	if total == 0 {
		// Say what was searched for, so an empty result is informative rather
		// than a dead end. A corpus with no true/false questions is a fact
		// worth stating plainly.
		text = fmt.Sprintf("**No questions found%s.**\n\nThe filters applied were: %s. Nothing in the indexed papers matches all of them.",
			scope, strings.Join(applied, ", "))
		if len(applied) == 0 {
			text = "**No questions found.** Nothing has been indexed for this course yet."
		}
	}

	w.event("sources", gin.H{"sources": sources, "questions": questions})
	w.stream(text)
	s.finish(c, w, t, domain.Answer{
		Intent:    domain.IntentEnumerate,
		Answer:    text,
		Sources:   sources,
		Questions: questions,
		Usage:     &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
	})
}

// answerExplain is the only path that reaches a model, and it does so with
// retrieved passages and a refusal instruction.
func (s *Server) answerExplain(c *gin.Context, w *writer, t turn, req askRequest, d router.Decision, started time.Time) {
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
		s.finish(c, w, t, domain.Answer{
			Intent:  d.Intent,
			Answer:  text,
			Sources: []domain.Citation{},
			Usage:   &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
		})
		return
	}

	sources := make([]domain.Citation, len(hits))
	var ctxBuf strings.Builder
	for i, h := range hits {
		sources[i] = h.Citation
		fmt.Fprintf(&ctxBuf, "[%d] (%s)\n%s\n\n", i+1, h.Citation.Label, h.Text)
	}

	// Earlier turns go in ahead of the passages so a follow-up like "why is
	// that?" has a referent. Retrieval still runs on the current question
	// alone: letting an old turn steer the search is how a thread drifts onto
	// material the student stopped asking about three questions ago.
	msgs := make([]provider.Message, 0, len(t.history)+2)
	msgs = append(msgs, provider.Message{Role: provider.RoleSystem, Content: systemPrompt})
	msgs = append(msgs, priorTurns(t.history)...)
	msgs = append(msgs, provider.Message{Role: provider.RoleUser, Content: fmt.Sprintf(
		"Retrieved passages:\n\n%s\nQuestion: %s", ctxBuf.String(), req.Question)})

	genReq := provider.GenerateRequest{
		Messages:    msgs,
		Temperature: 0.2,
		MaxTokens:   800,
	}

	// Open the stream BEFORE announcing sources.
	//
	// Citations are a claim that an answer is grounded in those passages. If
	// generation never starts there is no answer, so the claim is false, and
	// showing a list of sources beside an error message is precisely the
	// dishonesty this product exists to avoid. An unreachable provider fails
	// here, before a single citation has been sent.
	stream, err := s.llm.Stream(c, genReq)
	if err != nil {
		w.fail(err)
		return
	}

	// Generation is underway, so the citations are now a claim that will be
	// backed by an answer. Sent before the tokens so they render alongside the
	// text as it streams.
	w.event("sources", gin.H{"sources": sources})

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

	answer := full.String()

	// An answer that cites nothing is not grounded in the passages, whatever
	// was retrieved for it. The commonest case is the model correctly refusing
	// because the corpus does not cover the question, and attaching a list of
	// citations to a refusal claims provenance for an answer that was never
	// given. The `done` payload is authoritative, so the client corrects any
	// citations it rendered while streaming.
	// A refusal grounds nothing, so it must not carry citations. Replaced with
	// wording a student can act on, since the sentinel itself means nothing to
	// them. Empty array rather than nil: the contract types `sources` as a
	// required array and a client trusting the spec should not meet a null.
	if isRefusal(answer) {
		answer = "**No indexed passage covers that.**\n\nNothing in the uploaded material for this course answers the question. Try rephrasing it, or upload the relevant paper or notes."
		sources = []domain.Citation{}
		usage.RetrievedChunks = 0
	}

	s.finish(c, w, t, domain.Answer{
		Intent:  d.Intent,
		Answer:  answer,
		Sources: sources,
		Usage:   &usage,
	})
}

// priorTurns renders stored history as model messages.
//
// Assistant turns are truncated hard. A full previous answer can run to eight
// hundred tokens, and three of them would consume more of the window than the
// passages the current answer has to be grounded in. The opening is enough to
// resolve a pronoun, which is all history is here to do.
//
// Refusals are dropped entirely: replaying "no indexed passage covers that"
// teaches the model that refusing is the expected shape of an answer, and it
// starts refusing questions the corpus does cover.
func priorTurns(history []domain.Message) []provider.Message {
	const maxAssistantChars = 600

	out := make([]provider.Message, 0, len(history))
	for _, m := range history {
		content := m.Content
		role := provider.RoleUser
		if m.Role == domain.RoleAssistant {
			if isRefusal(content) || strings.HasPrefix(content, "**No indexed passage") {
				continue
			}
			role = provider.RoleAssistant
			// Cut on a rune boundary. The corpus is bilingual and slicing a
			// German umlaut in half produces an invalid UTF-8 sequence that
			// some providers reject outright.
			if utf8.RuneCountInString(content) > maxAssistantChars {
				content = string([]rune(content)[:maxAssistantChars]) + "..."
			}
		}
		if strings.TrimSpace(content) == "" {
			continue
		}
		out = append(out, provider.Message{Role: role, Content: content})
	}
	return out
}

// reRefusal matches the sentinel tolerantly.
//
// Models normalise punctuation in tokens they are told to emit verbatim:
// llama3.2 returns "INSUFFICIENT CONTEXT" with a space where the prompt asked
// for an underscore. Requiring the exact string made the check silently fail,
// so separators and case are both treated as noise.
var reRefusal = regexp.MustCompile(`(?i)insufficient[\s_\-]*context`)

// isRefusal reports whether the model signalled that the passages were
// insufficient.
//
// The whole answer is scanned, not just the opening. The prompt asks for the
// sentinel on the first line and nothing else, but models routinely explain
// themselves first and emit it at the end. Checking only the head let a
// refusal through with its citations still attached, and printed the sentinel
// to the user.
//
// Scanning everything risks a false positive if a model discusses the phrase in
// passing. Given the phrase, that is far less likely than the failure it
// prevents.
func isRefusal(answer string) bool {
	return reRefusal.MatchString(answer)
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
	c         *gin.Context
	streaming bool
	flusher   http.Flusher
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
