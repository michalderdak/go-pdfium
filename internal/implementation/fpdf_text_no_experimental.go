//go:build !pdfium_experimental
// +build !pdfium_experimental

package implementation

import (
	pdfium_errors "github.com/michalderdak/go-pdfium/errors"
	"github.com/michalderdak/go-pdfium/requests"
	"github.com/michalderdak/go-pdfium/responses"
)

// FPDFText_GetFontInfo returns the font name and flags of a particular character.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetFontInfo(request *requests.FPDFText_GetFontInfo) (*responses.FPDFText_GetFontInfo, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetFontWeight returns the font weight of a particular character.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetFontWeight(request *requests.FPDFText_GetFontWeight) (*responses.FPDFText_GetFontWeight, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetTextRenderMode returns the text rendering mode of character.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetTextRenderMode(request *requests.FPDFText_GetTextRenderMode) (*responses.FPDFText_GetTextRenderMode, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetFillColor returns the fill color of a particular character.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetFillColor(request *requests.FPDFText_GetFillColor) (*responses.FPDFText_GetFillColor, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetStrokeColor returns the stroke color of a particular character.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetStrokeColor(request *requests.FPDFText_GetStrokeColor) (*responses.FPDFText_GetStrokeColor, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetCharAngle returns the character rotation angle.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetCharAngle(request *requests.FPDFText_GetCharAngle) (*responses.FPDFText_GetCharAngle, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetLooseCharBox returns a "loose" bounding box of a particular character, i.e., covering
// the entire glyph bounds, without taking the actual glyph shape into
// account. All positions are measured in PDF "user space".
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetLooseCharBox(request *requests.FPDFText_GetLooseCharBox) (*responses.FPDFText_GetLooseCharBox, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_GetMatrix returns the effective transformation matrix for a particular character.
// All positions are measured in PDF "user space".
// Experimental API.
func (p *PdfiumImplementation) FPDFText_GetMatrix(request *requests.FPDFText_GetMatrix) (*responses.FPDFText_GetMatrix, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFLink_GetTextRange returns the start char index and char count for a link.
// Experimental API.
func (p *PdfiumImplementation) FPDFLink_GetTextRange(request *requests.FPDFLink_GetTextRange) (*responses.FPDFLink_GetTextRange, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
