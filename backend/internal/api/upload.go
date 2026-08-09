package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// uploadDocument accepts a file, stores it, and queues ingestion.
//
// It returns immediately rather than parsing inline. Ingesting a paper involves
// extraction, classification and embedding, which takes tens of seconds on a
// local model; holding an HTTP request open for that is a timeout waiting to
// happen and gives the user no progress.
func (s *Server) uploadDocument(c *gin.Context) {
	courseID, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}

	header, err := c.FormFile("file")
	if err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "A file is required.")
		return
	}
	if header.Size > int64(s.cfg.MaxUploadMB)*1024*1024 {
		fail(c, http.StatusRequestEntityTooLarge, "file_too_large",
			fmt.Sprintf("Maximum upload size is %d MB.", s.cfg.MaxUploadMB))
		return
	}

	kind := c.DefaultPostForm("kind", "exam")
	switch kind {
	case "exam", "solution", "notes", "textbook", "syllabus":
	default:
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "Unknown document kind.")
		return
	}

	src, err := header.Open()
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	defer src.Close()

	content, err := io.ReadAll(src)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	if !strings.HasPrefix(string(content), "%PDF") {
		fail(c, http.StatusUnsupportedMediaType, "unsupported_media",
			"Only PDF files are supported.")
		return
	}

	// Reject a scan up front rather than accepting it and indexing nothing.
	// The error names the missing capability so the operator knows the fix is
	// to deploy an OCR parser, not to re-upload.
	probe := ingest.ProbePDF(content)
	if !probe.HasTextLayer {
		fail(c, http.StatusUnprocessableEntity, "no_text_layer",
			"This PDF has no text layer, so it is a scan. OCR is not enabled in this deployment.")
		return
	}

	sum := sha256.Sum256(content)
	hash := hex.EncodeToString(sum[:])

	var year *int
	if raw := c.PostForm("year"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil {
			year = &v
		}
	}
	var term *string
	if raw := c.PostForm("term"); raw == "summer" || raw == "winter" {
		term = &raw
	}

	storageKey := filepath.Join(hash[:2], hash+".pdf")
	path := filepath.Join(s.cfg.StorageDir, storageKey)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		s.respond(c, err, nil)
		return
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		s.respond(c, err, nil)
		return
	}

	docID, existed, err := s.store.CreateDocument(c, courseID, header.Filename, kind,
		storageKey, hash, header.Size, year, term)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	if existed {
		// Deduplicated. 200 rather than 202 tells the client no new work was
		// queued, which is why the contract distinguishes the two.
		doc, err := s.store.DocumentStatus(c, docID)
		s.respond(c, err, gin.H{
			"id": docID, "status": doc.Status, "deduplicated": true,
			"message": "This document is already in the course.",
		})
		return
	}

	if err := s.queue.PublishIngest(c, docID, courseID, header.Filename, storageKey); err != nil {
		_ = s.store.FailDocument(c, docID, "queue: "+err.Error())
		s.respond(c, err, nil)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id": docID, "filename": header.Filename, "kind": kind,
		"status": "queued", "sizeBytes": header.Size, "contentHash": hash,
	})
}
