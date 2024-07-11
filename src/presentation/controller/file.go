package controller

import (
	"github.com/gin-gonic/gin"
)

// ApplyFileRoutes applies router to the gin Engine
func ApplyFileRoutes(router *gin.RouterGroup, controller *FileController) {
	posts := router.Group("/files")
	{
		posts.GET("/*objectKey", controller.Read)
	}
}
