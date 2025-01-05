package rest

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type PDFService interface {
	CompressPDF(inputFile string) (string, error)
}

type PDFHandler struct {
	Service PDFService
}

func NewPDFHandler(e *echo.Echo, svc PDFService) {
	handler := &PDFHandler{
		Service: svc,
	}
	e.POST("/pdf/compress", handler.CompressPDF)

}

func (a *PDFHandler) CompressPDF(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("", "input-*.pdf")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create temp file"})
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}

	inputFile := tempFile.Name()

	outputFile, err := a.Service.CompressPDF(inputFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to compress PDF"})
	}
	defer os.Remove(outputFile)

	compressedFile, err := os.Open(outputFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open compressed PDF"})
	}
	defer compressedFile.Close()

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+file.Filename+"_compressed.pdf")
	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

	if _, err := io.Copy(c.Response().Writer, compressedFile); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send compressed PDF"})
	}

	return nil
}
