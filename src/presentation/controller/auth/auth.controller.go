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

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login
// @Summary Redirects to Auth0 login page
// @Description Redirects to Auth0 login page for browser authentication
// @Tags Auth
// @Success 200
// @Failure 500 {object} ErrorResponse "Failed to generate random state"
// @Failure 500 {object} ErrorResponse "Failed to save session"
// @Router /services/files/auth/login [get]
func (controller *AuthController) Login(context *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Save the state inside the session.
	session := sessions.Default(context)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	context.Redirect(http.StatusTemporaryRedirect, controller.authService.AuthCodeURL(state))
}

// Callback
// @Summary Auth0 Callback
// @Description Callback for Auth0 browser authentication
// @Tags Auth
// @Success 200
// @Failure 401 {object} ErrorResponse "Failed to exchange an authorization code for a token"
// @Failure 500 {object} ErrorResponse "Failed to verify ID Token"
// @Router /services/files/auth/callback [get]
func (controller *AuthController) Callback(context *gin.Context) {
	session := sessions.Default(context)
	if context.Query("state") != session.Get("state") {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid state parameter"})
		return
	}

	// Exchange an authorization code for a token.
	token, err := controller.authService.Exchange(context.Request.Context(), context.Query("code"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Failed to exchange an authorization code for a token"})
		return
	}

	idToken, err := controller.authService.VerifyIDToken(context.Request.Context(), token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to verify ID Token."})
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
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
