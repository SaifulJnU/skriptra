package ingest

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

// Chapter extraction from a syllabus or a book's table of contents.
//
// The taxonomy has to come from somewhere. Until now it came only from the
// seed, so a course created through the product could never classify anything:
// classification scores a question against chapter vocabulary, and with no
// chapters there is nothing to score against.
//
// A contents page is the right source because it is what the course itself
// says its structure is. Inferring chapters from the questions instead would
// define the syllabus by what happened to be examined, which is backwards, and
// would make the analytics circular: "which chapters are tested most" cannot be
// answered by a taxonomy derived from what was tested.

var (
	// A contents line: "3 Least Squares Estimation . . . . . 41",
	// "Chapter 3: Least Squares", "3.  Least Squares Estimation".
	//
	// The trailing page number and dot leader are optional because a syllabus
	// often has neither, and a chapter with no page reference is still a
	// chapter.
	reOutlineLine = regexp.MustCompile(
		`(?im)^\s*(?:(?:chapter|kapitel|kap\.?|unit|teil)\s*)?` + // optional word
			`(\d{1,2})` + // the number
			`(?:\.\d+)*` + // reject sub-sections below, captured to detect them
			`\s*[.):\-]?\s+` + // separator
			`(\S[^\n]*?)` + // the title
			`(?:\s*[.\s·]{3,}\s*\d{1,4})?\s*$`) // optional dot leader and page

	// A sub-section, "3.1 Ordinary least squares". Matched so it can be
	// excluded: a table of contents lists both, and treating 3.1 as chapter 3
	// would produce a dozen duplicates of every chapter.
	reSubSection = regexp.MustCompile(`(?im)^\s*(?:chapter|kapitel|kap\.?)?\s*\d{1,2}\.\d`)

	// Front and back matter carry numbers in some books and are not chapters.
	reNotAChapter = regexp.MustCompile(
		`(?i)^(contents|inhalt\w*|preface|vorwort|foreword|introduction to the|index|` +
			`bibliograph\w*|references|literatur\w*|appendix|anhang|glossar\w*|` +
			`acknowledg\w*|about the author|notation|symbols|solutions|lösungen)\b`)
)

// Outline is a proposed taxonomy. It is proposed rather than saved because a
// contents page is messy and the person who owns the course is the only one who
// can say whether "3 Least Squares Estimation" is a chapter or a section
// heading inside one.
type Outline struct {
	Chapters []Chapter `json:"chapters"`
	// Source records how the list was produced, so a bad taxonomy can be traced
	// to the method rather than guessed at.
	Source string `json:"source"`
}

// ParseOutline reads a chapter list out of an already-extracted document.
//
// Rules first, for the same reason segmentation uses them: a table of contents
// is one of the most rigidly formatted things in publishing, and a regex over
// it is free, instant and testable where a model call is none of those. The
// model is the fallback for a syllabus written as prose, which is common enough
// for a course handout and impossible to match with a pattern.
func ParseOutline(ctx context.Context, pages []Page, llm provider.LLM) (Outline, error) {
	if chapters := outlineByRules(pages); len(chapters) >= 2 {
		return Outline{Chapters: chapters, Source: "rules"}, nil
	}

	// One chapter is more likely a false positive on a numbered list than a
	// real one-chapter course, so it is not accepted either.
	if llm == nil {
		return Outline{}, fmt.Errorf("no chapter list found, and no model is configured to read a prose outline")
	}

	chapters, err := outlineByLLM(ctx, pages, llm)
	if err != nil {
		return Outline{}, err
	}
	if len(chapters) < 2 {
		return Outline{}, fmt.Errorf("no chapter list could be read from this document")
	}
	return Outline{Chapters: chapters, Source: "llm"}, nil
}

func outlineByRules(pages []Page) []Chapter {
	// A contents page is near the front. Scanning a whole textbook would match
	// every numbered heading in the body and produce hundreds of "chapters".
	limit := len(pages)
	if limit > 12 {
		limit = 12
	}

	seen := map[int]Chapter{}
	for _, p := range pages[:limit] {
		for _, line := range strings.Split(p.Text, "\n") {
			if reSubSection.MatchString(line) {
				continue
			}
			m := reOutlineLine.FindStringSubmatch(line)
			if m == nil {
				continue
			}

			number, err := strconv.Atoi(m[1])
			if err != nil || number < 1 || number > 40 {
				continue
			}

			title := cleanTitle(m[2])
			if title == "" || reNotAChapter.MatchString(title) {
				continue
			}
			// A title of one short word is almost always a stray table cell or
			// a page header, not a chapter.
			if len([]rune(title)) < 4 {
				continue
			}

			// First occurrence wins. A contents page lists a chapter once; a
			// later match on the same number is the running header of the
			// chapter itself.
			if _, exists := seen[number]; !exists {
				seen[number] = Chapter{Number: number, Title: title, Topics: topicsFrom(title)}
			}
		}
	}

	if len(seen) == 0 {
		return nil
	}

	out := make([]Chapter, 0, len(seen))
	for _, ch := range seen {
		out = append(out, ch)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Number < out[j].Number })

	// The numbering has to be a plausible run.
	//
	// A contents page yields 1, 2, 3, 4. A page of scattered numbered prose
	// yields 3, 7, 12, and accepting that as a syllabus would be worse than
	// admitting nothing was found, because every question in the course would
	// then be classified against nonsense.
	//
	// The run is not required to start at 1: a book that numbers its preface
	// leaves the first real chapter at 2, and front matter has already been
	// filtered out above. What is required is density, that the numbers found
	// cover most of the range they span.
	first, last := out[0].Number, out[len(out)-1].Number
	span := last - first + 1
	if first > 3 || span > 2*len(out) {
		return nil
	}
	return out
}

// cleanTitle strips the dot leader, trailing page number and stray punctuation
// that a contents page leaves attached to a title.
func cleanTitle(s string) string {
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`[.\s·_]{3,}.*$`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`\s+\d{1,4}$`).ReplaceAllString(s, "")
	s = strings.Trim(s, " .·_-–\t")
	return strings.Join(strings.Fields(s), " ")
}

// topicsFrom seeds the classifier's vocabulary from the title itself.
//
// The title alone is thin evidence: a chapter called "Model Diagnostics" shares
// no words with a question about Cook's distance. Seeding topics from the title
// at least gives the keyword classifier something, and the review step exists
// so the owner can add the terms that actually appear in their exams. That
// editing step is where most of the classification quality comes from, which is
// why the taxonomy is proposed and confirmed rather than saved silently.
func topicsFrom(title string) []string {
	words := strings.FieldsFunc(strings.ToLower(title), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	stop := map[string]bool{
		"the": true, "and": true, "of": true, "to": true, "in": true, "for": true,
		"a": true, "an": true, "with": true, "der": true, "die": true, "das": true,
		"und": true, "von": true, "für": true, "im": true, "zur": true,
	}

	out := make([]string, 0, len(words))
	for _, w := range words {
		if len(w) > 2 && !stop[w] {
			out = append(out, w)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

const outlinePrompt = `You are reading a university course syllabus or a textbook table of contents.

Return ONLY a JSON array of chapters, no prose, in this exact shape:
[{"number": 1, "title": "The Linear Model", "topics": ["design matrix", "assumptions"]}]

Rules:
- number is the chapter number as printed, starting at 1.
- title is the chapter title, without the page number.
- topics are 2 to 5 short technical terms a student would expect to see in an exam question on that chapter. Take them from the chapter's own sub-headings where there are any.
- Skip the preface, index, bibliography, appendices and solutions.
- If the document does not contain a chapter list at all, return [].`

func outlineByLLM(ctx context.Context, pages []Page, llm provider.LLM) ([]Chapter, error) {
	var buf strings.Builder
	for i, p := range pages {
		if i >= 12 {
			break
		}
		buf.WriteString(p.Text)
		buf.WriteString("\n")
		// A contents page is short. Sending a whole textbook would blow the
		// context window and cost far more than the task is worth.
		if buf.Len() > 12000 {
			break
		}
	}

	res, err := llm.Generate(ctx, provider.GenerateRequest{
		Messages: []provider.Message{
			{Role: provider.RoleSystem, Content: outlinePrompt},
			{Role: provider.RoleUser, Content: buf.String()},
		},
		Temperature: 0,
		MaxTokens:   1200,
	})
	if err != nil {
		return nil, err
	}

	var parsed []Chapter
	if err := json.Unmarshal([]byte(extractJSONArray(res.Text)), &parsed); err != nil {
		return nil, fmt.Errorf("model did not return a usable chapter list: %w", err)
	}

	out := make([]Chapter, 0, len(parsed))
	for _, ch := range parsed {
		ch.Title = cleanTitle(ch.Title)
		if ch.Number < 1 || ch.Number > 40 || ch.Title == "" {
			continue
		}
		if len(ch.Topics) == 0 {
			ch.Topics = topicsFrom(ch.Title)
		}
		out = append(out, ch)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Number < out[j].Number })
	return out, nil
}

// extractJSONArray pulls the array out of a reply that wrapped it in prose or a
// code fence, which small models do regardless of instructions.
func extractJSONArray(s string) string {
	start := strings.Index(s, "[")
	end := strings.LastIndex(s, "]")
	if start == -1 || end == -1 || end < start {
		return "[]"
	}
	return s[start : end+1]
}
