package implementation_test

import (
	"os"
	"time"

	"github.com/michalderdak/go-pdfium/internal/implementation"
	"github.com/michalderdak/go-pdfium/shared_tests"
	"github.com/michalderdak/go-pdfium/single_threaded"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	implementation.InitLibrary(nil)

	// Set ENV to ensure resulting values.
	err := os.Setenv("TZ", "UTC")
	Expect(err).To(BeNil())

	pool := single_threaded.Init(single_threaded.Config{})
	shared_tests.PdfiumPool = pool

	instance, err := pool.GetInstance(time.Second * 30)
	Expect(err).To(BeNil())
	shared_tests.PdfiumInstance = instance
	shared_tests.TestDataPath = "../../shared_tests"
	shared_tests.TestType = "internal"
})

var _ = AfterSuite(func() {
	err := shared_tests.PdfiumInstance.Close()
	Expect(err).To(BeNil())

	err = shared_tests.PdfiumPool.Close()
	Expect(err).To(BeNil())
})

var _ = Describe("Implementation", func() {
	When("pinged", func() {
		It("pongs", func() {
			pdfium := implementation.Pdfium.GetInstance()
			resp, err := pdfium.Ping()
			Expect(err).To(BeNil())
			Expect(resp).To(Equal("Pong"))
			pdfium.Close()
		})
	})

	shared_tests.Import()
})
