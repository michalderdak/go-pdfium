package requests

import "github.com/michalderdak/go-pdfium/references"

type GetDestInfo struct {
	Document references.FPDF_DOCUMENT
	Dest     references.FPDF_DEST
}
