package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// ocrAvailable reports whether the OCR sidecar can be reached.
//
// Cached for a short window: this runs on every upload of a scan, and a health
// check per upload would add a round trip to a path the user is waiting on.
// Short enough that starting the service does not require restarting the API.
func (s *Server) ocrAvailable(ctx context.Context) bool {
	if s.cfg.OCRURL == "" {
		return false
	}
	if time.Since(s.ocrCheckedAt) < 30*time.Second {
		return s.ocrHealthy
	}
	parser := ingest.NewOCRParser(s.cfg.OCRURL)
	if parser == nil {
		return false
	}
	s.ocrHealthy = parser.Healthy(ctx)
	s.ocrCheckedAt = time.Now()
	return s.ocrHealthy
}

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

	// Format is decided by content, not by the filename. A phone that names a
	// photo "scan.pdf" is still sending an image.
	format, probe := ingest.ProbeDocument(content)
	if format == ingest.FormatUnknown {
		fail(c, http.StatusUnsupportedMediaType, "unsupported_media",
			fmt.Sprintf("Unsupported file type. Supported: %s.",
				strings.Join(ingest.SupportedFormats(), ", ")))
		return
	}

	// Anything without a text layer needs OCR. Rejecting it here, rather than
	// accepting it and having a worker fail out of sight, means the message
	// names the missing capability while the user is still looking at the
	// dialog.
	if !probe.HasTextLayer && !s.ocrAvailable(c) {
		msg := "This PDF has no text layer, so it is a scan, and OCR is not enabled in this deployment."
		if format == ingest.FormatImage {
			msg = "Reading a photo needs OCR, which is not enabled in this deployment."
		}
		fail(c, http.StatusUnprocessableEntity, "ocr_unavailable", msg)
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
