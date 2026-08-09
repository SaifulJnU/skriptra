package ingest

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Question is one segmented exam question, with the page it started on so a
// citation can point at it.
type Question struct {
	Number     string
	Ordinal    int
	Text       string
	Marks      *float64
	SourcePage int
}

// Rules before models, deliberately.
//
// Exam papers are among the most rigidly formatted documents that exist: every
// question is numbered, and the numbering is printed at the start of a line.
// A regex reads that structure exactly, costs nothing, and is testable. An LLM
// pass over the same text is slower, non-deterministic, and can silently
// hallucinate a question that is not there.
//
// Both English and German headings are matched, since a German university
// corpus mixes "Aufgabe 3" and "Question 3" freely, sometimes in one paper.
var (
	reHeading = regexp.MustCompile(
		`(?im)^[\s]*(?:(?:aufgabe|question|exercise|problem|frage|task)[\s.:]*)?` +
			`(\d{1,2}\s*(?:[a-h]\)?)?)` + // 4, 4b, 4b)
			`[\s]*[.):]\s+`)

	// "(10 marks)", "[10 Punkte]", "/ 10 P."
	reMarks = regexp.MustCompile(
		`(?i)[\(\[/]\s*(\d{1,3}(?:[.,]\d)?)\s*(?:marks?|points?|punkte?|p\.?|pts?)\s*[\)\]]?`)

	// Page furniture that must never be mistaken for question text.
	reNoise = regexp.MustCompile(
		`(?im)^\s*(?:seite|page)\s+\d+(?:\s*(?:von|of)\s*\d+)?\s*$|^\s*-\s*\d+\s*-\s*$`)
)

// SegmentQuestions splits parsed pages into individual questions.
//
// Returns nil when the document has no recognisable question structure, which
// is a real and expected outcome: a lecture-notes PDF or a textbook chapter has
// no questions in it and must not be forced into some.
func SegmentQuestions(pages []Page) []Question {
	type marker struct {
		number string
		page   int
		start  int
	}

	var full strings.Builder
	pageOf := []int{} // byte offset -> page, one entry per page boundary
	starts := []int{}

	for _, p := range pages {
		text := reNoise.ReplaceAllString(p.Text, "")
		starts = append(starts, full.Len())
		pageOf = append(pageOf, p.Number)
		full.WriteString(text)
		full.WriteString("\n")
	}
	doc := full.String()

	locs := reHeading.FindAllStringSubmatchIndex(doc, -1)
	if len(locs) < 2 {
		// One heading is more likely a false positive on a numbered list than
		// a one-question exam.
		return nil
	}

	pageAt := func(offset int) int {
		page := 1
		for i, s := range starts {
			if offset >= s {
				page = pageOf[i]
			}
		}
		return page
	}

	markers := make([]marker, 0, len(locs))
	for _, l := range locs {
		num := normaliseNumber(doc[l[2]:l[3]])
		markers = append(markers, marker{number: num, page: pageAt(l[0]), start: l[0]})
	}

	out := make([]Question, 0, len(markers))
	for i, m := range markers {
		end := len(doc)
		if i+1 < len(markers) {
			end = markers[i+1].start
		}
		body := strings.TrimSpace(stripHeading(doc[m.start:end]))
		if len(body) < 20 {
			// Too short to be a question; almost always a numbered sub-item
			// inside a list, or a stray table cell.
			continue
		}

		q := Question{
			Number:     m.number,
			Ordinal:    len(out) + 1,
			Text:       collapseWhitespace(body),
			SourcePage: m.page,
		}
		if mk := reMarks.FindStringSubmatch(body); mk != nil {
			if v, err := strconv.ParseFloat(strings.Replace(mk[1], ",", ".", 1), 64); err == nil {
				q.Marks = &v
			}
		}
		out = append(out, q)
	}

	if len(out) < 2 {
		return nil
	}
	return out
}

// stripHeading removes the matched "Aufgabe 4." prefix, leaving the question.
func stripHeading(s string) string {
	if loc := reHeading.FindStringIndex(s); loc != nil && loc[0] == 0 {
		return s[loc[1]:]
	}
	return s
}

// normaliseNumber turns "4 b)" into "4b" so numbering is comparable across
// papers that space and punctuate it differently.
func normaliseNumber(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// collapseWhitespace repairs PDF line wrapping. Extracted text breaks mid
// sentence at the column edge, and leaving those newlines in place would
// fragment chunks and corrupt full-text indexing.
func collapseWhitespace(s string) string {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\r' || r == '\t' || r == ' '
	})
	return strings.Join(fields, " ")
}
