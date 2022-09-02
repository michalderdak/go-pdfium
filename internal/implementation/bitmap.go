package implementation

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/google/uuid"
	"github.com/michalderdak/go-pdfium/references"
)

func (p *PdfiumImplementation) registerBitmap(bitmap C.FPDF_BITMAP) *BitmapHandle {
	ref := uuid.New()
	handle := &BitmapHandle{
		handle:    bitmap,
		nativeRef: references.FPDF_BITMAP(ref.String()),
	}

	p.bitmapRefs[handle.nativeRef] = handle

	return handle
}
