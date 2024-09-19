package data

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type FileRepository interface {
	ListFiles(objectKey string) ([]string, error)
	GetFile(objectKey string) ([]byte, error)
	SaveFile(filename string, data []byte) error
}

type S3Repository struct {
	s3Client   *s3.S3
	bucketName string
}

func NewS3Repository(s3Client *s3.S3, bucketName string) FileRepository {
	return &S3Repository{s3Client: s3Client, bucketName: bucketName}
}

func (repo *S3Repository) ListFiles(objectKey string) ([]string, error) {
	output, _ := repo.s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(repo.bucketName),
		Prefix: aws.String(objectKey[1:]),
	})

	var files []string
	for _, item := range output.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

func (repo *S3Repository) GetFile(objectKey string) ([]byte, error) {
	objectOutput, err := repo.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(repo.bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(objectOutput.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (repo *S3Repository) SaveFile(filename string, data []byte) error {
	_, err := repo.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      &repo.bucketName,
		Key:         &filename,
		Body:        aws.ReadSeekCloser(bytes.NewReader(data)),
		ContentType: aws.String("image/webp"),
	})
	return err
}
