// main.go
package main

import (
	"WiemanCDN/config"
	"WiemanCDN/src/data"
	"WiemanCDN/src/presentation"
	"WiemanCDN/src/presentation/controller"
	"WiemanCDN/src/s3"
	"WiemanCDN/src/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	appConfig := config.LoadConfig()

	s3Client := s3.NewS3Client(&appConfig.Region, &appConfig.S3Endpoint, appConfig.AccessKeyID, appConfig.SecretAccessKey)
	fileRepository := data.NewS3Repository(s3Client, appConfig.BucketName)
	fileService := service.NewFileService(fileRepository)
	fileController := controller.NewFileController(fileService)

	app := gin.Default()
	presentation.ApplyRoutes(app, fileController)
	app.Run(":" + appConfig.Port)
}
