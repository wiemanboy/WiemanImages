package service

import (
	"WiemanImages/mocks/WiemanImages/src/data"
	"testing"
)

var fileRepo data.MockFileRepository
var fileService FileService

func InitFileServiceTest(m *testing.M) {
	fileRepo := new(data.MockFileRepository)
	fileService = NewFileService(fileRepo)
	m.Run()
}
