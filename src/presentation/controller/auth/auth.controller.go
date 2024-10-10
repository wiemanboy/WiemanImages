package auth

import (
	"WiemanImages/src/service"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) Login(context *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Save the state inside the session.
	session := sessions.Default(context)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.Redirect(http.StatusTemporaryRedirect, controller.authService.AuthCodeURL(state))
}

func (controller *AuthController) Refresh(context *gin.Context) {
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
