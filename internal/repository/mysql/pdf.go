package mysql

import (
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
)

type PDFRepository struct {
	Conn *sql.DB
}

func NewPDFRepository(conn *sql.DB) *PDFRepository {
	return &PDFRepository{conn}
}
func (r *PDFRepository) Open(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Error("Failed to open PDF file: ", err)
		return nil, err
	}
	return file, nil
}

func (r *PDFRepository) Close(file *os.File) error {
	if err := file.Close(); err != nil {
		logrus.Error("Failed to close PDF file: ", err)
		return err
	}
	return nil
}

func (r *PDFRepository) CreateTempFile(prefix string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", prefix)
	if err != nil {
		logrus.Error("Failed to create temp file: ", err)
		return nil, err
	}
	return tempFile, nil
}
