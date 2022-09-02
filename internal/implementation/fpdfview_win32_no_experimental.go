//go:build windows && !pdfium_experimental
// +build windows,!pdfium_experimental

package implementation

import (
	pdfium_errors "github.com/michalderdak/go-pdfium/errors"

	"github.com/michalderdak/go-pdfium/requests"
	"github.com/michalderdak/go-pdfium/responses"
)

// FPDF_SetPrintMode sets printing mode when printing on Windows.
// Experimental API.
// Windows only!
func (p *PdfiumImplementation) FPDF_SetPrintMode(request *requests.FPDF_SetPrintMode) (*responses.FPDF_SetPrintMode, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
