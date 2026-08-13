package ingest

import (
	"bytes"
	"strings"
)

// Format is the container a document arrived in, decided by content rather than
// by the filename. A phone that names a photo `scan.pdf` should still be
// treated as an image, and an extension is a claim by whoever uploaded it.
type Format string

const (
	FormatPDF     Format = "pdf"
	FormatDOCX    Format = "docx"
	FormatImage   Format = "image"
	FormatUnknown Format = "unknown"
)

// magic numbers, checked in order of how specific they are.
var signatures = []struct {
	prefix []byte
	format Format
}{
	{[]byte("%PDF"), FormatPDF},
	{[]byte("\xFF\xD8\xFF"), FormatImage},      // jpeg
	{[]byte("\x89PNG\r\n\x1a\n"), FormatImage}, // png
	{[]byte("GIF87a"), FormatImage},
	{[]byte("GIF89a"), FormatImage},
	{[]byte("BM"), FormatImage},      // bmp
	{[]byte("II*\x00"), FormatImage}, // tiff, little endian
	{[]byte("MM\x00*"), FormatImage}, // tiff, big endian
}

// DetectFormat identifies a file from its contents.
func DetectFormat(buf []byte) Format {
	for _, sig := range signatures {
		if bytes.HasPrefix(buf, sig.prefix) {
			return sig.format
		}
	}

	// A .docx is a zip, and so is every other Office format, so the zip magic
	// alone is not enough. The word/ directory is what distinguishes it.
	if bytes.HasPrefix(buf, []byte("PK\x03\x04")) {
		if bytes.Contains(buf, []byte("word/document.xml")) {
			return FormatDOCX
		}
		// A large .docx may not carry its central directory within the window
		// searched above; fall back to the mimetype marker Office writes early.
		if bytes.Contains(buf, []byte("wordprocessingml")) {
			return FormatDOCX
		}
	}

	// HEIC/HEIF, which is what an iPhone produces by default. The brand appears
	// after a 4-byte box length.
	if len(buf) > 12 && bytes.Equal(buf[4:8], []byte("ftyp")) {
		brand := string(buf[8:12])
		if strings.HasPrefix(brand, "heic") || strings.HasPrefix(brand, "heif") ||
			strings.HasPrefix(brand, "mif1") || strings.HasPrefix(brand, "avif") {
			return FormatImage
		}
	}

	return FormatUnknown
}

// ProbeDocument inspects any supported format and reports what it needs.
//
// One entry point, so callers never have to know which prober to run. An image
// always needs OCR: there is no such thing as a photograph with a text layer.
func ProbeDocument(buf []byte) (Format, Probe) {
	format := DetectFormat(buf)

	var p Probe
	switch format {
	case FormatPDF:
		p = ProbePDF(buf)
	case FormatDOCX:
		p = ProbeDOCX(buf)
	case FormatImage:
		p = Probe{PageCount: 1, HasTextLayer: false, LikelyScanned: true}
	}
	// Carried on the probe so the Chain can filter on it without being told
	// separately, which is how a Word reader came to be chosen for a PDF.
	p.Format = format
	return format, p
}

// SupportedFormats lists what upload accepts, for error messages and for the
// client's accept attribute.
func SupportedFormats() []string {
	return []string{"PDF", "Word (.docx)", "JPEG", "PNG", "HEIC", "TIFF"}
}
