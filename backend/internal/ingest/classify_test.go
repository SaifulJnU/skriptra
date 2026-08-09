package ingest

import (
	"context"
	"errors"
	"testing"

	"github.com/skriptra/skriptra/backend/internal/provider"
)

var taxonomy = []Chapter{
	{Number: 1, Title: "The Linear Model", Topics: []string{"design matrix", "assumptions"}},
	{Number: 2, Title: "Least Squares Estimation", Topics: []string{"ordinary least squares", "Gauss-Markov", "BLUE"}},
	{Number: 3, Title: "Inference and Hypothesis Testing", Topics: []string{"F-test", "t-test", "confidence interval"}},
	{Number: 4, Title: "Model Diagnostics", Topics: []string{"residual", "leverage", "Cook's distance"}},
	{Number: 5, Title: "Generalized Linear Models", Topics: []string{"link function", "logistic regression", "Poisson"}},
}

// stubLLM records whether it was called, so the tests can assert that the
// expensive path is only taken when the cheap one is unsure.
type stubLLM struct {
	reply  string
	err    error
	called int
}

func (s *stubLLM) Info() provider.Info { return provider.Info{Provider: "stub", Model: "stub"} }
func (s *stubLLM) Stream(context.Context, provider.GenerateRequest) (<-chan provider.Chunk, error) {
	return nil, errors.New("not used")
}
func (s *stubLLM) Generate(context.Context, provider.GenerateRequest) (*provider.GenerateResponse, error) {
	s.called++
	if s.err != nil {
		return nil, s.err
	}
	return &provider.GenerateResponse{Text: s.reply}, nil
}

func TestClassifyByKeyword(t *testing.T) {
	cases := []struct {
		question string
		want     int
	}{
		{"State and prove the Gauss-Markov theorem and explain what BLUE means.", 2},
		{"Construct an F-test for the joint significance of two coefficients.", 3},
		{"Define leverage and Cook's distance, and explain influence.", 4},
		{"Derive the canonical link function for the Poisson distribution.", 5},
	}

	c := NewClassifier(taxonomy, nil)
	for _, tc := range cases {
		got := c.Classify(context.Background(), tc.question)
		if got.ChapterNumber != tc.want {
			t.Errorf("Classify(%.45q) = chapter %d, want %d", tc.question, got.ChapterNumber, tc.want)
		}
		if got.Source != "keyword" {
			t.Errorf("Classify(%.45q) source = %q, want keyword", tc.question, got.Source)
		}
	}
}

// "Generalized Linear Models" contains "Linear Model". Without length
// weighting every GLM question lands in chapter 1.
func TestClassifyPrefersMoreSpecificChapter(t *testing.T) {
	c := NewClassifier(taxonomy, nil)
	got := c.Classify(context.Background(), "Explain the link function in a generalized linear model.")
	if got.ChapterNumber != 5 {
		t.Errorf("chapter = %d, want 5 (specific beats generic)", got.ChapterNumber)
	}
}

// The whole point of the keyword-first design: the model is not called when
// the cheap path is already confident.
func TestClassifyDoesNotCallLLMWhenConfident(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 1, "confidence": 0.9}`}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(), "State and prove the Gauss-Markov theorem; explain BLUE and ordinary least squares.")
	if got.ChapterNumber != 2 {
		t.Fatalf("chapter = %d, want 2", got.ChapterNumber)
	}
	if llm.called != 0 {
		t.Errorf("LLM called %d times on a confident keyword match, want 0", llm.called)
	}
}

func TestClassifyEscalatesToLLMWhenUnsure(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 3, "confidence": 0.88}`}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(), "Show that the quantity in part (a) has the distribution stated in the lecture.")
	if llm.called == 0 {
		t.Fatal("LLM was not consulted for an ambiguous question")
	}
	if got.ChapterNumber != 3 || got.Source != "llm" {
		t.Errorf("got chapter %d from %q, want 3 from llm", got.ChapterNumber, got.Source)
	}
}

// A model that names a chapter the course does not have must be rejected, or
// the questions table ends up with dangling references.
func TestClassifyRejectsInventedChapter(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 99, "confidence": 0.95}`}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(), "Something entirely unrelated to any listed chapter whatsoever.")
	if got.ChapterNumber == 99 {
		t.Error("accepted chapter 99, which is not in the taxonomy")
	}
}

func TestClassifyHandlesModelWrappingJSONInProse(t *testing.T) {
	llm := &stubLLM{reply: "Sure! Here you go:\n```json\n{\"chapter\": 4, \"confidence\": 0.8}\n```"}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(), "Comment on the pattern visible in the plot from part (b).")
	if got.ChapterNumber != 4 {
		t.Errorf("chapter = %d, want 4 parsed out of the fenced reply", got.ChapterNumber)
	}
}

// Ingestion must not fail because Ollama is down. An unclassified question is
// recoverable; a failed import of a whole paper is not.
func TestClassifyDegradesWhenLLMUnavailable(t *testing.T) {
	llm := &stubLLM{err: provider.ErrUnavailable}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(), "Show that the estimator in part (a) attains the stated bound.")
	if got.Source != "keyword" {
		t.Errorf("source = %q, want the keyword result to survive an unavailable model", got.Source)
	}
}

// No chapter is a legitimate answer. Forcing one corrupts analytics silently.
func TestClassifyReturnsNoChapterWhenNothingMatches(t *testing.T) {
	c := NewClassifier(taxonomy, nil)
	got := c.Classify(context.Background(), "Write your matriculation number on every sheet before you begin.")
	if got.ChapterNumber != 0 {
		t.Errorf("chapter = %d, want 0 for administrative text", got.ChapterNumber)
	}
}

func TestClassifyFlagsLowConfidenceForReview(t *testing.T) {
	c := NewClassifier(taxonomy, nil)
	got := c.Classify(context.Background(), "Discuss the residual assumptions of the design matrix and the link function.")
	// Three chapters compete here, so confidence must not look decisive.
	if got.ChapterNumber != 0 && got.Confidence >= ConfidenceThreshold {
		t.Errorf("confidence = %.2f on a question matching several chapters; should be below %.2f",
			got.Confidence, ConfidenceThreshold)
	}
}
