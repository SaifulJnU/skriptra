package router

import "testing"

// Reported from the UI: "all short questions those that have only one mark from
// chapter 1 chapter 2" filtered on chapter 1 alone. The second chapter and the
// marks constraint were both discarded, and the empty result was reported as if
// the whole query had been answered.
func TestResolveMultipleChapters(t *testing.T) {
	cases := []struct {
		query string
		want  []int
	}{
		{"questions from chapter 1 chapter 2", []int{1, 2}},
		{"chapter 1 and chapter 3 questions", []int{1, 3}},
		{"chapters 1, 2 and 4", []int{1, 2, 4}},
		{"alle Fragen aus Kapitel 2 und Kapitel 3", []int{2, 3}},
		{"all chapter 3 questions", []int{3}},
		{"explain the residual sum of squares", nil},
	}

	for _, tc := range cases {
		got := Route(tc.query, chapters, 2026).ChapterNumbers
		if len(got) != len(tc.want) {
			t.Errorf("Route(%q) chapters = %v, want %v", tc.query, got, tc.want)
			continue
		}
		for i := range tc.want {
			if got[i] != tc.want[i] {
				t.Errorf("Route(%q) chapters = %v, want %v", tc.query, got, tc.want)
				break
			}
		}
	}
}

func TestResolveMarks(t *testing.T) {
	cases := []struct {
		query            string
		wantMin, wantMax *float64
		wantLabel        string
	}{
		{query: "questions worth 1 mark", wantMin: f(1), wantMax: f(1), wantLabel: "worth 1 mark"},
		{query: "the 10 mark questions", wantMin: f(10), wantMax: f(10), wantLabel: "worth 10 marks"},
		{query: "Fragen mit 5 Punkten", wantMin: f(5), wantMax: f(5), wantLabel: "worth 5 marks"},
		// Qualitative requests must state the threshold they assumed.
		{query: "all short questions from chapter 2", wantMax: f(3), wantLabel: "short questions (3 marks or fewer)"},
		{query: "the long questions", wantMin: f(4), wantLabel: "long questions (more than 3 marks)"},
	}

	for _, tc := range cases {
		got := Route(tc.query, chapters, 2026).Marks
		if got == nil {
			t.Errorf("Route(%q) resolved no marks filter", tc.query)
			continue
		}
		if !eq(got.Min, tc.wantMin) || !eq(got.Max, tc.wantMax) {
			t.Errorf("Route(%q) marks = [%v, %v], want [%v, %v]",
				tc.query, got.Min, got.Max, tc.wantMin, tc.wantMax)
		}
		if got.Label != tc.wantLabel {
			t.Errorf("Route(%q) label = %q, want %q", tc.query, got.Label, tc.wantLabel)
		}
	}
}

func TestNoMarksFilterWhenNoneRequested(t *testing.T) {
	for _, q := range []string{"all chapter 3 questions", "why is MLE used here"} {
		if got := Route(q, chapters, 2026).Marks; got != nil {
			t.Errorf("Route(%q) invented a marks filter: %+v", q, got)
		}
	}
}

// The reported query, end to end through the router.
func TestReportedShortQuestionQuery(t *testing.T) {
	d := Route("all short question those has only one Mark from chapter 1 chapter 2", chapters, 2026)

	if len(d.ChapterNumbers) != 2 {
		t.Errorf("chapters = %v, want both 1 and 2", d.ChapterNumbers)
	}
	if d.Marks == nil || d.Marks.Max == nil || *d.Marks.Max != 1 {
		t.Errorf("marks = %+v, want an explicit 1 mark filter", d.Marks)
	}
	if d.Intent != "enumerate" {
		t.Errorf("intent = %q, want enumerate", d.Intent)
	}
}

func f(v float64) *float64 { return &v }

func eq(a, b *float64) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}
