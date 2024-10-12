package files

import (
	"WiemanImages/src/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FileController struct {
	fileService service.FileService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type FileListResponse struct {
	Files []string `json:"files"`
}

type FileCreateResponse struct {
	FileName string `json:"fileName"`
	FileKey  string `json:"fileKey"`
}

func NewFileController(fileService service.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

// Read
// @Summary Fetch file
// @Description Fetches a file from the storage if it is an image it will be scaled to the requested size
// @Tags Files
// @Param key path string true "File key"
// @Success 200 {string} string "file contents"
// @Success 200 {object} FileListResponse "List of files"
// @Failure 404 {object} ErrorResponse "File not found"
// @Router /api/files/{key} [get]
func (controller *FileController) Read(context *gin.Context) {
	objectKey := context.Param("objectKey")
	imageSize := context.Query("size")

	fileList, fileContent, err := controller.fileService.GetFile(objectKey, imageSize)

	if err != nil {
		context.JSON(http.StatusNotFound, ErrorResponse{Error: "File not found"})
		return
	}
	if len(fileList) > 1 {
		context.JSON(http.StatusOK, FileListResponse{Files: fileList})
		return
	}

	context.Data(http.StatusOK, "image/webp", fileContent)
}

// Create
// @Summary Create file
// @Description Creates a file in the storage
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file to upload"
// @Param key formData string true "Key for the file"
// @Success 201 {object} FileCreateResponse "File successfully uploaded"
// @Failure 400 {object} ErrorResponse "Failed to retrieve image file from form"
// @Failure 400 {object} ErrorResponse "Failed to save image file please try again"
// @Router /api/files/ [post]
func (controller *FileController) Create(context *gin.Context) {
	formFile, fileHeader, err := context.Request.FormFile("image")
	if err != nil {

		context.JSON(http.StatusBadRequest, ErrorResponse{Error: "Failed to retrieve file from form"})
		return
	}

	key := context.PostForm("key")
	fileBytes, _ := io.ReadAll(formFile)

	err = controller.fileService.CreateFile(key, fileBytes)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: "Failed to save file please try again"})
		return
	}

	context.JSON(http.StatusCreated, FileCreateResponse{
		FileName: fileHeader.Filename,
		FileKey:  key,
	})
}
