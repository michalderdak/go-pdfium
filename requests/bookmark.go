package requests

import "github.com/michalderdak/go-pdfium/references"

type GetBookmarks struct {
	Document references.FPDF_DOCUMENT
}
