// Package eval measures whether retrieval actually works.
//
// RAG systems degrade silently. A chunking tweak, a prompt edit or a different
// embedding model can halve retrieval quality with no error, no failing test,
// and answers that still read fluently. Fluent and wrong is the failure mode
// this package exists to catch.
//
// The metrics are deliberately retrieval-first. If the right passage never came
// back, no amount of prompt engineering can save the answer, so recall is
// measured before anything about generation is considered.
package eval

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Case is one labelled example.
//
// RelevantQuestionNumbers is the ground truth: the questions a correct system
// must retrieve. Written by hand against real papers, which is what makes this
// a golden set rather than synthetic self-congratulation.
type Case struct {
	ID       string `json:"id"`
	Question string `json:"question"`

	// ExpectedIntent asserts the router's decision. Routing "give me all
	// chapter 3 questions" to retrieval instead of SQL is a correctness bug
	// no answer metric would catch.
	ExpectedIntent string `json:"expectedIntent"`

	ExpectedChapter         *int     `json:"expectedChapter,omitempty"`
	RelevantQuestionNumbers []string `json:"relevantQuestionNumbers,omitempty"`

	// ShouldRefuse marks questions the corpus cannot answer. A system that
	// answers these is hallucinating, and measuring that is as important as
	// measuring recall.
	ShouldRefuse bool `json:"shouldRefuse,omitempty"`
}

type Golden struct {
	CourseID uuid.UUID `json:"courseId"`
	Cases    []Case    `json:"cases"`
}

func LoadGolden(path string) (*Golden, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var g Golden
	if err := json.Unmarshal(raw, &g); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if len(g.Cases) == 0 {
		return nil, fmt.Errorf("%s contains no cases", path)
	}
	return &g, nil
}

// Result is the outcome for one case.
type Result struct {
	CaseID        string
	IntentCorrect bool
	ChapterOK     bool
	RecallAt5     float64
	ReciprocalRank float64
	Refused       bool
	RefusalOK     bool
	LatencyMs     int64
	Notes         string
}

// Report aggregates results and decides pass or fail.
type Report struct {
	Total          int     `json:"total"`
	IntentAccuracy float64 `json:"intentAccuracy"`
	ChapterAccuracy float64 `json:"chapterAccuracy"`
	RecallAt5      float64 `json:"recallAt5"`
	MRR            float64 `json:"mrr"`
	RefusalAccuracy float64 `json:"refusalAccuracy"`
	P50LatencyMs   int64   `json:"p50LatencyMs"`
	Results        []Result `json:"-"`
}

// Baseline is the committed quality floor. A change that drops below it fails
// CI. Numbers, not vibes: this is the entire point of the harness.
type Baseline struct {
	IntentAccuracy  float64 `json:"intentAccuracy"`
	ChapterAccuracy float64 `json:"chapterAccuracy"`
	RecallAt5       float64 `json:"recallAt5"`
	MRR             float64 `json:"mrr"`
	RefusalAccuracy float64 `json:"refusalAccuracy"`
}

func LoadBaseline(path string) (*Baseline, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var b Baseline
	return &b, json.Unmarshal(raw, &b)
}

// Retrieved is what the system returned for one case.
type Retrieved struct {
	Intent          string
	ChapterNumber   *int
	QuestionNumbers []string // in rank order
	Refused         bool
	Latency         time.Duration
}

// Score evaluates one case against what the system returned.
func Score(c Case, got Retrieved) Result {
	r := Result{
		CaseID:        c.ID,
		IntentCorrect: strings.EqualFold(c.ExpectedIntent, got.Intent),
		Refused:       got.Refused,
		LatencyMs:     got.Latency.Milliseconds(),
	}

	switch {
	case c.ExpectedChapter == nil && got.ChapterNumber == nil:
		r.ChapterOK = true
	case c.ExpectedChapter != nil && got.ChapterNumber != nil:
		r.ChapterOK = *c.ExpectedChapter == *got.ChapterNumber
	}

	r.RefusalOK = c.ShouldRefuse == got.Refused
	if c.ShouldRefuse {
		// Recall is meaningless when the right behaviour is to return nothing.
		r.RecallAt5, r.ReciprocalRank = 1, 1
		if !got.Refused {
			r.RecallAt5, r.ReciprocalRank = 0, 0
			r.Notes = "answered a question the corpus cannot support"
		}
		return r
	}

	if len(c.RelevantQuestionNumbers) == 0 {
		r.RecallAt5, r.ReciprocalRank = 1, 1
		return r
	}

	want := map[string]bool{}
	for _, n := range c.RelevantQuestionNumbers {
		want[strings.ToLower(n)] = true
	}

	found := 0
	for i, n := range got.QuestionNumbers {
		if !want[strings.ToLower(n)] {
			continue
		}
		if i < 5 {
			found++
		}
		if r.ReciprocalRank == 0 {
			r.ReciprocalRank = 1 / float64(i+1)
		}
	}
	r.RecallAt5 = float64(found) / float64(len(want))
	if r.RecallAt5 > 1 {
		r.RecallAt5 = 1
	}
	if r.RecallAt5 < 1 {
		r.Notes = fmt.Sprintf("recalled %d of %d relevant questions", found, len(want))
	}
	return r
}

func Aggregate(results []Result) Report {
	rep := Report{Total: len(results), Results: results}
	if len(results) == 0 {
		return rep
	}

	var intent, chapter, recall, mrr, refusal float64
	latencies := make([]int64, 0, len(results))
	for _, r := range results {
		if r.IntentCorrect {
			intent++
		}
		if r.ChapterOK {
			chapter++
		}
		if r.RefusalOK {
			refusal++
		}
		recall += r.RecallAt5
		mrr += r.ReciprocalRank
		latencies = append(latencies, r.LatencyMs)
	}

	n := float64(len(results))
	rep.IntentAccuracy = intent / n
	rep.ChapterAccuracy = chapter / n
	rep.RecallAt5 = recall / n
	rep.MRR = mrr / n
	rep.RefusalAccuracy = refusal / n

	sort.Slice(latencies, func(i, j int) bool { return latencies[i] < latencies[j] })
	rep.P50LatencyMs = latencies[len(latencies)/2]
	return rep
}

// Regressions compares a report against the committed baseline.
//
// A small tolerance absorbs the nondeterminism of a language model in the
// classification path. Without it the gate cries wolf and gets disabled, which
// is worse than having no gate.
const Tolerance = 0.02

func (r Report) Regressions(b Baseline) []string {
	var out []string
	check := func(name string, got, want float64) {
		if got < want-Tolerance {
			out = append(out, fmt.Sprintf("%s dropped to %.3f, baseline %.3f", name, got, want))
		}
	}
	check("intent accuracy", r.IntentAccuracy, b.IntentAccuracy)
	check("chapter accuracy", r.ChapterAccuracy, b.ChapterAccuracy)
	check("recall@5", r.RecallAt5, b.RecallAt5)
	check("MRR", r.MRR, b.MRR)
	check("refusal accuracy", r.RefusalAccuracy, b.RefusalAccuracy)
	return out
}

func (r Report) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "cases            %d\n", r.Total)
	fmt.Fprintf(&b, "intent accuracy  %.1f%%\n", r.IntentAccuracy*100)
	fmt.Fprintf(&b, "chapter accuracy %.1f%%\n", r.ChapterAccuracy*100)
	fmt.Fprintf(&b, "recall@5         %.3f\n", r.RecallAt5)
	fmt.Fprintf(&b, "MRR              %.3f\n", r.MRR)
	fmt.Fprintf(&b, "refusal accuracy %.1f%%\n", r.RefusalAccuracy*100)
	fmt.Fprintf(&b, "p50 latency      %dms\n", r.P50LatencyMs)
	return b.String()
}

// WriteBaseline records current numbers as the new floor. Run deliberately
// after a change that is understood to be an improvement, never automatically.
func WriteBaseline(path string, r Report) error {
	b := Baseline{
		IntentAccuracy:  round3(r.IntentAccuracy),
		ChapterAccuracy: round3(r.ChapterAccuracy),
		RecallAt5:       round3(r.RecallAt5),
		MRR:             round3(r.MRR),
		RefusalAccuracy: round3(r.RefusalAccuracy),
	}
	raw, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func round3(f float64) float64 { return math.Round(f*1000) / 1000 }

var _ = context.Background
