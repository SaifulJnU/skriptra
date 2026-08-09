// Package api is the HTTP layer.
//
// It does three things and nothing else: parse and validate input, call a
// service, and render the contract's response envelope. No SQL, no prompts, no
// business rules.
package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/config"
	"github.com/skriptra/skriptra/backend/internal/db"
	"github.com/skriptra/skriptra/backend/internal/provider"
)

// Publisher is the queue capability the API needs. Declared here rather than
// taking a concrete *queue.Queue so uploads can be tested without NATS.
type Publisher interface {
	PublishIngest(ctx context.Context, documentID, courseID uuid.UUID, filename, storageKey string) error
}

type Server struct {
	cfg      *config.Config
	store    *db.Store
	llm      provider.LLM
	embedder provider.Embedder
	queue    Publisher
	log      *slog.Logger
}

func New(cfg *config.Config, store *db.Store, llm provider.LLM, embedder provider.Embedder, q Publisher, log *slog.Logger) *Server {
	return &Server{cfg: cfg, store: store, llm: llm, embedder: embedder, queue: q, log: log}
}

func (s *Server) Routes() http.Handler {
	if s.cfg.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery(), s.requestLogger(), s.cors())

	// Everything under /api/v1 from day one. A v2 can be introduced later
	// without breaking a shipped mobile client.
	v1 := r.Group("/api/v1")
	{
		v1.GET("/healthz", s.health)
		v1.GET("/me", s.me)
		v1.GET("/providers", s.providers)

		v1.GET("/courses", s.listCourses)
		v1.GET("/courses/:courseId", s.getCourse)
		v1.GET("/courses/:courseId/chapters", s.listChapters)
		v1.GET("/courses/:courseId/exams", s.listExams)
		v1.GET("/courses/:courseId/questions", s.listQuestions)
		v1.GET("/courses/:courseId/documents", s.listDocuments)
		v1.POST("/courses/:courseId/documents", s.uploadDocument)
		v1.GET("/courses/:courseId/analytics/chapter-frequency", s.chapterFrequency)

		v1.GET("/exams/:examId", s.getExam)
		v1.GET("/questions/:questionId", s.getQuestion)
		v1.GET("/questions/:questionId/similar", s.similarQuestions)

		v1.POST("/search", s.search)
		v1.POST("/ask", s.ask)

		v1.GET("/documents/:documentId/status", s.documentStatus)
	}
	return r
}

// ------------------------------------------------------------- middleware --

func (s *Server) requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		lvl := slog.LevelInfo
		if c.Writer.Status() >= 500 {
			lvl = slog.LevelError
		}
		s.log.Log(c, lvl, "request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
		)
	}
}

func (s *Server) cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Development only: the Vite dev server runs on a different origin. In
		// production the frontend is served same-origin, so no allowance is
		// granted and this middleware is a no-op.
		if s.cfg.Env == "development" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// ----------------------------------------------------------------- errors --

type errBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details []any  `json:"details,omitempty"`
}

func fail(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, gin.H{"error": errBody{Code: code, Message: message}})
}

// respond maps a store error onto the contract's error envelope so no handler
// repeats this mapping.
func (s *Server) respond(c *gin.Context, err error, payload any) {
	switch {
	case err == nil:
		c.JSON(http.StatusOK, payload)
	case errors.Is(err, db.ErrNotFound):
		fail(c, http.StatusNotFound, "not_found", "No such resource.")
	case errors.Is(err, provider.ErrUnavailable):
		fail(c, http.StatusServiceUnavailable, "provider_unavailable",
			"The configured model provider is not reachable.")
	case errors.Is(err, provider.ErrRateLimited):
		fail(c, http.StatusTooManyRequests, "rate_limited", "The model provider is rate limiting requests.")
	default:
		s.log.Error("request failed", "path", c.Request.URL.Path, "err", err)
		fail(c, http.StatusInternalServerError, "internal", "Something went wrong.")
	}
}

// ------------------------------------------------------------------ input --

func uuidParam(c *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(name))
	if err != nil {
		fail(c, http.StatusBadRequest, "invalid_id", "That is not a valid identifier.")
		return uuid.Nil, false
	}
	return id, true
}

func intQuery(c *gin.Context, name string) *int {
	raw := c.Query(name)
	if raw == "" {
		return nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return nil
	}
	return &n
}

// pagination clamps page size so a client cannot ask for the whole corpus in
// one request.
func pagination(c *gin.Context, defaultSize, maxSize int) (page, size int) {
	page, size = 1, defaultSize
	if p := intQuery(c, "page"); p != nil && *p > 0 {
		page = *p
	}
	if s := intQuery(c, "pageSize"); s != nil && *s > 0 {
		size = *s
	}
	if size > maxSize {
		size = maxSize
	}
	return page, size
}

func totalPages(total, size int) int {
	if size <= 0 {
		return 1
	}
	n := total / size
	if total%size != 0 {
		n++
	}
	if n == 0 {
		n = 1
	}
	return n
}
