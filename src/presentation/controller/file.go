package controller

import (
	"github.com/gin-gonic/gin"
)

// ApplyFileRoutes applies router to the gin Engine
func ApplyFileRoutes(router *gin.RouterGroup, controller *FileController) {
	files := router.Group("/files")
	{
		files.GET("/*objectKey", controller.Read)
		files.POST("/", controller.Create)
	}
}
