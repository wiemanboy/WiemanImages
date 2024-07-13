package main

import (
	"WiemanCDN/config"
	"WiemanCDN/src/data"
	"WiemanCDN/src/presentation"
	"WiemanCDN/src/presentation/controller/auth"
	"WiemanCDN/src/presentation/controller/files"
	"WiemanCDN/src/s3"
	"WiemanCDN/src/service"
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

	fileRepository := data.NewS3Repository(s3Client, appConfig.BucketName)
	fileService := service.NewFileService(fileRepository)
	fileController := files.NewFileController(fileService)

	authService := service.NewAuthService(appConfig.JWTSecret, appConfig.JWTExpirationTime, appConfig.AdminUsername, appConfig.AdminPassword)
	authController := auth.NewController(authService, appConfig.JWTExpirationTime)

	app := gin.Default()
	presentation.ApplyRoutes(app, fileController, authController)
	app.Run(":" + appConfig.Port)
}
