package auth

import (
	"WiemanImages/src/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService      service.AuthService
	cookieExpiration int
}

func NewController(authService service.AuthService, cookieExpiration int) *AuthController {
	return &AuthController{
		authService:      authService,
		cookieExpiration: cookieExpiration,
	}
}

func (controller *AuthController) Login(context *gin.Context) {
	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(400)
		return
	}

	token, err := controller.authService.Login(body.Username, body.Password)

	if err != nil {
		context.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	context.SetCookie("token", token, controller.cookieExpiration, "/", "", true, true)

	context.JSON(200, gin.H{
		"token": token,
	})
}

func (controller *AuthController) Refresh(context *gin.Context) {

	cookieToken, err := context.Request.Cookie("token")

	if err != nil {
		context.JSON(400, gin.H{
			"error": "no token provided, please login first",
		})
		return
	}

	token, err := controller.authService.Refresh(cookieToken.Value)

	if err != nil {
		context.JSON(200, gin.H{
			"token": nil,
		})
		return
	}

	context.SetCookie("token", token, controller.cookieExpiration, "/", "", true, true)
	context.JSON(200, gin.H{
		"token": token,
	})
}
