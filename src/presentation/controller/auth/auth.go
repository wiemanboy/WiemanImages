package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup, controller *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/refresh", controller.Refresh)
	}
}
