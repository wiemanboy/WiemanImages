package config

import "os"

type Config struct {
	S3Endpoint      string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Region          string
	Port            string
}

func LoadConfig() *Config {
	return &Config{
		S3Endpoint:      os.Getenv("S3_ENDPOINT"),
		AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
		BucketName:      os.Getenv("BUCKET_NAME"),
		Region:          os.Getenv("REGION"),
		Port:            os.Getenv("PORT"),
	}
}
