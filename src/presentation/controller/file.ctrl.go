package controller

import (
	"WiemanCDN/src/data"
	"WiemanCDN/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {
	fileService data.FileRepository
}

func NewFileController(fileService *service.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

func (controller *FileController) Read(context *gin.Context) {
	objectKey := context.Param("objectKey")
	fileContent, err := controller.fileService.GetFile(objectKey)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	context.Data(http.StatusOK, "application/octet-stream", fileContent)
}
