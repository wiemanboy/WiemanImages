package presentation

import (
	"WiemanCDN/src/presentation/controller/auth"
	"WiemanCDN/src/presentation/controller/files"
	"github.com/gin-gonic/gin"
)

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.Engine, fileController *files.FileController, authController *auth.AuthController) {
	api := router.Group("/api")
	{
		api.GET("/ping", ping)
		files.ApplyRoutes(api, fileController)
		auth.ApplyRoutes(api, authController)
	}
}
