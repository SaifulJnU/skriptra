// Package ingest turns uploaded documents into indexable content.
//
// # Polyglot by necessity, not by default
//
// Document parsing is the one place in Skriptra where another language earns
// its keep. Modern OCR and layout analysis (Surya, PaddleOCR, docTR) have no Go
// equivalent, while everything else, embeddings, generation, retrieval,
// orchestration, is HTTP, SQL or ordinary application code that Go handles
// natively.
//
// So parsing is defined as an interface with a capability declaration, and a
// Chain picks the right implementation *per document* rather than per
// deployment. A digitally-generated PDF takes the fast in-process Go path. A
// scan takes whatever implementation can actually read it, today nothing, and
// later a Python sidecar over gRPC, registered exactly like any other parser.
//
// The point is that "add Python" is a new file and a compose entry, not a
// refactor. Nothing in the ingestion pipeline knows or cares which
// implementation ran.
package ingest

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sort"
	"sync"
)

// Page is one page of extracted text, keyed by the 1-indexed page number that
// citations point at.
type Page struct {
	Number int
	Text   string
	Width  float64
	Height float64
}

// Block is a positioned run of text. Only parsers that report Layout
// capability populate it; it is what bounding-box citation highlighting will
// need in phase 2, and it is defined now so adding that does not change this
// type.
type Block struct {
	Page           int
	Text           string
	X0, Y0, X1, Y1 float64
}

// ParsedDocument is the parser output contract.
type ParsedDocument struct {
	Pages     []Page
	Blocks    []Block
	PageCount int
	Language  string

	// ParsedBy names the implementation that produced this, so a bad
	// extraction can be traced to a parser rather than guessed at.
	ParsedBy string
}

// Capabilities describes what an implementation can actually do. The Chain
// dispatches on this, which is why it is data rather than a type switch.
type Capabilities struct {
	Name string

	// Formats this parser can read. Capability and quality are not enough on
	// their own: with three parsers registered, the highest-quality one was
	// selected for a PDF purely because it read text well, and it happened to
	// be the Word reader. Format is the first filter for that reason.
	Formats []Format

	// TextLayer: can extract an existing text layer. Every parser can.
	TextLayer bool
	// OCR: can read pages with no text layer (scans, photographs).
	OCR bool
	// Layout: reports positioned blocks, not just page text.
	Layout bool
	// Formulas: can render mathematics to LaTeX rather than mangled text.
	Formulas bool

	// Local is false for out-of-process parsers, so the scheduler can prefer
	// the cheap in-process path when quality is equal.
	Local bool

	// Quality ranks extraction fidelity, higher is better.
	//
	// Cost alone is the wrong basis for choosing. The pure-Go PDF reader is
	// free and in-process, and on a real LaTeX exam paper it returned
	// "Exercise1:(6Points)Decideforeach" with every space missing and no
	// position data to reconstruct them from. Cheap and wrong is not a
	// trade worth making, so fidelity is ranked first and cost breaks ties.
	Quality int
}

// Probe is what the Chain inspects before choosing. Cheap to compute: it only
// requires reading the first page or two.
type Probe struct {
	// Format is decided from the file's contents, never its extension.
	Format        Format
	HasTextLayer  bool
	LikelyScanned bool
	ContainsMath  bool
	PageCount     int
}

// DocumentParser extracts text and page structure from a document.
//
// Implementations must be safe for concurrent use and must respect context
// cancellation, an out-of-process parser that ignores it will pin an
// ingestion worker on a document that no longer matters.
type DocumentParser interface {
	Parse(ctx context.Context, r io.Reader, filename string) (*ParsedDocument, error)
	Capabilities() Capabilities
}

var (
	// ErrNoCapableParser means the document needs something no registered
	// parser provides, in practice, a scan with no OCR parser installed.
	// It is deliberately distinct from a parse failure: the fix is deploying
	// a capable parser, not retrying.
	ErrNoCapableParser = errors.New("no registered parser can handle this document")

	// ErrUnreadable means a capable parser tried and failed.
	ErrUnreadable = errors.New("document could not be read")
)

// Chain selects a parser per document.
type Chain struct {
	mu      sync.RWMutex
	parsers []DocumentParser
}

// Register adds an implementation. Call order does not matter; selection is by
// capability and cost, not registration order.
func (c *Chain) Register(p DocumentParser) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.parsers = append(c.parsers, p)
}

// Select returns the cheapest parser capable of handling the probed document.
//
// Preference order: satisfy the document's requirements first, then prefer an
// in-process parser over an out-of-process one. A scanned page needs OCR and
// there is no way around the sidecar; a clean digital PDF must never pay for
// one.
func (c *Chain) Select(p Probe) (DocumentParser, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	needsOCR := !p.HasTextLayer || p.LikelyScanned

	candidates := make([]DocumentParser, 0, len(c.parsers))
	for _, parser := range c.parsers {
		caps := parser.Capabilities()
		if !caps.handles(p.Format) {
			continue
		}
		if needsOCR && !caps.OCR {
			continue
		}
		if !needsOCR && !caps.TextLayer {
			continue
		}
		candidates = append(candidates, parser)
	}

	if len(candidates) == 0 {
		if needsOCR {
			return nil, fmt.Errorf("%w: document has no text layer and no OCR-capable parser is registered", ErrNoCapableParser)
		}
		return nil, fmt.Errorf("%w: no parser accepts format %q", ErrNoCapableParser, p.Format)
	}

	// Best first, then cheapest. Fidelity decides, because text that came back
	// wrong poisons every downstream stage silently; cost only breaks ties.
	sort.SliceStable(candidates, func(i, j int) bool {
		a, b := candidates[i].Capabilities(), candidates[j].Capabilities()
		if a.Quality != b.Quality {
			return a.Quality > b.Quality
		}
		if a.Local != b.Local {
			return a.Local
		}
		return weight(a) < weight(b)
	})

	return candidates[0], nil
}

// Parse probes, selects and runs, so callers never touch parser selection.
func (c *Chain) Parse(ctx context.Context, r io.Reader, filename string, p Probe) (*ParsedDocument, error) {
	parser, err := c.Select(p)
	if err != nil {
		return nil, err
	}
	doc, err := parser.Parse(ctx, r, filename)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrUnreadable, parser.Capabilities().Name, err)
	}
	if doc.ParsedBy == "" {
		doc.ParsedBy = parser.Capabilities().Name
	}
	return doc, nil
}

// Registered lists the parser names currently available, for /healthz and for
// telling a user why their scan was rejected.
func (c *Chain) Registered() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	names := make([]string, 0, len(c.parsers))
	for _, p := range c.parsers {
		names = append(names, p.Capabilities().Name)
	}
	sort.Strings(names)
	return names
}

// handles reports whether this parser accepts the format. An empty list means
// the parser did not declare one, which is treated as accepting anything so an
// adapter written before this field cannot silently stop being selected.
func (c Capabilities) handles(f Format) bool {
	if len(c.Formats) == 0 || f == "" {
		return true
	}
	for _, allowed := range c.Formats {
		if allowed == f {
			return true
		}
	}
	return false
}

func weight(c Capabilities) int {
	n := 0
	for _, on := range []bool{c.OCR, c.Layout, c.Formulas} {
		if on {
			n++
		}
	}
	return n
}

// PageLimiter is implemented by parsers that can stop after the first N pages.
//
// Optional rather than part of DocumentParser: a photograph has one page and a
// Word file has no page concept until it is rendered, so forcing every
// implementation to carry the parameter would add a field most of them ignore.
type PageLimiter interface {
	ParsePages(ctx context.Context, r io.Reader, filename string, maxPages int) (*ParsedDocument, error)
}

// ParseFirstPages reads only the front of a document.
//
// For the one job that needs it: finding a chapter list, which lives on a
// contents page near the front. Parsers that cannot stop early fall back to a
// full parse, so the caller always gets a result.
func (c *Chain) ParseFirstPages(ctx context.Context, r io.Reader, filename string, p Probe, maxPages int) (*ParsedDocument, error) {
	parser, err := c.Select(p)
	if err != nil {
		return nil, err
	}
	if limited, ok := parser.(PageLimiter); ok {
		doc, err := limited.ParsePages(ctx, r, filename, maxPages)
		if err != nil {
			return nil, fmt.Errorf("document could not be read: %w", err)
		}
		return doc, nil
	}
	return c.Parse(ctx, r, filename, p)
}
