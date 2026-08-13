package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// serveDocument streams the original PDF.
//
// This endpoint is what makes a citation a citation. "2025 Summer Exam, Q4,
// Page 3" is a claim; being able to click it and land on page 3 of the actual
// paper is what makes the claim checkable. Without this the product asks to be
// trusted, which is the opposite of its argument.
//
// http.ServeContent handles range requests, which PDF viewers rely on to fetch
// a single page out of a large file rather than downloading all of it.
func (s *Server) serveDocument(c *gin.Context) {
	id, ok := uuidParam(c, "documentId")
	if !ok {
		return
	}

	doc, err := s.store.GetDocumentFile(c, id)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	// The storage key comes from the database, but it still gets cleaned and
	// re-anchored under the storage root. A path escaping the root would be a
	// file-read primitive, and defence in depth here costs one line.
	clean := filepath.Clean("/" + doc.StorageKey)
	path := filepath.Join(s.cfg.StorageDir, clean)
	root, err := filepath.Abs(s.cfg.StorageDir)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	abs, err := filepath.Abs(path)
	if err != nil || !strings.HasPrefix(abs, root) {
		fail(c, http.StatusNotFound, "not_found", "No such document.")
		return
	}

	f, err := os.Open(abs)
	if err != nil {
		// The row exists but the file does not, which means storage and the
		// database have diverged. Say so rather than reporting a generic 500.
		s.log.Error("stored file missing", "document", id, "path", abs, "err", err)
		fail(c, http.StatusNotFound, "file_missing",
			"This document's file is no longer on disk.")
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	c.Header("Content-Type", "application/pdf")
	// inline, so a citation opens in the browser's viewer at the cited page
	// rather than downloading. The filename is quoted because paper names
	// routinely contain spaces.
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%q", doc.Filename))
	// Content is immutable: the key is a hash of the bytes, so a changed file is
	// a different key.
	c.Header("Cache-Control", "private, max-age=86400, immutable")

	http.ServeContent(c.Writer, c.Request, doc.Filename, info.ModTime(), f)
}
