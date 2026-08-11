package ingest

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// DOCXParser extracts text from Word documents.
//
// A .docx is a zip containing XML, so this needs nothing beyond the standard
// library. Students hand around lecture notes and question banks as Word files
// at least as often as PDFs, and refusing them would have been a gap with no
// technical justification behind it.
//
// There is no page concept in a .docx: pagination is decided by the renderer,
// not stored in the file. Rather than invent page numbers, every block is
// reported as page 1 and citations point at the document. Fabricating a page
// that Word might have produced would put a precise-looking number on a guess.
type DOCXParser struct{}

func (DOCXParser) Capabilities() Capabilities {
	return Capabilities{
		Name:      "docx-go",
		Formats:   []Format{FormatDOCX},
		TextLayer: true,
		OCR:       false,
		Layout:    false,
		Formulas:  false,
		Local:     true,
		// The words are literally in the XML, so nothing is being inferred.
		Quality: 5,
	}
}

func (p DOCXParser) Parse(ctx context.Context, r io.Reader, filename string) (*ParsedDocument, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}

	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, fmt.Errorf("open %s as docx: %w", filename, err)
	}

	var doc *zip.File
	for _, f := range zr.File {
		if f.Name == "word/document.xml" {
			doc = f
			break
		}
	}
	if doc == nil {
		return nil, fmt.Errorf("%s is a zip but not a Word document: word/document.xml is missing", filename)
	}

	rc, err := doc.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	text, err := extractDOCXText(ctx, rc)
	if err != nil {
		return nil, err
	}

	return &ParsedDocument{
		PageCount: 1,
		Pages:     []Page{{Number: 1, Text: text}},
		ParsedBy:  "docx-go",
	}, nil
}

// extractDOCXText walks the document body and reconstructs readable text.
//
// Streamed rather than unmarshalled into a struct, because Word's schema is
// large, deeply nested and full of formatting elements that carry no content.
// Only four element names matter here.
func extractDOCXText(ctx context.Context, r io.Reader) (string, error) {
	dec := xml.NewDecoder(r)
	var out strings.Builder
	var inText bool

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("parse document.xml: %w", err)
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "t": // w:t, a run of literal text
				inText = true
			case "tab":
				out.WriteByte(' ')
			case "br", "cr":
				out.WriteByte('\n')
			}

		case xml.CharData:
			if inText {
				out.Write(t)
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "t":
				inText = false
			case "p":
				// A paragraph break is what the segmenter keys on: exam
				// questions start at the beginning of a line.
				out.WriteByte('\n')
			}
		}
	}

	return out.String(), nil
}

// ProbeDOCX reports what a Word document offers.
//
// A .docx always carries a text layer if it carries anything, so there is no
// scanned equivalent. A document containing only embedded images extracts
// nothing, and that is reported honestly as no text layer rather than as an
// empty success.
func ProbeDOCX(buf []byte) Probe {
	p := Probe{PageCount: 1}

	doc, err := DOCXParser{}.Parse(context.Background(), bytes.NewReader(buf), "probe.docx")
	if err != nil {
		return Probe{LikelyScanned: true}
	}
	text := strings.TrimSpace(doc.Pages[0].Text)
	p.HasTextLayer = len(text) > 40
	p.LikelyScanned = !p.HasTextLayer
	p.ContainsMath = strings.ContainsAny(text, "∫∑√≤≥≠∂αβγσμθλ")
	return p
}
