package shared_tests

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/michalderdak/go-pdfium/references"
	"github.com/michalderdak/go-pdfium/requests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("fpdf_save", func() {
	BeforeEach(func() {
		Locker.Lock()
	})

	AfterEach(func() {
		Locker.Unlock()
	})

	Context("no document", func() {
		When("is opened", func() {
			It("returns an error when calling FPDF_SaveAsCopy", func() {
				FPDF_SaveAsCopy, err := PdfiumInstance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{})
				Expect(err).To(MatchError("document not given"))
				Expect(FPDF_SaveAsCopy).To(BeNil())
			})

			It("returns an error when calling FPDF_SaveWithVersion", func() {
				FPDF_SaveWithVersion, err := PdfiumInstance.FPDF_SaveWithVersion(&requests.FPDF_SaveWithVersion{})
				Expect(err).To(MatchError("document not given"))
				Expect(FPDF_SaveWithVersion).To(BeNil())
			})
		})
	})

	Context("a normal PDF file", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/test.pdf")
			Expect(err).To(BeNil())

			newDoc, err := PdfiumInstance.FPDF_LoadMemDocument(&requests.FPDF_LoadMemDocument{
				Data: &pdfData,
			})
			Expect(err).To(BeNil())

			doc = newDoc.Document
		})

		AfterEach(func() {
			FPDF_CloseDocument, err := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
				Document: doc,
			})
			Expect(err).To(BeNil())
			Expect(FPDF_CloseDocument).To(Not(BeNil()))
		})

		When("is opened", func() {
			Context("and saved to a byte array", func() {
				It("it returns the correct bytes", func() {
					FPDF_SaveAsCopy, err := PdfiumInstance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
						Document: doc,
					})
					Expect(err).To(BeNil())
					Expect(FPDF_SaveAsCopy).To(Not(BeNil()))
					Expect(FPDF_SaveAsCopy.FileBytes).To(Not(BeNil()))
					Expect(FPDF_SaveAsCopy.FileBytes).To(PointTo(HaveLen(11375)))
				})
			})

			Context("and saved to a file path", func() {
				It("it returns the correct bytes", func() {
					tempFile, err := ioutil.TempFile("", "")
					Expect(err).To(BeNil())
					defer tempFile.Close()
					defer os.Remove(tempFile.Name())

					tempFileName := tempFile.Name()
					FPDF_SaveAsCopy, err := PdfiumInstance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
						Document: doc,
						FilePath: &tempFileName,
					})

					Expect(err).To(BeNil())
					fileStat, err := tempFile.Stat()
					Expect(err).To(BeNil())
					Expect(FPDF_SaveAsCopy).To(Not(BeNil()))
					Expect(FPDF_SaveAsCopy.FileBytes).To(BeNil())
					Expect(fileStat.Size()).To(Equal(int64(11375)))
				})
			})

			Context("and saved to a file path that does not work", func() {
				It("it returns an error", func() {
					fakeFilePath := "/path/that/will/never/work"
					FPDF_SaveAsCopy, err := PdfiumInstance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
						Document: doc,
						FilePath: &fakeFilePath,
					})

					Expect(err).To(Not(BeNil()))
					Expect(FPDF_SaveAsCopy).To(BeNil())
				})
			})

			Context("and saved to a io.Writer", func() {
				BeforeEach(func() {
					if TestType == "multi" {
						Skip("Multi-threaded usage does not support io.Writer")
					}
				})

				It("it returns the correct bytes", func() {
					buffer := bytes.Buffer{}
					FPDF_SaveAsCopy, err := PdfiumInstance.FPDF_SaveAsCopy(&requests.FPDF_SaveAsCopy{
						Document:   doc,
						FileWriter: &buffer,
					})
					Expect(err).To(BeNil())
					Expect(FPDF_SaveAsCopy).To(Not(BeNil()))
					Expect(FPDF_SaveAsCopy.FileBytes).To(BeNil())
					Expect(buffer.Len()).To(Equal(11375))
				})
			})

			Context("and saved with another PDF version", func() {
				It("it returns the correct byte array", func() {
					FPDF_SaveWithVersion, err := PdfiumInstance.FPDF_SaveWithVersion(&requests.FPDF_SaveWithVersion{
						Document:    doc,
						FileVersion: 13,
					})
					Expect(err).To(BeNil())
					Expect(FPDF_SaveWithVersion).To(Not(BeNil()))
					Expect(FPDF_SaveWithVersion.FileBytes).To(Not(BeNil()))
					Expect(FPDF_SaveWithVersion.FileBytes).To(PointTo(HaveLen(11375)))
				})
			})
		})
	})
})
