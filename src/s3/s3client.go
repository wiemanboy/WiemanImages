package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Client(region *string, endpoint *string, accessKeyId string, secretAccessKey string) *s3.S3 {
	s3Session, _ := session.NewSession(&aws.Config{
		Region:           region,
		Endpoint:         endpoint,
		Credentials:      credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})

	return s3.New(s3Session)
}
