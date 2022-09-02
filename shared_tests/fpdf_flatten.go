package shared_tests

import (
	"github.com/michalderdak/go-pdfium/requests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fpdf_flatten", func() {
	BeforeEach(func() {
		Locker.Lock()
	})

	AfterEach(func() {
		Locker.Unlock()
	})

	Context("no document", func() {
		When("is opened", func() {
			It("returns an error when flattening a pdf page", func() {
				pageCount, err := PdfiumInstance.FPDFPage_Flatten(&requests.FPDFPage_Flatten{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Index: 0,
						},
					},
				})
				Expect(err).To(MatchError("document not given"))
				Expect(pageCount).To(BeNil())
			})
		})
	})
})
