package service

import (
	"WiemanImages/mocks/WiemanImages/src/data"
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileService_GetFile_returnsList(t *testing.T) {
	var fileRepo data.MockFileRepository
	fileRepo.On("ListFiles", "test").Return([]string{"test", "test"}, nil)
	var fileService = NewFileService(&fileRepo)

	fileList, file, err := fileService.GetFile("test", "sm")

	assert.Equal(t, []string{"test", "test"}, fileList)
	assert.Nil(t, file)
	assert.Nil(t, err)
}

func TestFileService_GetFile_returnsImage(t *testing.T) {
	testCases := []struct {
		imageSize   string
		minByteSize int
		maxByteSize int
	}{
		{"sm", 30000, 60000},
		{"md", 50000, 200000},
		{"lg", 200000, 600000},
		{"xl", 300000, 10000000},
	}

	for _, testCase := range testCases {
		t.Run(testCase.imageSize, func(t *testing.T) {
			var fileRepo data.MockFileRepository
			fileRepo.On("ListFiles", "test").Return([]string{"test"}, nil)
			fileRepo.On("GetFile", "test").Return(getImageBytes(t), nil)
			var fileService = NewFileService(&fileRepo)

			fileList, file, err := fileService.GetFile("test", testCase.imageSize)

			assert.Equal(t, []string{"test"}, fileList)
			assert.GreaterOrEqual(t, binary.Size(file), testCase.minByteSize)
			assert.LessOrEqual(t, binary.Size(file), testCase.maxByteSize)
			assert.Nil(t, err)
		})
	}
}

func getImageBytes(t *testing.T) []byte {
	imageBytes, err := os.ReadFile("../../dev/test.png")
	if err != nil {
		t.Fatalf("Failed to read image file: %v", err)
	}
	return imageBytes
}
