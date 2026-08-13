package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
	"github.com/skriptra/skriptra/backend/internal/provider"
)

// solutionPrompt is stricter than the general answer prompt.
//
// A student revising from a generated solution will act on it in an exam, so
// the instruction to flag uncertainty rather than smooth over it is the whole
// point. A confidently wrong derivation is worse than an incomplete one,
// because the student cannot tell which they received.
const solutionPrompt = `You are helping a student work through a past exam question for a university course.

Produce a worked solution:
- Show the steps, not just the final result. State what is being used at each step.
- Use the notation and conventions visible in the retrieved course material.
- Where the retrieved material does not settle a step, say so explicitly rather than inventing a justification.
- If the question asks for a proof, give the structure of the proof and the key steps.
- Keep it concise. This is revision material, not a textbook chapter.

You are not the official mark scheme. Do not claim to be.`

// generateSolution produces a worked solution for a question that has none.
//
// It is deliberately NOT persisted to questions.solution_text. That column
// means "the official solution as published by the course", and a generated
// attempt is a different kind of object. Storing them in the same place would
// make an uploaded mark scheme indistinguishable from a model's guess a week
// later, which is precisely the confusion this product exists to prevent.
func (s *Server) generateSolution(c *gin.Context) {
	questionID, ok := uuidParam(c, "questionId")
	if !ok {
		return
	}

	question, err := s.store.GetQuestion(c, questionID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	if question.SolutionText != "" {
		fail(c, http.StatusConflict, "solution_exists",
			"This question already has an official solution.")
		return
	}

	var body struct {
		Stream bool `json:"stream"`
	}
	_ = c.ShouldBindJSON(&body)

	started := time.Now()

	courseID, err := s.store.CourseIDForQuestion(c, questionID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	// Retrieve against the question text itself. Lecture notes and solutions
	// from other years are the material most likely to contain the method.
	vectors, err := s.embedder.Embed(c, []string{question.Text})
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	filters := domain.RetrievalFilters{}
	if question.Chapter != nil {
		filters.ChapterNumbers = []int{question.Chapter.Number}
	}

	hits, err := s.store.HybridSearch(c, courseID, question.Text, vectors[0], filters, 6)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	sources := make([]domain.Citation, 0, len(hits))
	var ctxBuf strings.Builder
	for i, h := range hits {
		// The question itself will rank first against its own text. Including
		// it as context is circular and as a citation is meaningless.
		if h.Citation.QuestionID != nil && *h.Citation.QuestionID == questionID {
			continue
		}
		sources = append(sources, h.Citation)
		fmt.Fprintf(&ctxBuf, "[%d] (%s)\n%s\n\n", i+1, h.Citation.Label, h.Text)
	}

	w := newWriter(c, body.Stream)
	defer w.close()

	// With nothing retrieved, a generated solution would come entirely from the
	// model's own knowledge while being presented inside a course workspace.
	// That is the failure this product exists to prevent, so it declines
	// instead. It happens legitimately: a course whose only uploaded document
	// is the exam paper itself has no material to work from.
	if len(sources) == 0 {
		const text = "**Not enough course material to work from.**\n\nNothing beyond the question itself has been indexed for this chapter, so any solution would come from general knowledge rather than from this course. Upload the lecture notes or a solution sheet from another year, and try again."
		w.event("sources", gin.H{"sources": []domain.Citation{}})
		w.stream(text)
		w.done(domain.Answer{
			ConversationID: uuid.New(),
			MessageID:      uuid.New(),
			Intent:         domain.IntentExplain,
			Answer:         text,
			Sources:        []domain.Citation{},
			Usage:          &domain.Usage{LatencyMs: time.Since(started).Milliseconds()},
		})
		return
	}

	userMsg := fmt.Sprintf("Course material:\n\n%s\nExam question (%d marks):\n%s",
		ctxBuf.String(), int(marksOf(question)), question.Text)

	stream, err := s.llm.Stream(c, provider.GenerateRequest{
		Messages: []provider.Message{
			{Role: provider.RoleSystem, Content: solutionPrompt},
			{Role: provider.RoleUser, Content: userMsg},
		},
		Temperature: 0.3,
		MaxTokens:   1000,
	})
	if err != nil {
		// Same rule as /ask: no citations are announced unless generation
		// actually starts.
		w.fail(err)
		return
	}

	w.event("sources", gin.H{"sources": sources})

	var full strings.Builder
	usage := domain.Usage{RetrievedChunks: len(sources)}
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
	// Same rule as /ask: a refusal grounds nothing, so it keeps no citations
	// and the sentinel never reaches the user.
	if isRefusal(answer) {
		answer = "**Not enough course material to work from.**\n\nThe indexed material for this chapter does not contain the method this question needs. Upload the lecture notes or a solution sheet from another year."
		sources = []domain.Citation{}
		usage.RetrievedChunks = 0
	}

	w.done(domain.Answer{
		ConversationID: uuid.New(),
		MessageID:      uuid.New(),
		Intent:         domain.IntentExplain,
		Answer:         answer,
		Sources:        sources,
		Usage:          &usage,
	})
}

func marksOf(q *domain.QuestionDetail) float64 {
	if q.Marks == nil {
		return 0
	}
	return *q.Marks
}
