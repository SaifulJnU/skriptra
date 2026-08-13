package ingest

import "testing"

func pagesFrom(texts ...string) []Page {
	out := make([]Page, len(texts))
	for i, t := range texts {
		out[i] = Page{Number: i + 1, Text: t}
	}
	return out
}

func TestSegmentEnglishExam(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
Question 1. Derive the ordinary least squares estimator for beta and state the
assumptions required for it to be unbiased. (12 marks)

Question 2. State and prove the Gauss-Markov theorem. Explain what "best"
means in the acronym BLUE. (14 marks)
`))

	if len(qs) != 2 {
		t.Fatalf("segmented %d questions, want 2", len(qs))
	}
	if qs[0].Number != "1" || qs[1].Number != "2" {
		t.Errorf("numbers = %q, %q; want 1, 2", qs[0].Number, qs[1].Number)
	}
	if qs[0].Marks == nil || *qs[0].Marks != 12 {
		t.Errorf("marks = %v, want 12", qs[0].Marks)
	}
	// The heading itself must not survive into the question body.
	if got := qs[0].Text; got[:6] != "Derive" {
		t.Errorf("text starts with %q, want the heading stripped", got[:20])
	}
}

func TestSegmentGermanExam(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
Aufgabe 1: Leiten Sie den Kleinste-Quadrate-Schaetzer her und nennen Sie die
noetigen Annahmen. [10 Punkte]

Aufgabe 2: Formulieren und beweisen Sie das Gauss-Markov-Theorem. [14 Punkte]
`))

	if len(qs) != 2 {
		t.Fatalf("segmented %d questions from a German paper, want 2", len(qs))
	}
	if qs[1].Marks == nil || *qs[1].Marks != 14 {
		t.Errorf("marks = %v, want 14 from '[14 Punkte]'", qs[1].Marks)
	}
}

func TestSegmentSubLetteredQuestions(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
4a) Compute the residual sum of squares for the fitted model given in Table 1.
4b) Using your answer above, construct a 95% confidence interval for sigma squared.
5) Explain the difference between leverage and influence in this context.
`))

	if len(qs) != 3 {
		t.Fatalf("segmented %d, want 3", len(qs))
	}
	if qs[0].Number != "4a" || qs[1].Number != "4b" {
		t.Errorf("numbers = %q, %q; want 4a, 4b", qs[0].Number, qs[1].Number)
	}
}

// Page numbers drive citations. If they are wrong the whole product's promise
// breaks, so this is asserted explicitly.
func TestSegmentTracksSourcePageAcrossPages(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(
		"Question 1. Derive the least squares estimator and prove it is unbiased.",
		"Question 2. Construct an F-test for the joint significance of two coefficients.",
		"Question 3. Define leverage and explain how it differs from influence here.",
	))

	if len(qs) != 3 {
		t.Fatalf("segmented %d, want 3", len(qs))
	}
	for i, want := range []int{1, 2, 3} {
		if qs[i].SourcePage != want {
			t.Errorf("question %d sourcePage = %d, want %d", i+1, qs[i].SourcePage, want)
		}
	}
}

// A lecture-notes PDF has no questions. Forcing structure onto it would put
// junk in the question index, which the enumerate path then reports as fact.
func TestSegmentReturnsNilForProse(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
The linear model expresses a response as a linear combination of predictors
plus an error term. Linearity refers to the parameters, not the predictors,
so a polynomial in x is still a linear model.
`))
	if qs != nil {
		t.Errorf("segmented %d questions from prose, want nil", len(qs))
	}
}

func TestSegmentIgnoresPageFurniture(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
Seite 1 von 8
Question 1. Derive the ordinary least squares estimator and state its assumptions.
- 2 -
Question 2. Prove the Gauss-Markov theorem for the normal linear model.
Page 2 of 8
`))

	if len(qs) != 2 {
		t.Fatalf("segmented %d, want 2", len(qs))
	}
	for _, q := range qs {
		if containsAny(q.Text, "Seite 1 von 8", "Page 2 of 8", "- 2 -") {
			t.Errorf("page furniture leaked into question text: %q", q.Text)
		}
	}
}

func TestSegmentRepairsWrappedLines(t *testing.T) {
	qs := SegmentQuestions(pagesFrom(`
Question 1. Derive the ordinary
least squares
estimator for beta.
Question 2. Prove the Gauss-Markov theorem for the normal linear model here.
`))

	if len(qs) != 2 {
		t.Fatalf("segmented %d, want 2", len(qs))
	}
	want := "Derive the ordinary least squares estimator for beta."
	if qs[0].Text != want {
		t.Errorf("text = %q, want %q (wrapped lines rejoined)", qs[0].Text, want)
	}
}

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if len(sub) > 0 && len(s) >= len(sub) {
			for i := 0; i+len(sub) <= len(s); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
		}
	}
	return false
}
