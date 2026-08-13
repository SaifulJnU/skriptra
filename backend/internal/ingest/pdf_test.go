package ingest

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
)

// testdata/sample_exam.pdf is a synthetic Linear Models paper written for this
// repository. It is a real PDF produced by a real renderer, not a hand-built
// fixture, so these tests exercise the actual extraction path rather than a
// convenient approximation of it.
func loadSample(t *testing.T) []byte {
	t.Helper()
	buf, err := os.ReadFile("testdata/sample_exam.pdf")
	if err != nil {
		t.Skipf("sample PDF unavailable: %v", err)
	}
	return buf
}

func TestProbeDetectsTextLayer(t *testing.T) {
	probe := ProbePDF(loadSample(t))

	if !probe.HasTextLayer {
		t.Error("HasTextLayer = false on a digitally generated PDF")
	}
	if probe.LikelyScanned {
		t.Error("LikelyScanned = true on a digitally generated PDF")
	}
	if probe.PageCount < 1 {
		t.Errorf("PageCount = %d, want at least 1", probe.PageCount)
	}
}

// A PDF with no text layer must route to OCR, not be indexed as nothing.
func TestProbeTreatsUnreadableInputAsScanned(t *testing.T) {
	probe := ProbePDF([]byte("%PDF-1.4\nnot actually a pdf"))
	if !probe.LikelyScanned {
		t.Error("LikelyScanned = false for an unreadable file; it must not be treated as text")
	}
	if probe.HasTextLayer {
		t.Error("HasTextLayer = true for an unreadable file")
	}
}

func TestParseExtractsPagesAndText(t *testing.T) {
	doc, err := PDFParser{}.Parse(context.Background(), bytes.NewReader(loadSample(t)), "sample_exam.pdf")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if doc.PageCount < 1 {
		t.Fatalf("PageCount = %d", doc.PageCount)
	}
	if doc.ParsedBy != "pdf-go" {
		t.Errorf("ParsedBy = %q, want the parser to record itself", doc.ParsedBy)
	}

	var all strings.Builder
	for _, p := range doc.Pages {
		all.WriteString(p.Text)
	}
	text := all.String()

	for _, want := range []string{"least squares", "Gauss-Markov", "Poisson"} {
		if !strings.Contains(text, want) {
			t.Errorf("extracted text is missing %q", want)
		}
	}
}

// The end-to-end check that matters: a real PDF in, correctly numbered and
// chapter-attributable questions out.
func TestParseThenSegmentRealPDF(t *testing.T) {
	doc, err := PDFParser{}.Parse(context.Background(), bytes.NewReader(loadSample(t)), "sample_exam.pdf")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	questions := SegmentQuestions(doc.Pages)
	if len(questions) < 5 {
		t.Fatalf("segmented %d questions from the sample paper, want at least 5", len(questions))
	}

	if questions[0].Number != "1" {
		t.Errorf("first question number = %q, want 1", questions[0].Number)
	}
	if questions[0].Marks == nil || *questions[0].Marks != 12 {
		t.Errorf("first question marks = %v, want 12", questions[0].Marks)
	}
	for _, q := range questions {
		if q.SourcePage < 1 {
			t.Errorf("question %s has sourcePage %d; citations depend on this", q.Number, q.SourcePage)
		}
		if strings.Contains(q.Text, "Question ") && strings.Index(q.Text, "Question ") == 0 {
			t.Errorf("question %s still carries its heading: %.40q", q.Number, q.Text)
		}
	}

	// And the segmented output must classify sensibly, which is the whole
	// chain working rather than each half working alone.
	c := NewClassifier(taxonomy, nil)
	got := c.Classify(context.Background(), questions[1].Text)
	if got.ChapterNumber != 2 {
		t.Errorf("Gauss-Markov question classified to chapter %d, want 2", got.ChapterNumber)
	}
}
