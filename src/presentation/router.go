package presentation

import (
	"WiemanImages/src/presentation/controller/auth"
	"WiemanImages/src/presentation/controller/files"
	"WiemanImages/src/presentation/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type PingResponse struct {
	Message string `json:"message"`
}

// ping
// @Summary Ping
// @Description Health check endpoint
// @Tags Services
// @Success 200 {object} PingResponse "pong"
// @Router /services/files/ping [get]
func ping(context *gin.Context) {
	context.JSON(200, PingResponse{Message: "pong"})
}

func ApplyRoutes(router *gin.Engine, fileController *files.FileController, authController *auth.AuthController, authMiddleware *middleware.AuthorizedMiddleware) {

	services := router.Group("/services/files")
	{
		services.GET("/ping", ping)
		services.GET("/docs/*any", authMiddleware.Check, ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth.ApplyRoutes(services, authController)
	}

	api := router.Group("/api")
	{
		files.ApplyRoutes(api, fileController, authMiddleware)
	}
}
