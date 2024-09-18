package files

import (
	"WiemanImages/src/presentation/middleware"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(router *gin.RouterGroup, controller *FileController, authMiddleware *middleware.AuthorizedMiddleware) {
	files := router.Group("/files")
	{
		files.GET("/*objectKey", controller.Read)
		files.POST("/", authMiddleware.Check, controller.Create)
	}
}
