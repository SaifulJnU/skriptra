package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// Course setup: defining what the course actually covers.
//
// A taxonomy is not optional decoration. Chapters are the primary navigation,
// the filter that makes "give me the Chapter 3 questions" answerable, and the
// axis the analytics are computed over. Without one a course can hold a hundred
// indexed questions and answer almost nothing useful about them.
//
// The source is the course's own contents page rather than the questions,
// because a taxonomy inferred from what was examined would make the analytics
// circular: "which chapters are tested most often" cannot be answered by a list
// of chapters derived from what was tested.

const (
	maxChapters = 40

	// outlinePageLimit is how far into a document a chapter list is looked
	// for. Twelve covers a front cover, a title page and a contents page that
	// runs to several pages, without paying for the body.
	outlinePageLimit = 12
)

// extractOutline proposes a taxonomy from an uploaded syllabus or contents
// page. It saves nothing.
//
// Proposing rather than saving is deliberate. A contents page is messy, and
// only the person who owns the course knows whether a line is a chapter or a
// section heading inside one. The review step is also where the topic lists get
// the exam vocabulary that most of the classification quality comes from.
func (s *Server) extractOutline(c *gin.Context) {
	courseID, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}
	_ = courseID // membership is already enforced by the route guard

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "No file was uploaded.")
		return
	}
	defer file.Close()

	// Read it all: the probe needs to see the magic bytes and the parser needs
	// the whole file, and a syllabus is small. The upload limit still applies
	// through Gin's MaxMultipartMemory.
	buf, err := io.ReadAll(io.LimitReader(file, int64(s.cfg.MaxUploadMB)<<20))
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	format, probe := ingest.ProbeDocument(buf)
	if format == ingest.FormatUnknown {
		fail(c, http.StatusUnsupportedMediaType, "unsupported_media",
			"That file type is not supported. Accepted: "+strings.Join(ingest.SupportedFormats(), ", "))
		return
	}

	// Only the front of the document. A contents page is at the front by
	// definition, and reading a whole textbook to find it is waste that has
	// already killed the OCR container once.
	parsed, err := s.parsers.ParseFirstPages(c, bytes.NewReader(buf), header.Filename, probe, outlinePageLimit)
	if err != nil {
		fail(c, http.StatusUnprocessableEntity, "parse_failed",
			"That file could not be read: "+err.Error())
		return
	}

	outline, err := ingest.ParseOutline(c, parsed.Pages, s.llm)
	if err != nil {
		// Not a server fault: the user uploaded something that is not a
		// contents page. Saying so plainly beats returning an empty list that
		// looks like a taxonomy with no chapters in it.
		fail(c, http.StatusUnprocessableEntity, "no_outline_found",
			"No chapter list could be read from that file. A syllabus or a book's contents page works best.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chapters": proposedChapters(outline.Chapters),
		"source":   outline.Source,
		"filename": header.Filename,
	})
}

// proposedChapter is the wire shape of a chapter that does not exist yet.
//
// The internal ingest type is not serialised directly: it carries no JSON tags,
// so it went out as Number and Title while every other field in the contract is
// lowerCamelCase, and a client written against the spec could not read it.
type proposedChapter struct {
	Number int      `json:"number"`
	Title  string   `json:"title"`
	Topics []string `json:"topics"`
}

func proposedChapters(chapters []ingest.Chapter) []proposedChapter {
	out := make([]proposedChapter, 0, len(chapters))
	for _, ch := range chapters {
		topics := ch.Topics
		if topics == nil {
			// The contract types topics as an array; a client trusting it
			// should not meet a null.
			topics = []string{}
		}
		out = append(out, proposedChapter{Number: ch.Number, Title: ch.Title, Topics: topics})
	}
	return out
}

type saveChaptersRequest struct {
	Chapters []struct {
		Number int      `json:"number"`
		Title  string   `json:"title"`
		Topics []string `json:"topics"`
	} `json:"chapters"`
}

// saveChapters commits a confirmed taxonomy and applies it to everything
// already in the course.
//
// The re-classification is the point. Students upload the papers they have and
// find the syllabus afterwards, so without it a taxonomy would only ever apply
// to future uploads and the questions already indexed would stay unclassified
// forever.
func (s *Server) saveChapters(c *gin.Context) {
	courseID, ok := uuidParam(c, "courseId")
	if !ok {
		return
	}

	var req saveChaptersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}
	if len(req.Chapters) == 0 {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "A course needs at least one chapter.")
		return
	}
	if len(req.Chapters) > maxChapters {
		fail(c, http.StatusUnprocessableEntity, "validation_failed",
			fmt.Sprintf("A course cannot have more than %d chapters.", maxChapters))
		return
	}

	chapters := make([]ingest.Chapter, 0, len(req.Chapters))
	seen := map[int]bool{}
	for _, ch := range req.Chapters {
		title := strings.TrimSpace(ch.Title)
		if ch.Number < 1 || ch.Number > 99 || title == "" {
			fail(c, http.StatusUnprocessableEntity, "validation_failed",
				"Every chapter needs a number between 1 and 99 and a title.")
			return
		}
		if seen[ch.Number] {
			fail(c, http.StatusUnprocessableEntity, "validation_failed",
				fmt.Sprintf("Chapter %d appears twice.", ch.Number))
			return
		}
		seen[ch.Number] = true

		topics := make([]string, 0, len(ch.Topics))
		for _, t := range ch.Topics {
			if t = strings.TrimSpace(t); t != "" {
				topics = append(topics, t)
			}
		}
		chapters = append(chapters, ingest.Chapter{Number: ch.Number, Title: title, Topics: topics})
	}
	sort.Slice(chapters, func(i, j int) bool { return chapters[i].Number < chapters[j].Number })

	if err := s.store.ReplaceChapters(c, courseID, chapters); err != nil {
		s.respond(c, err, nil)
		return
	}

	classified, err := s.classifyPending(c, courseID, chapters)
	if err != nil {
		// The taxonomy is saved; only the back-fill failed. Reporting a 500
		// here would suggest the whole request was lost, and the user would
		// re-submit a taxonomy that is already stored.
		s.log.Error("reclassify after taxonomy change", "course", courseID, "error", err)
	}

	// The analytics are computed over chapters, so every cached aggregate for
	// this course is now wrong.
	s.cache.DeletePrefix(c, "analytics:"+courseID.String())

	saved, err := s.store.ListChapters(c, courseID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":                saved,
		"questionsClassified": classified,
	})
}

// classifyPending assigns a chapter to every question in the course that has
// none, against the taxonomy just saved.
//
// Runs inline rather than on the queue. It is bounded by the questions already
// in the course, the keyword pass is free, and the model is consulted only for
// the genuinely ambiguous minority. Doing it inline also means the response can
// report how many were classified, which is the one number that tells the user
// whether their taxonomy was any good.
func (s *Server) classifyPending(c *gin.Context, courseID uuid.UUID, chapters []ingest.Chapter) (int, error) {
	pending, err := s.store.UnclassifiedQuestions(c, courseID)
	if err != nil {
		return 0, err
	}

	classifier := ingest.NewClassifier(chapters, s.llm)
	assigned := 0
	for _, q := range pending {
		result := classifier.Classify(c, q.Text)
		if result.ChapterNumber == 0 {
			// Genuinely unclassifiable: rubric, a cover page, "Answer the
			// following questions". Left alone on purpose, because a wrong
			// chapter silently corrupts both enumeration and the analytics.
			continue
		}
		if err := s.store.SetQuestionChapter(c, q.ID, courseID, result.ChapterNumber, result.Confidence); err != nil {
			return assigned, err
		}
		assigned++
	}
	return assigned, nil
}
