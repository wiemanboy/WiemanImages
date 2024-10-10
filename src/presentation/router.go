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

	services := router.Group("/services/files")
	{
		services.GET("/ping", ping)
		services.GET("/docs", ping)
		auth.ApplyRoutes(services, authController)
	}

	api := router.Group("/api")
	{
		files.ApplyRoutes(api, fileController, authMiddleware)
		api.GET("/ping", authMiddleware.Check, ping)
	}
}
