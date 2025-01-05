package pdf

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type MockPDFRepository struct {
	mock.Mock
}

func (m *MockPDFRepository) Open(filePath string) (*os.File, error) {
	args := m.Called(filePath)
	return args.Get(0).(*os.File), args.Error(1)
}

func (m *MockPDFRepository) Close(file *os.File) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockPDFRepository) CreateTempFile(prefix string) (*os.File, error) {
	args := m.Called(prefix)
	return args.Get(0).(*os.File), args.Error(1)
}

func TestCompressPDF(t *testing.T) {
	mockRepo := new(MockPDFRepository)
	service := NewService(mockRepo)

	// Mock the os.CreateTemp function
	tempFile, _ := os.CreateTemp("", "compressed-*.pdf")
	defer os.Remove(tempFile.Name())

	mockRepo.On("CreateTempFile", "compressed-").Return(tempFile, nil)

	// api.OptimizeFile = func(inputFile, outputFile string, nilfunc func()) error {
	// 	return nil
	// }

	outputFileName, err := service.CompressPDF("input.pdf")
	assert.NoError(t, err)
	assert.Equal(t, tempFile.Name(), outputFileName)

	mockRepo.AssertExpectations(t)
}

func TestCompressPDF_Error(t *testing.T) {
	mockRepo := new(MockPDFRepository)
	service := NewService(mockRepo)

	// Mock the os.CreateTemp function to return an error
	mockRepo.On("CreateTempFile", "compressed-").Return(nil, errors.New("failed to create temp file"))

	_, err := service.CompressPDF("input.pdf")
	assert.Error(t, err)
	assert.Equal(t, "failed to create temp file", err.Error())

	mockRepo.AssertExpectations(t)
}
