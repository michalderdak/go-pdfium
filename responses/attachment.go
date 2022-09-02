package responses

import "github.com/michalderdak/go-pdfium/enums"

type AttachmentValue struct {
	Key         string
	ValueType   enums.FPDF_OBJECT_TYPE
	StringValue string
}

type Attachment struct {
	Name    string
	Content []byte
	Values  []AttachmentValue
}

type GetAttachments struct {
	Attachments []Attachment
}
