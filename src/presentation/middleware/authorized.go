package middleware

import (
	"WiemanCDN/src/service"
	"github.com/gin-gonic/gin"
)

type AuthorizedMiddleware struct {
	authService *service.AuthService
}

func NewAuthorizedMiddleware(authService *service.AuthService) *AuthorizedMiddleware {
	return &AuthorizedMiddleware{authService: authService}
}

func (middleware *AuthorizedMiddleware) Check(context *gin.Context) {
	cookieToken, err := context.Request.Cookie("token")

	if err != nil || !middleware.authService.Check(cookieToken.Value) {
		context.JSON(401, gin.H{
			"error": "invalid token, please login again",
		})
		context.Abort()
		return
	}

	context.Next()
}
