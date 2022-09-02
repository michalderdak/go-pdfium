package worker

import (
	"github.com/michalderdak/go-pdfium"
	"github.com/michalderdak/go-pdfium/internal/implementation"
)

func StartWorker(config *pdfium.LibraryConfig) {
	implementation.StartPlugin(config)
}
