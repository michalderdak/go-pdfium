package implementation

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"errors"
	"github.com/google/uuid"

	pdfium_errors "github.com/klippa-app/go-pdfium/errors"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

// FPDF_LoadPage loads a page and returns a reference.
func (p *PdfiumImplementation) FPDF_LoadPage(request *requests.FPDF_LoadPage) (*responses.FPDF_LoadPage, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	pageObject := C.FPDF_LoadPage(nativeDoc.currentDoc, C.int(request.Index))
	if pageObject == nil {
		return nil, pdfium_errors.ErrPage
	}

	pageRef := uuid.New()
	nativePage := &NativePage{
		page:        pageObject,
		index:       request.Index,
		nativeRef:   references.FPDF_PAGE(pageRef.String()),
		documentRef: nativeDoc.nativeRef,
	}

	nativeDoc.pageRefs[nativePage.nativeRef] = nativePage
	p.pageRefs[nativePage.nativeRef] = nativePage

	return &responses.FPDF_LoadPage{
		Page: nativePage.nativeRef,
	}, nil
}

// FPDF_ClosePage unloads a page by reference.
func (p *PdfiumImplementation) FPDF_ClosePage(request *requests.FPDF_ClosePage) (*responses.FPDF_ClosePage, error) {
	p.Lock()
	defer p.Unlock()

	pageRef, err := p.getNativePage(request.Page)
	if err != nil {
		return nil, err
	}

	pageRef.Close()
	delete(p.pageRefs, request.Page)

	// Remove page reference from document.
	nativeDoc, err := p.getNativeDocument(pageRef.documentRef)
	if err != nil {
		return nil, err
	}
	delete(nativeDoc.pageRefs, request.Page)

	return &responses.FPDF_ClosePage{}, nil
}

// FPDF_GetFileVersion returns the version of the PDF file.
func (p *PdfiumImplementation) FPDF_GetFileVersion(request *requests.FPDF_GetFileVersion) (*responses.FPDF_GetFileVersion, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	fileVersion := C.int(0)

	success := C.FPDF_GetFileVersion(nativeDoc.currentDoc, &fileVersion)
	if int(success) == 0 {
		return nil, errors.New("could not get file version")
	}

	return &responses.FPDF_GetFileVersion{
		FileVersion: int(fileVersion),
	}, nil
}

// FPDF_GetDocPermissions returns the permissions of the PDF.
func (p *PdfiumImplementation) FPDF_GetDocPermissions(request *requests.FPDF_GetDocPermissions) (*responses.FPDF_GetDocPermissions, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	permissions := C.FPDF_GetDocPermissions(nativeDoc.currentDoc)

	docPermissions := &responses.FPDF_GetDocPermissions{
		DocPermissions: uint32(permissions),
	}

	PrintDocument := uint32(1 << 2)
	ModifyContents := uint32(1 << 3)
	CopyOrExtractText := uint32(1 << 4)
	AddOrModifyTextAnnotations := uint32(1 << 5)
	FillInExistingInteractiveFormFields := uint32(1 << 8)
	ExtractTextAndGraphics := uint32(1 << 9)
	AssembleDocument := uint32(1 << 10)
	PrintDocumentAsFaithfulDigitalCopy := uint32(1 << 11)

	hasPermission := func(permission uint32) bool {
		if docPermissions.DocPermissions&permission > 0 {
			return true
		}

		return false
	}

	docPermissions.PrintDocument = hasPermission(PrintDocument)
	docPermissions.ModifyContents = hasPermission(ModifyContents)
	docPermissions.CopyOrExtractText = hasPermission(CopyOrExtractText)
	docPermissions.AddOrModifyTextAnnotations = hasPermission(AddOrModifyTextAnnotations)
	docPermissions.FillInInteractiveFormFields = hasPermission(AddOrModifyTextAnnotations)
	docPermissions.FillInExistingInteractiveFormFields = hasPermission(FillInExistingInteractiveFormFields)
	docPermissions.ExtractTextAndGraphics = hasPermission(ExtractTextAndGraphics)
	docPermissions.AssembleDocument = hasPermission(AssembleDocument)
	docPermissions.PrintDocumentAsFaithfulDigitalCopy = hasPermission(PrintDocumentAsFaithfulDigitalCopy)

	// Calculated permissions
	docPermissions.CreateOrModifyInteractiveFormFields = docPermissions.ModifyContents && docPermissions.AddOrModifyTextAnnotations

	return docPermissions, nil
}

// FPDF_GetSecurityHandlerRevision returns the revision number of security handlers of the file.
func (p *PdfiumImplementation) FPDF_GetSecurityHandlerRevision(request *requests.FPDF_GetSecurityHandlerRevision) (*responses.FPDF_GetSecurityHandlerRevision, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	securityHandlerRevision := C.FPDF_GetSecurityHandlerRevision(nativeDoc.currentDoc)

	return &responses.FPDF_GetSecurityHandlerRevision{
		SecurityHandlerRevision: int(securityHandlerRevision),
	}, nil
}

// FPDF_GetPageCount counts the amount of pages.
func (p *PdfiumImplementation) FPDF_GetPageCount(request *requests.FPDF_GetPageCount) (*responses.FPDF_GetPageCount, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	return &responses.FPDF_GetPageCount{
		PageCount: int(C.FPDF_GetPageCount(nativeDoc.currentDoc)),
	}, nil
}
