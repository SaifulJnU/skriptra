package ingest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/ledongthuc/pdf"
)

// PDFParser is the in-process fallback for PDFs.
//
// Pure Go, no cgo, so it builds anywhere Go builds. It is genuinely limited:
// on PDFs that do not encode their own spaces, which includes most LaTeX
// output, it returns every word run together. A real exam paper came back as
// "Exercise1:(6Points)Decideforeachofthefollowing".
//
// Reconstructing the gaps from geometry was tried and abandoned: the library
// reports X, W and FontSize as zero for exactly the files that need it, and
// ordering rows by Y reversed the reading order on files that did not. It
// therefore declares Quality 1, and the Chain prefers the pdftotext parser
// whenever the document service is running.
//
// It stays because a deployment without that service should still handle
// ordinary PDFs rather than refuse everything.
//
// It declares no OCR capability, so the Chain will refuse a scan rather than
// hand back empty text that would silently poison the index.
type PDFParser struct{}

func (PDFParser) Capabilities() Capabilities {
	return Capabilities{
		Name:      "pdf-go",
		Formats:   []Format{FormatPDF},
		TextLayer: true,
		OCR:       false,
		Layout:    false,
		Formulas:  false,
		Local:     true,
		// Low: it loses spacing on PDFs that do not encode them, which is most
		// LaTeX output. Adequate as a fallback when no sidecar is running.
		Quality: 1,
	}
}

func (p PDFParser) Parse(ctx context.Context, r io.Reader, filename string) (*ParsedDocument, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}

	reader, err := pdf.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filename, err)
	}

	total := reader.NumPage()
	doc := &ParsedDocument{
		PageCount: total,
		Pages:     make([]Page, 0, total),
		ParsedBy:  "pdf-go",
	}

	for i := 1; i <= total; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		page := reader.Page(i)
		if page.V.IsNull() {
			doc.Pages = append(doc.Pages, Page{Number: i})
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			// One malformed page must not lose the other twenty. Record it
			// empty and carry on; Probe will notice if the whole document
			// came back blank.
			doc.Pages = append(doc.Pages, Page{Number: i})
			continue
		}
		doc.Pages = append(doc.Pages, Page{Number: i, Text: text})
	}

	return doc, nil
}

// ProbePDF inspects a document cheaply to decide which parser it needs.
//
// The signal that matters is whether a usable text layer exists. A scan
// produces a PDF with pages and no extractable text, and that is precisely the
// case that must route to OCR rather than be indexed as nothing.
func ProbePDF(buf []byte) Probe {
	p := Probe{}

	reader, err := pdf.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return Probe{LikelyScanned: true}
	}
	p.PageCount = reader.NumPage()

	// Sampling the first few pages is enough and keeps upload responsive on a
	// 200-page textbook.
	sample := p.PageCount
	if sample > 3 {
		sample = 3
	}

	var extracted strings.Builder
	for i := 1; i <= sample; i++ {
		page := reader.Page(i)
		if page.V.IsNull() {
			continue
		}
		if text, err := page.GetPlainText(nil); err == nil {
			extracted.WriteString(text)
		}
	}

	text := strings.TrimSpace(extracted.String())
	// A handful of stray characters is not a text layer. The threshold is
	// per sampled page so a short cover page does not skew the verdict.
	p.HasTextLayer = len(text) > 40*sample
	p.LikelyScanned = !p.HasTextLayer
	p.ContainsMath = strings.ContainsAny(text, "∫∑√≤≥≠∂αβγσμθλ")

	return p
}
