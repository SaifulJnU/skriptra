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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/auth"
	"github.com/skriptra/skriptra/backend/internal/cache"
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
	cache    cache.Cache
	issuer   *auth.Issuer
	log      *slog.Logger

	// Cached OCR reachability, see ocrAvailable.
	ocrHealthy   bool
	ocrCheckedAt time.Time
}

func New(cfg *config.Config, store *db.Store, llm provider.LLM, embedder provider.Embedder, q Publisher, c cache.Cache, issuer *auth.Issuer, log *slog.Logger) *Server {
	if c == nil {
		c = cache.NoOp{}
	}
	return &Server{cfg: cfg, store: store, llm: llm, embedder: embedder,
		queue: q, cache: c, issuer: issuer, log: log}
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
		// Open: liveness is what a load balancer polls, and the auth endpoints
		// are how a caller gets a token in the first place.
		v1.GET("/healthz", s.health)
		v1.POST("/auth/signup", s.signup)
		v1.POST("/auth/login", s.login)
		v1.POST("/auth/refresh", s.refresh)
		v1.POST("/auth/logout", s.logout)
	}

	// Everything below requires a valid access token. Guarding the group
	// rather than listing routes means a new endpoint is protected by default:
	// an allow-list eventually misses one, and a missed one is an open
	// endpoint nobody notices.
	authed := r.Group("/api/v1", s.requireAuth())
	{
		v1 := authed
		v1.GET("/me", s.me)
		v1.GET("/providers", s.providers)

		v1.GET("/courses", s.listCourses)
		v1.POST("/courses", s.createCourse)

		// Membership is enforced on the path parameter, not inside each
		// handler. One guard cannot be forgotten; seventeen repetitions of it
		// eventually are.
		course := v1.Group("/courses/:courseId", s.requireCourseMember())
		{
			course.GET("", s.getCourse)
			course.GET("/chapters", s.listChapters)
			course.GET("/exams", s.listExams)
			course.GET("/questions", s.listQuestions)
			course.GET("/documents", s.listDocuments)
			course.POST("/documents", s.uploadDocument)
			course.GET("/analytics/chapter-frequency", s.chapterFrequency)
			course.GET("/conversations", s.listConversations)
		}

		// Already scoped by user in the query: a thread belonging to someone
		// else is indistinguishable from one that does not exist.
		v1.GET("/conversations/:conversationId", s.getConversation)
		v1.DELETE("/conversations/:conversationId", s.deleteConversation)

		// Addressed by a child id, so the course is resolved first and then
		// checked. Without this an id leaked from another user's course would
		// read straight through.
		exam := v1.Group("/exams/:examId", s.courseGuard("examId", s.store.CourseIDForExam))
		exam.GET("", s.getExam)

		question := v1.Group("/questions/:questionId", s.courseGuard("questionId", s.store.CourseIDForQuestion))
		{
			question.GET("", s.getQuestion)
			question.GET("/similar", s.similarQuestions)
			question.POST("/solution", s.generateSolution)
		}

		document := v1.Group("/documents/:documentId", s.courseGuard("documentId", s.store.CourseIDForDocument))
		{
			document.GET("/status", s.documentStatus)
			document.GET("/file", s.serveDocument)
		}

		// Both carry the course in the body, so they check it themselves.
		v1.POST("/search", s.search)
		v1.POST("/ask", s.ask)
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
			// The refresh cookie only travels if the browser is told to send
			// credentials cross-origin, which the dev server is.
			c.Header("Access-Control-Allow-Credentials", "true")
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
