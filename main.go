package main

import (
	"WiemanImages/config"
	"WiemanImages/src/data"
	"WiemanImages/src/presentation"
	"WiemanImages/src/presentation/controller/auth"
	"WiemanImages/src/presentation/controller/files"
	"WiemanImages/src/presentation/middleware"
	"WiemanImages/src/s3"
	"WiemanImages/src/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	exception := godotenv.Load()
	if exception != nil {
		print("Failed loading from .env file")
	}

	appConfig := config.LoadConfig()

	s3Client := s3.NewS3Client(&appConfig.Region, &appConfig.S3Endpoint, appConfig.AccessKeyID, appConfig.SecretAccessKey)

	authService := service.NewAuthService()
	authController := auth.NewController(authService)
	authMiddleware := middleware.NewAuthorizedMiddleware(&authService)

	fileRepository := data.NewS3Repository(s3Client, appConfig.BucketName)
	fileService := service.NewFileService(fileRepository)
	fileController := files.NewFileController(fileService)

	app := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("auth-session", store))

	presentation.ApplyRoutes(app, fileController, authController, authMiddleware)
	app.Run(":" + appConfig.Port)
}
