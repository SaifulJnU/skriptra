package router

import (
	"testing"

	"github.com/skriptra/skriptra/backend/internal/domain"
)

var chapters = []Chapter{
	{Number: 1, Title: "The Linear Model", Topics: []string{"design matrix"}},
	{Number: 2, Title: "Least Squares Estimation", Topics: []string{"OLS", "Gauss-Markov"}},
	{Number: 3, Title: "Inference and Hypothesis Testing", Topics: []string{"F-test"}},
	{Number: 5, Title: "Generalized Linear Models", Topics: []string{"logistic"}},
}

func TestRouteIntent(t *testing.T) {
	cases := []struct {
		question string
		want     domain.QueryIntent
	}{
		// Enumerate: exhaustive, must never go to top-k retrieval.
		{"Give me all Chapter 3 questions from the last five years", domain.IntentEnumerate},
		{"list every question on hypothesis testing", domain.IntentEnumerate},
		{"Zeig mir alle Fragen aus Kapitel 2", domain.IntentEnumerate},

		// Analyse: an aggregate; the model is not in this path at all.
		{"Which chapters are tested most often?", domain.IntentAnalyse},
		{"how many times has chapter 3 come up", domain.IntentAnalyse},

		{"Has a question about the Gauss-Markov theorem appeared before?", domain.IntentSimilar},
		{"show me similar questions", domain.IntentSimilar},

		{"Why is maximum likelihood used in this question?", domain.IntentExplain},
		{"What is regression?", domain.IntentExplain},
		{"Warum ist der F-Test hier korrekt?", domain.IntentExplain},

		// Both a selection and an explanation.
		{"explain all chapter 3 questions", domain.IntentHybrid},

		// A bare chapter reference is a browse, not a question.
		{"chapter 3", domain.IntentEnumerate},
	}

	for _, tc := range cases {
		if got := Route(tc.question, chapters, 2026).Intent; got != tc.want {
			t.Errorf("Route(%q) intent = %q, want %q", tc.question, got, tc.want)
		}
	}
}

func TestResolveChapterFromNumberWordAndTitle(t *testing.T) {
	cases := []struct {
		question string
		want     int
	}{
		{"all chapter 3 questions", 3},
		{"questions from Chapter two", 2},
		{"alle Fragen aus Kapitel 2", 2},
		{"ch. 5 please", 5},
		{"questions on Inference and Hypothesis Testing", 3},
		{"anything about the F-test", 3},
	}

	for _, tc := range cases {
		got := Route(tc.question, chapters, 2026).ChapterNumber
		if got == nil {
			t.Errorf("Route(%q) resolved no chapter, want %d", tc.question, tc.want)
			continue
		}
		if *got != tc.want {
			t.Errorf("Route(%q) chapter = %d, want %d", tc.question, *got, tc.want)
		}
	}
}

// "Generalized Linear Models" contains "Linear Model". Longest match must win,
// or every GLM question would be misfiled into chapter 1.
func TestResolveChapterPrefersLongestTitleMatch(t *testing.T) {
	got := Route("questions on Generalized Linear Models", chapters, 2026).ChapterNumber
	if got == nil || *got != 5 {
		t.Fatalf("chapter = %v, want 5 (longest title match)", got)
	}
}

func TestResolveChapterAbsent(t *testing.T) {
	if got := Route("explain the residual sum of squares", chapters, 2026).ChapterNumber; got != nil {
		t.Errorf("chapter = %d, want nil when no chapter is referenced", *got)
	}
}

func TestResolveYears(t *testing.T) {
	cases := []struct {
		question           string
		wantFrom, wantTo   int
	}{
		{"all chapter 3 questions from the last five years", 2022, 2026},
		{"questions between 2020-2023", 2020, 2023},
		{"the 2025 exam", 2025, 2025},
	}

	for _, tc := range cases {
		d := Route(tc.question, chapters, 2026)
		if d.YearFrom == nil || d.YearTo == nil {
			t.Errorf("Route(%q) resolved no year range", tc.question)
			continue
		}
		if *d.YearFrom != tc.wantFrom || *d.YearTo != tc.wantTo {
			t.Errorf("Route(%q) years = %d-%d, want %d-%d",
				tc.question, *d.YearFrom, *d.YearTo, tc.wantFrom, tc.wantTo)
		}
	}
}
