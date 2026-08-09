// Package router decides how a question gets answered before any work is done.
//
// This is the core design claim of Skriptra: most questions a student asks are
// not retrieval questions.
//
//	"give me all chapter 3 questions"     -> exhaustive SQL. Top-k cannot do this.
//	"which chapters are tested most"      -> an aggregate. No model involved.
//	"has this been asked before"          -> vector k-NN over questions.
//	"why is MLE used here"                -> retrieval, then generation.
//
// Routing to the right one is what makes answers correct rather than merely
// plausible, and it means most requests never pay for a model call.
//
// Rules first, deliberately. The patterns below are fast, free, deterministic
// and testable; an LLM classifier is the fallback for genuinely ambiguous
// phrasing, not the default path.
package router

import (
	"regexp"
	"strings"

	"github.com/skriptra/skriptra/backend/internal/domain"
)

// Decision is the routing outcome plus anything extracted along the way.
type Decision struct {
	Intent domain.QueryIntent
	// ChapterNumber is resolved from the query text when present. This is the
	// step that turns "chapter two" into a WHERE clause instead of an
	// embedding — chapter membership is a property of course structure, not a
	// semantic property of the question text.
	ChapterNumber *int
	YearFrom      *int
	YearTo        *int
	// Confidence is low when the rules did not match strongly; callers may use
	// it to decide whether to consult the LLM classifier.
	Confidence float64
}

var (
	reAnalyse = regexp.MustCompile(`(?i)\b(most|frequen\w*|often|trend|distribution|how many|statistic\w*|häufig\w*|verteilung)\b`)
	reSimilar = regexp.MustCompile(`(?i)\b(similar|like this|same kind|repeated|appeared before|come up before|ähnlich\w*)\b`)
	reList    = regexp.MustCompile(`(?i)\b(all|list|every|show me|give me|which questions|alle|liste|zeig)\b`)
	reExplain = regexp.MustCompile(`(?i)\b(why|how|explain|derive|prove|meaning|what is|what does|warum|wie|erkläre|erklär|was ist)\b`)

	// Digits and German/English number words, so "chapter two" and "Kapitel 2"
	// resolve identically.
	reChapterDigit = regexp.MustCompile(`(?i)\b(?:chapter|kapitel|ch\.?|kap\.?)\s*(\d{1,2})\b`)
	reChapterWord  = regexp.MustCompile(`(?i)\b(?:chapter|kapitel)\s+([a-zäöüß]+)`)

	reYearRange = regexp.MustCompile(`\b(19|20)(\d{2})\s*(?:-|–|to|bis)\s*(19|20)(\d{2})\b`)
	reYear      = regexp.MustCompile(`\b(19|20)(\d{2})\b`)
	reLastNYear = regexp.MustCompile(`(?i)\blast\s+(\d{1,2}|two|three|four|five|ten)\s+years?\b`)
)

var numberWords = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
	"eins": 1, "zwei": 2, "drei": 3, "vier": 4, "fünf": 5,
	"sechs": 6, "sieben": 7, "acht": 8, "neun": 9, "zehn": 10,
}

// Chapter is the minimum the router needs to match a chapter by its title.
type Chapter struct {
	Number int
	Title  string
	Topics []string
}

// Route classifies a question. `chapters` is the course taxonomy, used so a
// query naming a chapter by title ("the maximum likelihood chapter") resolves
// as well as one naming it by number.
func Route(question string, chapters []Chapter, currentYear int) Decision {
	q := strings.ToLower(strings.TrimSpace(question))
	d := Decision{Intent: domain.IntentExplain, Confidence: 0.5}

	d.ChapterNumber = resolveChapter(q, chapters)
	d.YearFrom, d.YearTo = resolveYears(q, currentYear)

	wantsList := reList.MatchString(q)
	wantsExplanation := reExplain.MatchString(q)

	switch {
	case reAnalyse.MatchString(q):
		d.Intent = domain.IntentAnalyse
		d.Confidence = 0.9

	case reSimilar.MatchString(q):
		d.Intent = domain.IntentSimilar
		d.Confidence = 0.85

	case wantsList && wantsExplanation:
		// "explain the most common chapter 3 question" — select, then explain.
		d.Intent = domain.IntentHybrid
		d.Confidence = 0.75

	case wantsList:
		d.Intent = domain.IntentEnumerate
		d.Confidence = 0.9

	case wantsExplanation:
		d.Intent = domain.IntentExplain
		d.Confidence = 0.85

	default:
		// A bare chapter reference with no verb ("chapter 3") is a browse.
		if d.ChapterNumber != nil {
			d.Intent = domain.IntentEnumerate
			d.Confidence = 0.6
		}
	}
	return d
}

func resolveChapter(q string, chapters []Chapter) *int {
	if m := reChapterDigit.FindStringSubmatch(q); m != nil {
		n := atoi(m[1])
		if n > 0 {
			return &n
		}
	}
	if m := reChapterWord.FindStringSubmatch(q); m != nil {
		if n, ok := numberWords[m[1]]; ok {
			return &n
		}
	}
	// Match by chapter title, longest first so "Generalized Linear Models"
	// wins over "Linear Model" when both appear.
	best := -1
	bestLen := 0
	for _, c := range chapters {
		t := strings.ToLower(c.Title)
		if t != "" && len(t) > bestLen && strings.Contains(q, t) {
			best, bestLen = c.Number, len(t)
		}
		for _, topic := range c.Topics {
			tp := strings.ToLower(topic)
			if len(tp) > 3 && len(tp) > bestLen && strings.Contains(q, tp) {
				best, bestLen = c.Number, len(tp)
			}
		}
	}
	if best > 0 {
		return &best
	}
	return nil
}

func resolveYears(q string, currentYear int) (*int, *int) {
	if m := reYearRange.FindStringSubmatch(q); m != nil {
		from := atoi(m[1] + m[2])
		to := atoi(m[3] + m[4])
		if from > to {
			from, to = to, from
		}
		return &from, &to
	}
	if m := reLastNYear.FindStringSubmatch(q); m != nil {
		n := atoi(m[1])
		if n == 0 {
			n = numberWords[strings.ToLower(m[1])]
		}
		if n > 0 && currentYear > 0 {
			from := currentYear - n + 1
			to := currentYear
			return &from, &to
		}
	}
	if m := reYear.FindStringSubmatch(q); m != nil {
		y := atoi(m[1] + m[2])
		return &y, &y
	}
	return nil, nil
}

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
