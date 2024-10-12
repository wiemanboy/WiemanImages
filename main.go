package main

import (
	"WiemanImages/config"
	_ "WiemanImages/docs"
	"WiemanImages/src/client"
	"WiemanImages/src/data"
	"WiemanImages/src/presentation"
	"WiemanImages/src/presentation/controller/auth"
	"WiemanImages/src/presentation/controller/files"
	"WiemanImages/src/presentation/middleware"
	"WiemanImages/src/service"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:generate swag init

// @title Wieman Images API
// @description This is the Wieman Images service API.
func main() {
	exception := godotenv.Load()
	if exception != nil {
		print("Failed loading from .env file")
	}

	appConfig := config.LoadConfig()

	auth0Client := client.NewAuth0Client("https://" + appConfig.Auth0Domain)
	s3Client := client.NewS3Client(&appConfig.Region, &appConfig.S3Endpoint, appConfig.AccessKeyID, appConfig.SecretAccessKey)

	authService := service.NewAuthService(*auth0Client, appConfig.Auth0Domain, appConfig.Auth0ClientId, appConfig.Auth0ClientSecret, appConfig.Auth0CallbackUrl)
	authController := auth.NewController(authService)
	authMiddleware := middleware.NewAuthorizedMiddleware(&authService)

	fileRepository := data.NewS3Repository(s3Client, appConfig.BucketName)
	fileService := service.NewFileService(fileRepository)
	fileController := files.NewFileController(fileService)

	app := gin.Default()

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("auth-session", store))

	presentation.ApplyRoutes(app, fileController, authController, authMiddleware)
	_ = app.Run(":" + appConfig.Port)
}
