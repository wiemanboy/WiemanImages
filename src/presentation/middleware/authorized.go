package middleware

import (
	"WiemanImages/src/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthorizedMiddleware struct {
	authService *service.AuthService
}

func NewAuthorizedMiddleware(authService *service.AuthService) *AuthorizedMiddleware {
	return &AuthorizedMiddleware{authService: authService}
}

func (middleware *AuthorizedMiddleware) Check(context *gin.Context) {
	token := context.GetHeader("Authorization")

	if middleware.authService.CheckJwt(strings.TrimPrefix(token, "Bearer ")) {
		context.Next()
		return
	}

	if sessions.Default(context).Get("profile") != nil {
		context.Next()
		return
	}

	context.Redirect(http.StatusSeeOther, "/services/files/auth/login")
	context.AbortWithStatus(http.StatusUnauthorized)
}
