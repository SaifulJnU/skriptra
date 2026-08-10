package ingest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"time"
)

// OCRParser routes documents with no text layer to the OCR sidecar.
//
// This is the implementation the Chain was built for. Registering it is the
// whole integration: nothing in the pipeline changes, because nothing in the
// pipeline knows which parser ran. Photographs and scans start working the
// moment the service is reachable, and stop working cleanly when it is not.
type OCRParser struct {
	baseURL string
	http    *http.Client
}

// NewOCRParser returns nil when no address is configured, so callers can
// register unconditionally and get the correct behaviour either way.
func NewOCRParser(baseURL string) *OCRParser {
	if baseURL == "" {
		return nil
	}
	return &OCRParser{
		baseURL: baseURL,
		http: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			},
			// OCR on a 40-page scan legitimately takes minutes on CPU. The
			// ceiling is high because the alternative is failing a document
			// that would have succeeded; cancellation is the caller's context.
			Timeout: 15 * time.Minute,
		},
	}
}

func (*OCRParser) Capabilities() Capabilities {
	return Capabilities{
		Name:      "tesseract-ocr",
		TextLayer: true,
		OCR:       true,
		Layout:    false,
		Formulas:  false,
		// Out of process, so the Chain prefers the in-process Go parser
		// whenever a document does not actually need OCR.
		Local: false,
	}
}

type ocrResponse struct {
	Pages []struct {
		Number int    `json:"number"`
		Text   string `json:"text"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"pages"`
	PageCount           int    `json:"pageCount"`
	ParsedBy            string `json:"parsedBy"`
	CharactersExtracted int    `json:"charactersExtracted"`
}

func (o *OCRParser) Parse(ctx context.Context, r io.Reader, filename string) (*ParsedDocument, error) {
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)

	part, err := mw.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, r); err != nil {
		return nil, err
	}
	if err := mw.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+"/ocr", &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	res, err := o.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ocr service unreachable: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(res.Body, 512))
		return nil, fmt.Errorf("ocr service returned %d: %s", res.StatusCode, bytes.TrimSpace(msg))
	}

	var parsed ocrResponse
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("decode ocr response: %w", err)
	}

	// A successful call that recovered almost nothing is a failure worth
	// naming. Indexing a blank document would put an entry in the corpus that
	// silently matches nothing, and the operator would have no idea why.
	if parsed.CharactersExtracted < 40 {
		return nil, fmt.Errorf("%w: OCR recovered only %d characters from %s, the image may be too blurred or too low contrast",
			ErrUnreadable, parsed.CharactersExtracted, filename)
	}

	doc := &ParsedDocument{
		PageCount: parsed.PageCount,
		Pages:     make([]Page, 0, len(parsed.Pages)),
		ParsedBy:  "tesseract-ocr",
	}
	for _, p := range parsed.Pages {
		doc.Pages = append(doc.Pages, Page{
			Number: p.Number,
			Text:   p.Text,
			Width:  float64(p.Width),
			Height: float64(p.Height),
		})
	}
	return doc, nil
}

// Healthy reports whether the sidecar is reachable, so the API can tell a user
// that photo upload is unavailable rather than accepting a file and failing
// later in a worker they cannot see.
func (o *OCRParser) Healthy(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, o.baseURL+"/healthz", nil)
	if err != nil {
		return false
	}
	res, err := o.http.Do(req)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	return res.StatusCode == http.StatusOK
}
