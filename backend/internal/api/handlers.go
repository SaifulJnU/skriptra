package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/cache"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

func (s *Server) health(c *gin.Context) {
	if err := s.store.Ping(c); err != nil {
		fail(c, http.StatusServiceUnavailable, "database_unavailable", "Database is not reachable.")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// me returns the single-user stub. The authorization *checks* are the point of
// having this now; the identity behind them gets real in phase 2.
func (s *Server) me(c *gin.Context) {
	c.JSON(http.StatusOK, domain.User{
		ID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		DisplayName: "Saiful",
	})
}

// providers surfaces what is actually configured, so the UI can show it and a
// self-hosting user can confirm at a glance that nothing leaves their machine.
func (s *Server) providers(c *gin.Context) {
	llm := s.llm.Info()
	emb := s.embedder.Info()
	c.JSON(http.StatusOK, domain.Providers{
		LLM: domain.ProviderInfo{
			Provider: llm.Provider, Model: llm.Model, Local: llm.Local,
		},
		Embedding: domain.ProviderInfo{
			Provider: emb.Provider, Model: emb.Model, Local: emb.Local, Dimensions: emb.Dimensions,
		},
	})
}

func (s *Server) listCourses(c *gin.Context) {
	page, size := pagination(c, 20, 100)
	courses, total, err := s.store.ListCourses(c, page, size)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, domain.Paged[domain.Course]{
		Data: courses,
		Meta: domain.PageMeta{Page: page, PageSize: size, Total: total, TotalPages: totalPages(total, size)},
	})
}

func (s *Server) getCourse(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	course, err := s.store.GetCourse(c, id)
	s.respond(c, err, course)
}

func (s *Server) listChapters(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	chapters, err := s.store.ListChapters(c, id)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": chapters})
}

func (s *Server) listExams(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	page, size := pagination(c, 50, 100)
	exams, total, err := s.store.ListExams(c, id, page, size)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, domain.Paged[domain.Exam]{
		Data: exams,
		Meta: domain.PageMeta{Page: page, PageSize: size, Total: total, TotalPages: totalPages(total, size)},
	})
}

func (s *Server) getExam(c *gin.Context) {
	id, ok := uuidParam(c, "examId")
	if !ok {
		return
	}
	exam, err := s.store.GetExam(c, id)
	s.respond(c, err, exam)
}

// listQuestions is the enumerate path. Filters are structured, applied in SQL,
// and the result is exhaustive within the requested page, not a top-k sample.
func (s *Server) listQuestions(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	page, size := pagination(c, 20, 100)

	f := domain.QuestionFilters{
		ChapterNumber: intQuery(c, "chapterNumber"),
		YearFrom:      intQuery(c, "yearFrom"),
		YearTo:        intQuery(c, "yearTo"),
		Query:         c.Query("q"),
		Sort:          c.DefaultQuery("sort", "newest"),
		Page:          page,
		PageSize:      size,
	}
	if raw := c.Query("chapterId"); raw != "" {
		if chID, err := uuid.Parse(raw); err == nil {
			f.ChapterID = &chID
		}
	}
	if raw := c.Query("term"); raw != "" {
		t := domain.Term(raw)
		f.Term = &t
	}

	questions, total, err := s.store.ListQuestions(c, id, f, nil)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, domain.Paged[domain.Question]{
		Data: questions,
		Meta: domain.PageMeta{Page: page, PageSize: size, Total: total, TotalPages: totalPages(total, size)},
	})
}

func (s *Server) getQuestion(c *gin.Context) {
	id, ok := uuidParam(c, "questionId")
	if !ok {
		return
	}
	q, err := s.store.GetQuestion(c, id)
	s.respond(c, err, q)
}

func (s *Server) similarQuestions(c *gin.Context) {
	id, ok := uuidParam(c, "questionId")
	if !ok {
		return
	}
	limit := 10
	if l := intQuery(c, "limit"); l != nil && *l > 0 && *l <= 50 {
		limit = *l
	}
	minScore := 0.5
	if raw := c.Query("minScore"); raw != "" {
		if v, err := strconv.ParseFloat(raw, 64); err == nil && v >= 0 && v <= 1 {
			minScore = v
		}
	}

	similar, err := s.store.SimilarQuestions(c, id, limit, minScore)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": similar})
}

// chapterFrequency is cached because it is a full aggregate over every question
// in a course, it is read on two screens, and it only changes when a document
// finishes indexing. The worker drops the course scope at that point, so the
// TTL is a backstop rather than the correctness mechanism.
func (s *Server) chapterFrequency(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	yearFrom, yearTo := intQuery(c, "yearFrom"), intQuery(c, "yearTo")

	// The year range is part of the key: two different ranges are two
	// different aggregates, and sharing one entry would serve the wrong one.
	variant := fmt.Sprintf("chapter-freq:%s:%s", intPtr(yearFrom), intPtr(yearTo))
	key := cache.AnalyticsKey(id.String(), variant)

	if raw, hit := s.cache.Get(c, key); hit {
		var cached domain.ChapterFrequencyResponse
		if err := json.Unmarshal(raw, &cached); err == nil {
			c.JSON(http.StatusOK, cached)
			return
		}
		// A stale encoding is treated as a miss rather than an error.
	}

	freq, err := s.store.ChapterFrequency(c, id, yearFrom, yearTo)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	if raw, err := json.Marshal(freq); err == nil {
		s.cache.Set(c, key, raw, 6*time.Hour)
	}
	c.JSON(http.StatusOK, freq)
}

func intPtr(v *int) string {
	if v == nil {
		return "all"
	}
	return strconv.Itoa(*v)
}

func (s *Server) listDocuments(c *gin.Context) {
	id, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	page, size := pagination(c, 50, 100)
	docs, total, err := s.store.ListDocuments(c, id, page, size)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, domain.Paged[domain.Document]{
		Data: docs,
		Meta: domain.PageMeta{Page: page, PageSize: size, Total: total, TotalPages: totalPages(total, size)},
	})
}

func (s *Server) documentStatus(c *gin.Context) {
	id, ok := uuidParam(c, "documentId")
	if !ok {
		return
	}
	st, err := s.store.DocumentStatus(c, id)
	s.respond(c, err, st)
}

type searchRequest struct {
	CourseID uuid.UUID               `json:"courseId" binding:"required"`
	Query    string                  `json:"query"    binding:"required"`
	Filters  domain.RetrievalFilters `json:"filters"`
	Limit    int                     `json:"limit"`
}

func (s *Server) search(c *gin.Context) {
	var req searchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}
	if req.Limit <= 0 || req.Limit > 50 {
		req.Limit = 10
	}

	vectors, err := s.embedder.Embed(c, []string{req.Query})
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	hits, err := s.store.HybridSearch(c, req.CourseID, req.Query, vectors[0], req.Filters, req.Limit)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hits})
}
