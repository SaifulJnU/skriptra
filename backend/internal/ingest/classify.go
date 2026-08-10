package ingest

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

// Chapter is one entry in the course taxonomy that questions are classified
// against. Built once at course setup from a syllabus or a textbook contents
// page.
type Chapter struct {
	Number int
	Title  string
	Topics []string
}

// Classification is the outcome for one question.
//
// ChapterNumber is zero when no chapter could be assigned. That is a real
// answer, not a failure: leaving a question unclassified and visibly flagged is
// better than filing it under a chapter it does not belong to, because a wrong
// chapter silently corrupts both the enumerate path and the analytics.
type Classification struct {
	ChapterNumber int
	Confidence    float64
	// Source is "keyword" or "llm", recorded so a bad batch can be traced to
	// the method that produced it.
	Source string
	Topic  string
}

// ConfidenceThreshold is the point below which a keyword match is considered
// too weak to trust on its own and the LLM is consulted.
const ConfidenceThreshold = 0.70

// Classifier assigns questions to chapters.
//
// Keyword scoring runs first because it is free, instant and deterministic. The
// LLM is a fallback for the genuinely ambiguous minority, which keeps ingestion
// fast and cheap: a 30-question paper typically needs only a handful of model
// calls rather than 30.
type Classifier struct {
	chapters []Chapter
	llm      provider.LLM
}

func NewClassifier(chapters []Chapter, llm provider.LLM) *Classifier {
	return &Classifier{chapters: chapters, llm: llm}
}

// Classify assigns one question, escalating to the LLM only when keyword
// scoring is not confident.
func (c *Classifier) Classify(ctx context.Context, questionText string) Classification {
	best := c.byKeyword(questionText)
	if best.Confidence >= ConfidenceThreshold || c.llm == nil {
		return best
	}

	if llmResult, err := c.byLLM(ctx, questionText); err == nil && llmResult.ChapterNumber != 0 {
		return llmResult
	}
	// The model was unavailable or unhelpful. Return the weak keyword result
	// rather than nothing, so the UI can show it flagged for review.
	return best
}

// byKeyword scores each chapter by how much of its topic vocabulary appears in
// the question.
//
// Longer terms score higher: matching "generalized linear model" is far more
// informative than matching "model", and without that weighting every question
// mentioning a common word drifts to whichever chapter listed it.
func (c *Classifier) byKeyword(text string) Classification {
	norm := normaliseForMatch(text)

	type scored struct {
		number  int
		score   float64
		topic   string
		matched int
	}
	results := make([]scored, 0, len(c.chapters))

	for _, ch := range c.chapters {
		var total float64
		var bestTopic string
		var bestTopicScore float64
		matched := 0

		terms := append([]string{ch.Title}, ch.Topics...)
		for _, term := range terms {
			t := normaliseForMatch(term)
			if len(t) < 3 || !strings.Contains(norm, t) {
				continue
			}
			// Weight by term length: specific phrases are strong evidence,
			// short common words are weak.
			w := float64(len(strings.Fields(t))) + float64(len(t))/12.0
			total += w
			matched++
			if w > bestTopicScore {
				bestTopicScore, bestTopic = w, term
			}
		}
		if matched > 0 {
			results = append(results, scored{ch.Number, total, bestTopic, matched})
		}
	}

	if len(results) == 0 {
		return Classification{Source: "keyword"}
	}

	sort.Slice(results, func(i, j int) bool { return results[i].score > results[j].score })
	top := results[0]

	// Confidence reflects how decisively the winner beat the runner-up. Two
	// chapters scoring almost equally is exactly the ambiguous case the LLM
	// should arbitrate, so it must not be reported as confident.
	confidence := 0.55 + 0.1*float64(top.matched)
	if len(results) > 1 {
		margin := (top.score - results[1].score) / top.score
		confidence = 0.5 + 0.45*margin
	}
	if confidence > 0.97 {
		confidence = 0.97
	}

	return Classification{
		ChapterNumber: top.number,
		Confidence:    confidence,
		Source:        "keyword",
		Topic:         top.topic,
	}
}

func (c *Classifier) byLLM(ctx context.Context, text string) (Classification, error) {
	var list strings.Builder
	for _, ch := range c.chapters {
		fmt.Fprintf(&list, "%d. %s", ch.Number, ch.Title)
		if len(ch.Topics) > 0 {
			fmt.Fprintf(&list, " (%s)", strings.Join(ch.Topics, ", "))
		}
		list.WriteString("\n")
	}

	// Explicitly permitting 0 matters. Without it the model always picks
	// something, and a forced guess is worse than an honest gap.
	prompt := fmt.Sprintf(`Course chapters:
%s
Exam question:
%s

Which chapter does this question belong to? Reply with JSON only:
{"chapter": <number>, "confidence": <0.0-1.0>}
Use {"chapter": 0} if it does not clearly belong to any of them.`, list.String(), text)

	res, err := c.llm.Generate(ctx, provider.GenerateRequest{
		Messages: []provider.Message{
			{Role: provider.RoleSystem, Content: "You classify exam questions into course chapters. Reply with JSON only."},
			{Role: provider.RoleUser, Content: prompt},
		},
		Temperature: 0,
		MaxTokens:   64,
	})
	if err != nil {
		return Classification{}, err
	}

	var parsed struct {
		Chapter    int     `json:"chapter"`
		Confidence float64 `json:"confidence"`
	}
	if err := json.Unmarshal([]byte(extractJSON(res.Text)), &parsed); err != nil {
		return Classification{}, fmt.Errorf("classifier returned unparseable output: %w", err)
	}
	if parsed.Chapter != 0 && !c.hasChapter(parsed.Chapter) {
		// Models invent chapter numbers. Reject rather than store a dangling
		// reference.
		return Classification{}, fmt.Errorf("classifier returned unknown chapter %d", parsed.Chapter)
	}
	if parsed.Confidence <= 0 || parsed.Confidence > 1 {
		parsed.Confidence = 0.75
	}
	return Classification{
		ChapterNumber: parsed.Chapter,
		Confidence:    parsed.Confidence,
		Source:        "llm",
	}, nil
}

func (c *Classifier) hasChapter(n int) bool {
	for _, ch := range c.chapters {
		if ch.Number == n {
			return true
		}
	}
	return false
}

// extractJSON pulls the first {...} out of a response, since models routinely
// wrap JSON in prose or a markdown fence however firmly they were told not to.
func extractJSON(s string) string {
	start := strings.Index(s, "{")
	end := strings.LastIndex(s, "}")
	if start >= 0 && end > start {
		return s[start : end+1]
	}
	return s
}

// normaliseForMatch lowercases and reduces punctuation to spaces so that
// "Gauss-Markov", "gauss markov" and "Gauss, Markov" all match.
func normaliseForMatch(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prevSpace := false
	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			prevSpace = false
		default:
			if !prevSpace {
				b.WriteByte(' ')
				prevSpace = true
			}
		}
	}
	return strings.TrimSpace(b.String())
}
