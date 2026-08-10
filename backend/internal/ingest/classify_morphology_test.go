package ingest

import (
	"context"
	"testing"
)

// Reported from the UI: "all questions from chapter 1" returned a question that
// plainly belongs in chapter 2.
//
// The cause was exact substring matching. The chapter is titled "Least Squares
// Estimation" and the question asks for the "least squares estimator", so
// chapter 2 scored zero and the question fell to chapter 1 on the strength of
// the words "the linear model". These cases pin the morphology that broke it.
func TestClassifyHandlesMorphologicalVariation(t *testing.T) {
	cases := []struct {
		name     string
		question string
		want     int
	}{
		{
			name:     "estimator against a chapter titled estimation",
			question: "Derive the ordinary least squares estimator for beta in the linear model y = X beta + epsilon, and state precisely the assumptions required for it to be unbiased.",
			want:     2,
		},
		{
			name:     "tests against a chapter titled testing",
			question: "Carry out the appropriate test of the joint hypothesis using an F-test.",
			want:     3,
		},
		{
			name:     "plural residuals against singular residual",
			question: "Plot the residuals against the fitted values and comment on what the leverage of point 7 implies.",
			want:     4,
		},
		{
			name:     "a genuine chapter 1 question stays in chapter 1",
			question: "State the assumptions of the classical linear model and give the design matrix for a model with two predictors.",
			want:     1,
		},
	}

	c := NewClassifier(taxonomy, nil)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := c.Classify(context.Background(), tc.question)
			if got.ChapterNumber != tc.want {
				t.Errorf("chapter = %d, want %d (confidence %.2f, matched %q)",
					got.ChapterNumber, tc.want, got.Confidence, got.Topic)
			}
		})
	}
}

// Regression: with a chat model configured, "write your matriculation number on
// every sheet" was filed under Generalized Linear Models at full confidence.
// The keyword pass had correctly matched nothing, and escalating that to the
// model turned an honest gap into a confident error. A model asked to pick will
// pick.
func TestClassifyDoesNotEscalateWhenNothingMatches(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 5, "confidence": 1.0}`}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(),
		"Write your matriculation number on every sheet before handing in your work, and hand in this question paper together with your answers.")

	if got.ChapterNumber != 0 {
		t.Errorf("chapter = %d, want 0 for administrative text", got.ChapterNumber)
	}
	if llm.called != 0 {
		t.Errorf("LLM called %d times when no chapter matched at all, want 0", llm.called)
	}
}

// The fallback must still run when chapters genuinely compete.
func TestClassifyStillEscalatesOnAmbiguity(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 3, "confidence": 0.9}`}
	c := NewClassifier(taxonomy, llm)

	c.Classify(context.Background(),
		"Discuss the residual assumptions of the design matrix and the link function together.")

	if llm.called == 0 {
		t.Error("LLM was not consulted for a question matching several chapters")
	}
}

// Confidence reported by the model is capped: it was asked precisely because
// the evidence was weak.
func TestClassifyCapsLLMConfidence(t *testing.T) {
	llm := &stubLLM{reply: `{"chapter": 3, "confidence": 1.0}`}
	c := NewClassifier(taxonomy, llm)

	got := c.Classify(context.Background(),
		"Discuss the residual assumptions of the design matrix and the link function together.")

	if got.Source == "llm" && got.Confidence > 0.9 {
		t.Errorf("confidence = %.2f from the fallback, want it capped at 0.9", got.Confidence)
	}
}

func TestStem(t *testing.T) {
	cases := map[string]string{
		"estimator":   "estim",
		"estimation":  "estim",
		"testing":     "test",
		"tests":       "test",
		"residuals":   "residual",
		"regressions": "regression",
		// Short words must survive: over-stemming collides unrelated terms.
		"test": "test",
		"blue": "blue",
		"ols":  "ols",
	}
	for in, want := range cases {
		if got := stem(in); got != want {
			t.Errorf("stem(%q) = %q, want %q", in, got, want)
		}
	}
}

// A chapter titled "The Linear Model" must not match every question containing
// the word "the".
func TestTokensDropStopWords(t *testing.T) {
	got := tokenSet("The Linear Model")
	if got["the"] {
		t.Error("stop word 'the' survived tokenisation")
	}
	if !got["linear"] || !got["model"] {
		t.Errorf("meaningful tokens missing from %v", got)
	}
}
