package shared_tests

import (
	"github.com/klippa-app/go-pdfium/errors"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

func RunDocumentTests(pdfiumContainer pdfium.Pdfium, testsPath string, prefix string) {
	Describe("NewDocumentFromBytes", func() {
		Context("a normal PDF file with 1 page", func() {
			var doc pdfium.Document

			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/test.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
			})

			When("is opened", func() {
				It("returns the correct file version", func() {
					fileVersion, err := doc.GetFileVersion(&requests.GetFileVersion{})
					Expect(err).To(BeNil())
					Expect(fileVersion).To(Equal(&responses.GetFileVersion{
						FileVersion: 15,
					}))
				})

				It("returns the correct document permissions", func() {
					docPermissions, err := doc.GetDocPermissions(&requests.GetDocPermissions{})
					Expect(err).To(BeNil())
					Expect(docPermissions).To(Equal(&responses.GetDocPermissions{
						DocPermissions: 0xffffffff, // 0xffffffff (4294967295) = not protected
					}))
				})

				It("returns the correct security handler revision", func() {
					securityHandlerRevision, err := doc.GetSecurityHandlerRevision(&requests.GetSecurityHandlerRevision{})
					Expect(err).To(BeNil())
					Expect(securityHandlerRevision).To(Equal(&responses.GetSecurityHandlerRevision{
						SecurityHandlerRevision: -1, // -1 = no security handler.
					}))
				})

				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 1,
					}))
				})

				It("returns the correct page mode", func() {
					pageMode, err := doc.GetPageMode(&requests.GetPageMode{})
					Expect(err).To(BeNil())
					Expect(pageMode).To(Equal(&responses.GetPageMode{
						PageMode: responses.PageModeUseNone,
					}))
				})

				It("returns the correct metadata", func() {
					metadata, err := doc.GetMetadata(&requests.GetMetadata{
						Tag: "Producer",
					})
					Expect(err).To(BeNil())
					Expect(metadata).To(Equal(&responses.GetMetadata{
						Tag:   "Producer",
						Value: "cairo 1.16.0 (https://cairographics.org)",
					}))
				})
			})
		})

		Context("a normal PDF file with multiple pages", func() {
			var doc pdfium.Document

			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/test_multipage.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
			})

			When("is opened", func() {
				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 2,
					}))
				})
			})
		})

		Context("a password protected PDF file", func() {
			pdfData, _ := ioutil.ReadFile(testsPath + "/testdata/password_test123.pdf")
			When("is opened with no password", func() {
				It("returns the password error", func() {
					doc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the wrong password", func() {
				It("returns the password error", func() {
					wrongPassword := "test"
					doc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData, pdfium.OpenDocumentWithPasswordOption(wrongPassword))
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the correct password", func() {
				It("does not return an error", func() {
					pdfPassword := "test123"
					doc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData, pdfium.OpenDocumentWithPasswordOption(pdfPassword))
					Expect(err).To(BeNil())
					Expect(doc).To(Not(BeNil()))
					doc.Close()
				})
			})
		})
	})

	Describe("NewDocumentFromFilePath", func() {
		Context("a normal PDF file with 1 page", func() {
			var doc pdfium.Document

			BeforeEach(func() {
				newDoc, err := pdfiumContainer.NewDocumentFromFilePath(testsPath + "/testdata/test.pdf")
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
			})

			When("is opened", func() {
				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 1,
					}))
				})

				It("returns the correct metadata", func() {
					metadata, err := doc.GetMetadata(&requests.GetMetadata{
						Tag: "Producer",
					})
					Expect(err).To(BeNil())
					Expect(metadata).To(Equal(&responses.GetMetadata{
						Tag:   "Producer",
						Value: "cairo 1.16.0 (https://cairographics.org)",
					}))
				})
			})
		})

		Context("a normal PDF file with multiple pages", func() {
			var doc pdfium.Document

			BeforeEach(func() {
				newDoc, err := pdfiumContainer.NewDocumentFromFilePath(testsPath + "/testdata/test_multipage.pdf")
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
			})

			When("is opened", func() {
				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 2,
					}))
				})
			})
		})

		Context("a password protected PDF file", func() {
			filePath := testsPath + "/testdata/password_test123.pdf"
			When("is opened with no password", func() {
				It("returns the password error", func() {
					doc, err := pdfiumContainer.NewDocumentFromFilePath(filePath)
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the wrong password", func() {
				It("returns the password error", func() {
					wrongPassword := "test"
					doc, err := pdfiumContainer.NewDocumentFromFilePath(filePath, pdfium.OpenDocumentWithPasswordOption(wrongPassword))
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the correct password", func() {
				It("does not return an error", func() {
					pdfPassword := "test123"
					doc, err := pdfiumContainer.NewDocumentFromFilePath(filePath, pdfium.OpenDocumentWithPasswordOption(pdfPassword))
					Expect(err).To(BeNil())
					Expect(doc).To(Not(BeNil()))
					doc.Close()
				})
			})
		})

		Context("a non-existent file", func() {
			filePath := testsPath + "/testdata/i_dont_exist.pdf"
			When("is opened", func() {
				It("returns the file error", func() {
					doc, err := pdfiumContainer.NewDocumentFromFilePath(filePath)
					Expect(err).To(MatchError(errors.ErrFile.Error()))
					Expect(doc).To(BeNil())
				})
			})
		})
	})

	Describe("NewDocumentFromReader", func() {
		Context("a normal PDF file with 1 page", func() {
			var doc pdfium.Document

			BeforeEach(func() {
				file, err := os.Open(testsPath + "/testdata/test.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}
				fileStat, err := file.Stat()
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()))
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
			})

			When("is opened", func() {
				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 1,
					}))
				})

				It("returns the correct metadata", func() {
					metadata, err := doc.GetMetadata(&requests.GetMetadata{
						Tag: "Producer",
					})
					Expect(err).To(BeNil())
					Expect(metadata).To(Equal(&responses.GetMetadata{
						Tag:   "Producer",
						Value: "cairo 1.16.0 (https://cairographics.org)",
					}))
				})
			})
		})

		Context("a normal PDF file with multiple pages", func() {
			var doc pdfium.Document
			var file *os.File

			BeforeEach(func() {
				pdfFile, err := os.Open(testsPath + "/testdata/test_multipage.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}
				file = pdfFile
				fileStat, err := file.Stat()
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()))
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
				file.Close()
			})

			When("is opened", func() {
				It("returns the correct file version", func() {
					pageCount, err := doc.GetFileVersion(&requests.GetFileVersion{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetFileVersion{
						FileVersion: 15,
					}))
				})

				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 2,
					}))
				})
			})
		})

		Context("a normal PDF file with alpha channel", func() {
			var doc pdfium.Document
			var file *os.File

			BeforeEach(func() {
				pdfFile, err := os.Open(testsPath + "/testdata/alpha_channel.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}
				file = pdfFile
				fileStat, err := file.Stat()
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()))
				if err != nil {
					return
				}

				doc = newDoc
			})

			AfterEach(func() {
				doc.Close()
				file.Close()
			})

			When("is opened", func() {
				It("returns the correct file version", func() {
					pageCount, err := doc.GetFileVersion(&requests.GetFileVersion{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetFileVersion{
						FileVersion: 17,
					}))
				})

				It("returns the correct page count", func() {
					pageCount, err := doc.GetPageCount(&requests.GetPageCount{})
					Expect(err).To(BeNil())
					Expect(pageCount).To(Equal(&responses.GetPageCount{
						PageCount: 1,
					}))
				})
			})
		})

		Context("a password protected PDF file", func() {
			When("is opened with no password", func() {
				It("returns the password error", func() {
					file, err := os.Open(testsPath + "/testdata/password_test123.pdf")
					Expect(err).To(BeNil())
					if err != nil {
						return
					}
					fileStat, err := file.Stat()
					Expect(err).To(BeNil())
					if err != nil {
						return
					}

					doc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()))
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the wrong password", func() {
				It("returns the password error", func() {
					file, err := os.Open(testsPath + "/testdata/password_test123.pdf")
					Expect(err).To(BeNil())
					if err != nil {
						return
					}
					fileStat, err := file.Stat()
					Expect(err).To(BeNil())
					if err != nil {
						return
					}

					wrongPassword := "test"
					doc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()), pdfium.OpenDocumentWithPasswordOption(wrongPassword))
					Expect(err).To(MatchError(errors.ErrPassword.Error()))
					Expect(doc).To(BeNil())
				})
			})
			When("is opened with the correct password", func() {
				It("does not return an error", func() {
					file, err := os.Open(testsPath + "/testdata/password_test123.pdf")
					Expect(err).To(BeNil())
					if err != nil {
						return
					}
					fileStat, err := file.Stat()
					Expect(err).To(BeNil())
					if err != nil {
						return
					}

					pdfPassword := "test123"
					doc, err := pdfiumContainer.NewDocumentFromReader(file, int(fileStat.Size()), pdfium.OpenDocumentWithPasswordOption(pdfPassword))
					Expect(err).To(BeNil())
					Expect(doc).To(Not(BeNil()))

					fileVersion, err := doc.GetFileVersion(&requests.GetFileVersion{})
					Expect(err).To(BeNil())
					Expect(fileVersion).To(Equal(&responses.GetFileVersion{
						FileVersion: 15,
					}))

					docPermissions, err := doc.GetDocPermissions(&requests.GetDocPermissions{})
					Expect(err).To(BeNil())
					Expect(docPermissions).To(Equal(&responses.GetDocPermissions{
						DocPermissions: 0xFFFFFFFC, // 0xFFFFFFFC (4294967292) = owner password
					}))

					securityHandlerRevision, err := doc.GetSecurityHandlerRevision(&requests.GetSecurityHandlerRevision{})
					Expect(err).To(BeNil())
					Expect(securityHandlerRevision).To(Equal(&responses.GetSecurityHandlerRevision{
						SecurityHandlerRevision: 3,
					}))

					doc.Close()
				})
			})
		})
	})
}
