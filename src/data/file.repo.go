package data

import (
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type FileRepository interface {
	GetFile(objectKey string) ([]byte, error)
}

type S3Repository struct {
	s3Client   *s3.S3
	bucketName string
}

func NewS3Repository(s3Client *s3.S3, bucketName string) FileRepository {
	return &S3Repository{s3Client: s3Client, bucketName: bucketName}
}

func (repo *S3Repository) GetFile(objectKey string) ([]byte, error) {
	objectOutput, exception := repo.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(repo.bucketName),
		Key:    aws.String(objectKey),
	})
	if exception != nil {
		return nil, exception
	}
	defer objectOutput.Body.Close()

	bodyBytes, exception := ioutil.ReadAll(objectOutput.Body)
	if exception != nil {
		return nil, exception
	}

	return bodyBytes, nil
}

func SaveToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
