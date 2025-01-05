package pdf

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFService struct {
	pdfRepo PDFRepository
}

type PDFRepository interface {
	Open(filePath string) (*os.File, error)
	Close(file *os.File) error
	CreateTempFile(prefix string) (*os.File, error)
}

func NewService(p PDFRepository) *PDFService {
	return &PDFService{
		pdfRepo: p,
	}
}

func (s *PDFService) CompressPDF(inputFile string) (string, error) {
	outputFile, err := os.CreateTemp("", "compressed-*.pdf")
	if err != nil {
		return "", err
	}
	outputFileName := outputFile.Name()
	outputFile.Close()
	err = api.OptimizeFile(inputFile, outputFileName, nil)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
