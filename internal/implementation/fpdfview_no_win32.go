//go:build !windows
// +build !windows

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
	return nil, pdfium_errors.ErrWindowsUnsupported
}

// FPDF_RenderPage renders contents of a page to a device (screen, bitmap, or printer).
// This feature does not work on multi-threaded usage as you will need to give a device handle.
// Windows only!
func (p *PdfiumImplementation) FPDF_RenderPage(request *requests.FPDF_RenderPage) (*responses.FPDF_RenderPage, error) {
	return nil, pdfium_errors.ErrWindowsUnsupported
}
