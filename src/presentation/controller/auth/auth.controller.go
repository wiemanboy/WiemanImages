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

func (controller *AuthController) Callback(context *gin.Context) {
	session := sessions.Default(context)
	if context.Query("state") != session.Get("state") {
		context.String(http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := controller.authService.Exchange(context.Request.Context(), context.Query("code"))
	if err != nil {
		context.String(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
		return
	}

	idToken, err := controller.authService.VerifyIDToken(context.Request.Context(), token)
	if err != nil {
		context.String(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to docs page.
	context.Redirect(http.StatusTemporaryRedirect, "/services/files/docs")
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
