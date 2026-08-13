package ingest

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

// fakeParser stands in for a real implementation so the selection rules can be
// tested without a PDF or a running sidecar.
type fakeParser struct{ caps Capabilities }

func (f fakeParser) Capabilities() Capabilities { return f.caps }
func (f fakeParser) Parse(_ context.Context, _ io.Reader, _ string) (*ParsedDocument, error) {
	return &ParsedDocument{PageCount: 1, ParsedBy: f.caps.Name}, nil
}

var (
	goNative  = fakeParser{Capabilities{Name: "go-fitz", TextLayer: true, Local: true}}
	pySidecar = fakeParser{Capabilities{
		Name: "python-ocr", TextLayer: true, OCR: true, Layout: true, Formulas: true, Local: false,
	}}
)

func TestSelectPrefersLocalParserForDigitalPDF(t *testing.T) {
	c := &Chain{}
	// Registered in the "expensive first" order on purpose: selection must not
	// depend on registration order.
	c.Register(pySidecar)
	c.Register(goNative)

	got, err := c.Select(Probe{HasTextLayer: true})
	if err != nil {
		t.Fatalf("Select() error = %v", err)
	}
	if name := got.Capabilities().Name; name != "go-fitz" {
		t.Errorf("digital PDF routed to %q, want the in-process parser %q", name, "go-fitz")
	}
}

func TestSelectRoutesScanToOCRParser(t *testing.T) {
	c := &Chain{}
	c.Register(goNative)
	c.Register(pySidecar)

	got, err := c.Select(Probe{HasTextLayer: false, LikelyScanned: true})
	if err != nil {
		t.Fatalf("Select() error = %v", err)
	}
	if name := got.Capabilities().Name; name != "python-ocr" {
		t.Errorf("scan routed to %q, want an OCR-capable parser", name)
	}
}

// The Go-only MVP state: a scan arrives and nothing can read it. That must be a
// distinct, actionable error rather than a confusing empty extraction.
func TestSelectFailsLoudlyWhenScanArrivesWithoutOCR(t *testing.T) {
	c := &Chain{}
	c.Register(goNative)

	_, err := c.Select(Probe{HasTextLayer: false, LikelyScanned: true})
	if !errors.Is(err, ErrNoCapableParser) {
		t.Fatalf("Select() error = %v, want ErrNoCapableParser", err)
	}
	if !strings.Contains(err.Error(), "no text layer") {
		t.Errorf("error should say why the document was rejected, got: %v", err)
	}
}

func TestParseStampsTheImplementationUsed(t *testing.T) {
	c := &Chain{}
	c.Register(goNative)

	doc, err := c.Parse(context.Background(), strings.NewReader(""), "exam.pdf", Probe{HasTextLayer: true})
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if doc.ParsedBy != "go-fitz" {
		t.Errorf("ParsedBy = %q, want the parser that ran to be recorded", doc.ParsedBy)
	}
}

func TestRegisteredListsParserNames(t *testing.T) {
	c := &Chain{}
	c.Register(pySidecar)
	c.Register(goNative)

	got := c.Registered()
	want := []string{"go-fitz", "python-ocr"}
	if len(got) != len(want) {
		t.Fatalf("Registered() = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Registered()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
