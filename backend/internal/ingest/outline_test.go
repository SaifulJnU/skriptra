package ingest

import (
	"context"
	"testing"
)

func TestParseOutlineFromContentsPage(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
Contents

Preface . . . . . . . . . . . . . . . . . . vii

1  The Linear Model . . . . . . . . . . . . . 1
   1.1 The design matrix . . . . . . . . . .  3
   1.2 Classical assumptions . . . . . . . .  9
2  Least Squares Estimation  . . . . . . . . 17
3  Inference and Hypothesis Testing . . . . 41
4  Model Diagnostics . . . . . . . . . . . . 73
5  Generalized Linear Models . . . . . . . . 98

Bibliography . . . . . . . . . . . . . . . 140
Index . . . . . . . . . . . . . . . . . . . 147
`}}

	got, err := ParseOutline(context.Background(), pages, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got.Source != "rules" {
		t.Fatalf("source = %q, want rules, the model should not be needed here", got.Source)
	}
	if len(got.Chapters) != 5 {
		t.Fatalf("got %d chapters, want 5: %+v", len(got.Chapters), got.Chapters)
	}

	want := []string{
		"The Linear Model",
		"Least Squares Estimation",
		"Inference and Hypothesis Testing",
		"Model Diagnostics",
		"Generalized Linear Models",
	}
	for i, w := range want {
		if got.Chapters[i].Number != i+1 {
			t.Errorf("chapter %d: number = %d", i, got.Chapters[i].Number)
		}
		if got.Chapters[i].Title != w {
			t.Errorf("chapter %d: title = %q, want %q", i+1, got.Chapters[i].Title, w)
		}
	}
}

// A contents page lists sub-sections too. Treating 1.1 and 1.2 as chapter 1
// would produce a dozen duplicates of every chapter.
func TestParseOutlineIgnoresSubSections(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
1 The Linear Model 1
1.1 The design matrix 3
1.2 Classical assumptions 9
1.3 Interpretation 12
2 Least Squares Estimation 17
2.1 Ordinary least squares 18
`}}
	got, err := ParseOutline(context.Background(), pages, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Chapters) != 2 {
		t.Fatalf("got %d chapters, want 2: %+v", len(got.Chapters), got.Chapters)
	}
}

// Front and back matter carry numbers in some books and are not chapters.
func TestParseOutlineSkipsFrontAndBackMatter(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
1 Preface 1
2 The Linear Model 5
3 Least Squares Estimation 20
4 Appendix A: Matrix algebra 90
5 Index 99
`}}
	got, err := ParseOutline(context.Background(), pages, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, ch := range got.Chapters {
		switch ch.Title {
		case "Preface", "Index":
			t.Fatalf("front or back matter accepted as a chapter: %q", ch.Title)
		}
	}
}

// German syllabus wording, since the corpus is bilingual.
func TestParseOutlineGerman(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
Inhaltsverzeichnis

Kapitel 1: Das lineare Modell . . . . . 1
Kapitel 2: Kleinste-Quadrate-Schaetzung . . . 18
Kapitel 3: Hypothesentests . . . . . . 44
`}}
	got, err := ParseOutline(context.Background(), pages, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Chapters) != 3 {
		t.Fatalf("got %d chapters, want 3: %+v", len(got.Chapters), got.Chapters)
	}
	if got.Chapters[0].Title != "Das lineare Modell" {
		t.Fatalf("title = %q", got.Chapters[0].Title)
	}
}

// Scattered numbered prose is not a syllabus. Accepting it would fill a course
// with nonsense chapters that every question then gets classified against.
func TestParseOutlineRejectsNonContents(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
Some notes on the exam.

3 marks are available for part one and you should
7 lines of working are expected here
12 candidates sat this paper last year
`}}
	if _, err := ParseOutline(context.Background(), pages, nil); err == nil {
		t.Fatal("scattered numbers were accepted as a chapter list")
	}
}

// Topics seed the classifier's vocabulary. Without them a chapter matches only
// its own title, which is thin evidence.
func TestParseOutlineSeedsTopicsFromTitle(t *testing.T) {
	pages := []Page{{Number: 1, Text: `
1 The Linear Model 1
2 Least Squares Estimation 17
`}}
	got, err := ParseOutline(context.Background(), pages, nil)
	if err != nil {
		t.Fatal(err)
	}
	topics := got.Chapters[1].Topics
	if len(topics) == 0 {
		t.Fatal("no topics seeded")
	}
	for _, tp := range topics {
		if tp == "the" || tp == "of" {
			t.Fatalf("stop word kept as a topic: %q", tp)
		}
	}
}

func TestCleanTitle(t *testing.T) {
	cases := map[string]string{
		"Least Squares Estimation . . . . . . 17": "Least Squares Estimation",
		"Model Diagnostics                    73": "Model Diagnostics",
		"  Inference and Hypothesis Testing  ":    "Inference and Hypothesis Testing",
		"Generalized Linear Models·······98":      "Generalized Linear Models",
	}
	for in, want := range cases {
		if got := cleanTitle(in); got != want {
			t.Errorf("cleanTitle(%q) = %q, want %q", in, got, want)
		}
	}
}
