package mocks

import (
	"os"

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
