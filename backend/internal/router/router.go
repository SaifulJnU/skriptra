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
	"fmt"
	"regexp"
	"strings"

	"github.com/skriptra/skriptra/backend/internal/domain"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// MarksFilter is a constraint on question weight.
//
// Label carries the interpretation back to the user, because "short questions"
// is a judgement rather than a number. Applying a threshold silently would be
// the same defect as dropping a filter: the answer looks precise while resting
// on an assumption nobody stated.
type MarksFilter struct {
	Min   *float64
	Max   *float64
	Label string
}

// Decision is the routing outcome plus anything extracted along the way.
type Decision struct {
	Intent domain.QueryIntent
	// ChapterNumber is the first chapter mentioned, kept for callers that only
	// handle one. Prefer ChapterNumbers.
	ChapterNumber *int
	// ChapterNumbers holds every chapter named in the query. "chapter 1 and
	// chapter 2" is two filters, and answering it with one is answering a
	// different question. This is the step that turns "chapter two" into a
	// WHERE clause instead of an embedding: chapter membership is a property of
	// course structure, not a semantic property of the question text.
	ChapterNumbers []int
	// Marks is set when the query constrains question weight.
	Marks    *MarksFilter
	YearFrom *int
	YearTo   *int
	// QuestionType is the format the user asked for ("true/false questions"),
	// empty when they named none. Filtering on a format the corpus does not
	// contain must return nothing rather than silently widening the query.
	QuestionType string
	// Confidence is low when the rules did not match strongly; callers may use
	// it to decide whether to consult the LLM classifier.
	Confidence float64
}

var (
	// Analyse asks a question *about the corpus*, not about the subject.
	//
	// The words that distinguish it are dangerous on their own, because this is
	// a statistics course and its vocabulary is the same vocabulary. The
	// evaluation harness caught "Derive the canonical link for the Poisson
	// distribution" being answered with a chapter-frequency table, because the
	// rule matched the bare word "distribution". "Statistic" was worse still:
	// every F-statistic and test statistic in the corpus would have tripped it.
	//
	// The corpus-level words are now only accepted when attached to something
	// about the corpus: a distribution *of questions*, the *most tested*
	// chapter. A distribution of a random variable stays an explanation.
	reAnalyse = regexp.MustCompile(`(?i)` +
		`\b(?:how many|frequen\w*|häufig\w*)\b` +
		`|\bmost\s+(?:frequently|commonly|often|tested|asked|examined)\b` +
		`|\b(?:tested|asked|examined|appears?|appeared)\s+most\b` +
		`|\b(?:distribution|verteilung|breakdown|trend)\s+(?:of|über|von)\s+` +
		`(?:the\s+|die\s+|der\s+)?(?:question|topic|chapter|exam|mark|point|kapitel|frage|aufgabe)\w*\b` +
		`|\bwhich\s+(?:chapter|topic|kapitel|thema)\w*\s+(?:is|are|were|was|ist|sind)\b` +
		`|\bstatistics\s+(?:on|about|for|über)\b`)
	reSimilar = regexp.MustCompile(`(?i)\b(similar|like this|same kind|repeated|appeared before|come up before|ähnlich\w*)\b`)
	reList    = regexp.MustCompile(`(?i)\b(all|list|every|show me|give me|which questions|alle|liste|zeig)\b`)
	reExplain = regexp.MustCompile(`(?i)\b(why|how|explain|derive|prove|meaning|what is|what does|warum|wie|erkläre|erklär|was ist)\b`)

	// Digits and German/English number words, so "chapter two" and "Kapitel 2"
	// resolve identically.
	reChapterDigit = regexp.MustCompile(`(?i)\b(?:chapters?|kapitel|ch\.?|kap\.?)\s*(\d{1,2})\b`)
	reChapterWord  = regexp.MustCompile(`(?i)\b(?:chapters?|kapitel)\s+([a-zäöüß]+)`)

	reYearRange = regexp.MustCompile(`\b(19|20)(\d{2})\s*(?:-|, |to|bis)\s*(19|20)(\d{2})\b`)
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

	d.ChapterNumbers = resolveChapters(q, chapters)
	if len(d.ChapterNumbers) > 0 {
		first := d.ChapterNumbers[0]
		d.ChapterNumber = &first
	}
	d.YearFrom, d.YearTo = resolveYears(q, currentYear)
	d.Marks = resolveMarks(q)

	if t := ingest.ParseQuestionType(q); t != ingest.TypeUnknown {
		d.QuestionType = string(t)
	}

	// Naming a format or a marks constraint is itself a request to list.
	// "true/false questions from chapter 2" and "the one mark questions"
	// contain no verb the list patterns match, but both are plainly listings.
	if (d.QuestionType != "" || d.Marks != nil) && !reExplain.MatchString(q) {
		d.Intent = domain.IntentEnumerate
		d.Confidence = 0.85
		return d
	}

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
		// "explain the most common chapter 3 question", select, then explain.
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

// reMarks matches an explicit weight: "1 mark", "worth 2 marks", "10 Punkte".
var reMarks = regexp.MustCompile(`(?i)\b(?:worth\s+)?(\d{1,3}|one|two|three|four|five)[\s-]*(?:marks?|points?|punkt(?:e|en)?)\b`)

// reShort matches a qualitative request for small questions.
var reShort = regexp.MustCompile(`(?i)\b(short|small|quick|brief|kurze?|kurz)\s*(?:answer\s*)?(?:question|frage)?s?\b`)

// reLong matches the opposite.
var reLong = regexp.MustCompile(`(?i)\b(long|big|major|large|lange?)\s*(?:answer\s*)?(?:question|frage)?s?\b`)

// shortQuestionThreshold is where "short" is taken to end.
//
// An arbitrary line, which is exactly why the label naming it is surfaced to
// the user. Silently applying a threshold and reporting the result as exact
// would be the same failure as dropping a filter.
const shortQuestionThreshold = 3.0

func resolveMarks(q string) *MarksFilter {
	// An explicit number always wins over a qualitative word.
	if m := reMarks.FindStringSubmatch(q); m != nil {
		n := atoi(m[1])
		if n == 0 {
			n = numberWords[strings.ToLower(m[1])]
		}
		if n > 0 {
			v := float64(n)
			unit := "marks"
			if n == 1 {
				unit = "mark"
			}
			return &MarksFilter{Min: &v, Max: &v, Label: fmt.Sprintf("worth %d %s", n, unit)}
		}
	}

	if reShort.MatchString(q) {
		max := shortQuestionThreshold
		return &MarksFilter{
			Max:   &max,
			Label: fmt.Sprintf("short questions (%.0f marks or fewer)", shortQuestionThreshold),
		}
	}
	if reLong.MatchString(q) {
		min := shortQuestionThreshold + 1
		return &MarksFilter{
			Min:   &min,
			Label: fmt.Sprintf("long questions (more than %.0f marks)", shortQuestionThreshold),
		}
	}
	return nil
}

// resolveChapters returns every chapter the query names, in the order given.
//
// "chapter 1 chapter 2" is two constraints. Returning only the first silently
// answers a narrower question, which is how a request spanning two chapters
// came back reporting nothing found in one of them.
func resolveChapters(q string, chapters []Chapter) []int {
	seen := map[int]bool{}
	var out []int
	add := func(n int) {
		if n > 0 && !seen[n] {
			seen[n] = true
			out = append(out, n)
		}
	}

	for _, m := range reChapterDigit.FindAllStringSubmatch(q, -1) {
		add(atoi(m[1]))
	}
	for _, m := range reChapterWord.FindAllStringSubmatch(q, -1) {
		add(numberWords[m[1]])
	}
	// A bare "1 and 2" following a chapter reference, as in "chapter 1 and 2".
	if len(out) > 0 {
		if m := reChapterList.FindStringSubmatch(q); m != nil {
			for _, part := range regexp.MustCompile(`\d+`).FindAllString(m[0], -1) {
				add(atoi(part))
			}
		}
	}

	if len(out) == 0 {
		if n := resolveChapter(q, chapters); n != nil {
			add(*n)
		}
	}
	return out
}

// reChapterList catches "chapters 1, 2 and 4" where only the first number
// follows the word "chapter".
var reChapterList = regexp.MustCompile(`(?i)\b(?:chapters?|kapitel)\s*\d+(?:\s*(?:,|and|und|&|to|bis|-)\s*\d+)+`)

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
