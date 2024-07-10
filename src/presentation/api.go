package presentation

import (
	"WiemanCDN/src/presentation/controller"
	"github.com/gin-gonic/gin"
)

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.Engine, fileController *controller.FileController) {
	api := router.Group("/api")
	{
		api.GET("/ping", ping)
		controller.ApplyFileRoutes(api, fileController)
	}
}
