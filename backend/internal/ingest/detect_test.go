package ingest

import (
	"archive/zip"
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestDetectFormat(t *testing.T) {
	cases := []struct {
		name string
		data []byte
		want Format
	}{
		{"pdf", []byte("%PDF-1.7\n..."), FormatPDF},
		{"jpeg", []byte("\xFF\xD8\xFF\xE0 jfif"), FormatImage},
		{"png", []byte("\x89PNG\r\n\x1a\n...."), FormatImage},
		{"tiff le", []byte("II*\x00...."), FormatImage},
		{"heic from an iPhone", append([]byte{0, 0, 0, 24}, []byte("ftypheic....")...), FormatImage},
		{"plain text", []byte("Question 1. Derive the estimator."), FormatUnknown},
		{"empty", []byte{}, FormatUnknown},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := DetectFormat(tc.data); got != tc.want {
				t.Errorf("DetectFormat = %q, want %q", got, tc.want)
			}
		})
	}
}

// A photo named "scan.pdf" is still a photo. Trusting the extension would send
// it to the PDF parser, which would fail confusingly instead of routing to OCR.
func TestDetectFormatIgnoresTheFilename(t *testing.T) {
	jpeg := []byte("\xFF\xD8\xFF\xE0 this is really a jpeg")
	if got := DetectFormat(jpeg); got != FormatImage {
		t.Errorf("got %q, want image regardless of what it might be called", got)
	}
}

// buildDOCX writes a minimal but genuine .docx: a zip containing the one entry
// that matters.
func buildDOCX(t *testing.T, paragraphs ...string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)

	w, err := zw.Create("word/document.xml")
	if err != nil {
		t.Fatal(err)
	}
	var body strings.Builder
	body.WriteString(`<?xml version="1.0"?><w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:body>`)
	for _, p := range paragraphs {
		body.WriteString("<w:p><w:r><w:t>" + p + "</w:t></w:r></w:p>")
	}
	body.WriteString(`</w:body></w:document>`)

	if _, err := w.Write([]byte(body.String())); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestDetectDOCX(t *testing.T) {
	data := buildDOCX(t, "Question 1. Derive the estimator.")
	if got := DetectFormat(data); got != FormatDOCX {
		t.Errorf("DetectFormat = %q, want docx", got)
	}
}

func TestParseDOCX(t *testing.T) {
	data := buildDOCX(t,
		"Question 1. Derive the ordinary least squares estimator and state the assumptions. (12 marks)",
		"Question 2. State and prove the Gauss-Markov theorem. (14 marks)",
	)

	doc, err := DOCXParser{}.Parse(context.Background(), bytes.NewReader(data), "paper.docx")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if doc.ParsedBy != "docx-go" {
		t.Errorf("ParsedBy = %q", doc.ParsedBy)
	}
	if !strings.Contains(doc.Pages[0].Text, "Gauss-Markov") {
		t.Errorf("text is missing expected content: %q", doc.Pages[0].Text)
	}

	// The end-to-end point: a Word paper must segment like a PDF one, since a
	// paragraph break is what the segmenter keys on.
	questions := SegmentQuestions(doc.Pages)
	if len(questions) != 2 {
		t.Fatalf("segmented %d questions from a .docx, want 2", len(questions))
	}
	if questions[0].Marks == nil || *questions[0].Marks != 12 {
		t.Errorf("marks = %v, want 12", questions[0].Marks)
	}
}

// A .docx with no extractable text, one holding only images, must be reported
// as needing OCR rather than as an empty success.
func TestProbeDOCXWithoutText(t *testing.T) {
	p := ProbeDOCX(buildDOCX(t, "x"))
	if p.HasTextLayer {
		t.Error("HasTextLayer = true for a document with almost no text")
	}
	if !p.LikelyScanned {
		t.Error("LikelyScanned = false, so this would never route to OCR")
	}
}

func TestProbeDocumentRoutesImagesToOCR(t *testing.T) {
	_, p := ProbeDocument([]byte("\xFF\xD8\xFF\xE0 jpeg body"))
	if p.HasTextLayer {
		t.Error("an image was reported as having a text layer")
	}
	if !p.LikelyScanned {
		t.Error("an image must be routed to OCR")
	}
}

// With OCR registered, a photo must reach it; without, the Chain must refuse
// rather than hand back an empty extraction.
func TestChainRoutesImageToOCRWhenRegistered(t *testing.T) {
	_, probe := ProbeDocument([]byte("\xFF\xD8\xFF\xE0 jpeg body"))

	bare := &Chain{}
	bare.Register(PDFParser{})
	bare.Register(DOCXParser{})
	if _, err := bare.Select(probe); err == nil {
		t.Error("a photo was accepted with no OCR parser registered")
	}

	withOCR := &Chain{}
	withOCR.Register(PDFParser{})
	withOCR.Register(DOCXParser{})
	withOCR.Register(NewOCRParser("http://ocr:50052"))

	got, err := withOCR.Select(probe)
	if err != nil {
		t.Fatalf("Select() error = %v", err)
	}
	if got.Capabilities().Name != "tesseract-ocr" {
		t.Errorf("photo routed to %q, want the OCR parser", got.Capabilities().Name)
	}
}
