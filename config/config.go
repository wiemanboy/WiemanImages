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
	Auth0Domain       string
	Auth0ClientId     string
	Auth0ClientSecret string
	Auth0CallbackUrl  string
}

func LoadConfig() *Config {
	return &Config{
		S3Endpoint:        os.Getenv("S3_ENDPOINT"),
		AccessKeyID:       os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey:   os.Getenv("SECRET_ACCESS_KEY"),
		BucketName:        os.Getenv("BUCKET_NAME"),
		Region:            os.Getenv("REGION"),
		Port:              os.Getenv("PORT"),
		Auth0Domain:       os.Getenv("AUTH0_DOMAIN"),
		Auth0ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Auth0CallbackUrl:  os.Getenv("AUTH0_CALLBACK_URL"),
	}
}

func toInt(string string) int {
	integer, _ := strconv.Atoi(string)
	return integer
}
