package presentation

import (
	"WiemanImages/src/presentation/controller/auth"
	"WiemanImages/src/presentation/controller/files"
	"WiemanImages/src/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.Engine, fileController *files.FileController, authController *auth.AuthController, authMiddleware *middleware.AuthorizedMiddleware) {
	api := router.Group("/api")
	{
		api.GET("/ping", ping)
		files.ApplyRoutes(api, fileController, authMiddleware)
		auth.ApplyRoutes(api, authController)
	}
}
