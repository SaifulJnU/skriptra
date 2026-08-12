package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// Conversation history.
//
// Additive to the frozen v1 contract: new paths, no change to any existing
// response shape. `conversationId` was already documented on /ask and already
// returned; these endpoints are what make it mean something.

func (s *Server) listConversations(c *gin.Context) {
	courseID, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	page, size := pagination(c, 20, 100)

	convos, total, err := s.store.ListConversations(c, courseID, currentUser(c), page, size)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convos,
		"meta": domain.PageMeta{Page: page, PageSize: size, Total: total, TotalPages: totalPages(total, size)},
	})
}

func (s *Server) getConversation(c *gin.Context) {
	id, ok := uuidParam(c, "conversationId")
	if !ok {
		return
	}
	// Scoped by user in the query rather than fetched and then checked, so a
	// thread belonging to someone else is indistinguishable from one that does
	// not exist. Returning 403 would confirm the id is real.
	convo, err := s.store.GetConversation(c, id, currentUser(c))
	s.respond(c, err, convo)
}

func (s *Server) deleteConversation(c *gin.Context) {
	id, ok := uuidParam(c, "conversationId")
	if !ok {
		return
	}
	if err := s.store.DeleteConversation(c, id, currentUser(c)); err != nil {
		s.respond(c, err, nil)
		return
	}
	c.Status(http.StatusNoContent)
}
