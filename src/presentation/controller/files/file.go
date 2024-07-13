package files

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(router *gin.RouterGroup, controller *FileController) {
	files := router.Group("/files")
	{
		files.GET("/*objectKey", controller.Read)
		files.POST("/", controller.Create)
	}
}
