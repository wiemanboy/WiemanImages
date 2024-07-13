package files

import (
	"WiemanCDN/src/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FileController struct {
	fileService service.FileService
}

func NewFileController(fileService service.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

func (controller *FileController) Read(context *gin.Context) {
	objectKey := context.Param("objectKey")
	imageSize := context.Query("size")
	cookieToken, _ := context.Request.Cookie("token")

	token := ""
	if cookieToken != nil {
		token = cookieToken.Value
	}

	fileContent, err := controller.fileService.GetFile(objectKey, imageSize, token)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	context.Data(http.StatusOK, "image/webp", fileContent)
}

func (controller *FileController) Create(context *gin.Context) {
	formFile, fileHeader, err := context.Request.FormFile("image")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve image file"})
		return
	}

	key := context.PostForm("key")
	locked := context.PostForm("locked")
	fileBytes, _ := io.ReadAll(formFile)

	err = controller.fileService.CreateFile(key, fileBytes, locked)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save image file"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"file":   fileHeader.Filename,
		"key":    key,
		"locked": locked,
	})
}