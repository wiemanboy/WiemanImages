package config

import (
	"os"
	"strconv"
)

type Config struct {
	S3Endpoint        string
	AccessKeyID       string
	SecretAccessKey   string
	BucketName        string
	Region            string
	Port              string
	JWTSecret         string
	JWTExpirationTime int
	AdminUsername     string
	AdminPassword     string
}

func LoadConfig() *Config {
	return &Config{
		S3Endpoint:        os.Getenv("S3_ENDPOINT"),
		AccessKeyID:       os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey:   os.Getenv("SECRET_ACCESS_KEY"),
		BucketName:        os.Getenv("BUCKET_NAME"),
		Region:            os.Getenv("REGION"),
		Port:              os.Getenv("PORT"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpirationTime: toInt(os.Getenv("JWT_EXPIRATION_TIME")),
		AdminUsername:     os.Getenv("ADMIN_USERNAME"),
		AdminPassword:     os.Getenv("ADMIN_PASSWORD"),
	}
}

func toInt(string string) int {
	integer, _ := strconv.Atoi(string)
	return integer
}
