package middleware

import (
	"WiemanImages/src/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizedMiddleware struct {
	authService *service.AuthService
}

func NewAuthorizedMiddleware(authService *service.AuthService) *AuthorizedMiddleware {
	return &AuthorizedMiddleware{authService: authService}
}

func (middleware *AuthorizedMiddleware) Check(context *gin.Context) {
	if sessions.Default(context).Get("profile") == nil {
		context.Redirect(http.StatusSeeOther, "/services/files/auth/login")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Next()
}
