package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup, controller *AuthController) {
	auth := r.Group("/auth")
	{
		auth.GET("/login", controller.Login)
		auth.GET("/refresh", controller.Refresh)
	}
}
