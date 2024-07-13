package data

import (
	"bytes"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type FileRepository interface {
	GetFile(objectKey string, lockAccess bool) ([]byte, error)
	SaveFile(filename string, data []byte, locked string) error
}

type S3Repository struct {
	s3Client   *s3.S3
	bucketName string
}

func NewS3Repository(s3Client *s3.S3, bucketName string) FileRepository {
	return &S3Repository{s3Client: s3Client, bucketName: bucketName}
}

func (repo *S3Repository) GetFile(objectKey string, lockAccess bool) ([]byte, error) {
	objectOutput, err := repo.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(repo.bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}

	locked := objectOutput.Metadata["Locked"]
	if !lockAccess && locked != nil && *locked == "true" {
		return nil, errors.New("file is locked")
	}

	defer objectOutput.Body.Close()

	bodyBytes, err := io.ReadAll(objectOutput.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (repo *S3Repository) SaveFile(filename string, data []byte, locked string) error {
	_, err := repo.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: &repo.bucketName,
		Key:    &filename,
		Body:   aws.ReadSeekCloser(bytes.NewReader(data)),
		Metadata: map[string]*string{
			"locked": aws.String(locked),
		},
		ContentType: aws.String("image/webp"),
	})
	return err
}
