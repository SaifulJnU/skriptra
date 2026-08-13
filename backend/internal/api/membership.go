package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Course membership enforcement.
//
// Authentication answers who the caller is. This answers what they may see,
// and it is the half that actually keeps one student's corpus out of another's
// results.

// requireCourseMember guards every route addressed by :courseId.
//
// Middleware on the path parameter rather than a check repeated in each
// handler. Seventeen handlers each remembering to call the same three lines is
// seventeen chances to forget, and the one that forgets is the hole.
func (s *Server) requireCourseMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, ok := uuidParam(c, "courseId")
		if !ok {
			return
		}
		if !s.mayAccessCourse(c, courseID) {
			return
		}
		c.Next()
	}
}

// mayAccessCourse reports whether the caller belongs to a course, and writes
// the refusal itself when they do not.
//
// The refusal is 404, not 403. A 403 confirms the course exists, which lets
// anyone with a course id learn whether it is real. Not being a member and not
// existing are the same answer from outside.
func (s *Server) mayAccessCourse(c *gin.Context, courseID uuid.UUID) bool {
	member, err := s.store.IsCourseMember(c, courseID, currentUser(c))
	if err != nil {
		s.respond(c, err, nil)
		return false
	}
	if !member {
		fail(c, http.StatusNotFound, "not_found", "No such resource.")
		return false
	}
	return true
}

// courseGuard resolves the course behind a child identifier and checks
// membership against it, for the routes addressed by an exam, question or
// document id rather than by a course id.
func (s *Server) courseGuard(param string, lookup func(context.Context, uuid.UUID) (uuid.UUID, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := uuidParam(c, param)
		if !ok {
			return
		}
		courseID, err := lookup(c, id)
		if err != nil {
			s.respond(c, err, nil)
			return
		}
		if !s.mayAccessCourse(c, courseID) {
			return
		}
		c.Next()
	}
}

type createCourseRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code"`
	Institution string `json:"institution"`
	Language    string `json:"language"`
}

// createCourse makes a course and makes the caller its owner.
//
// The entry point for a new account: a fresh user belongs to nothing, so
// without this the product is empty for them and there is no way out of it.
func (s *Server) createCourse(c *gin.Context) {
	var req createCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" || len([]rune(name)) > 200 {
		fail(c, http.StatusUnprocessableEntity, "validation_failed",
			"Course name must be between 1 and 200 characters.")
		return
	}

	course, err := s.store.CreateCourse(c, currentUser(c), name, req.Code, req.Institution, req.Language)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusCreated, course)
}
