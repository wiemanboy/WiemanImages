package service

import (
	"WiemanCDN/src/data"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

type FileService struct {
	fileRepository data.FileRepository
}

func NewFileService(fileRepository data.FileRepository) FileService {
	return FileService{
		fileRepository: fileRepository,
	}
}

func (service *FileService) GetFile(objectKey string, imageSize string) ([]byte, error) {
	imageBytes, err := service.fileRepository.GetFile(objectKey)
	if err != nil {
		return nil, err
	}

	imageBuffer := bytes.NewReader(imageBytes)
	inputImage, imageFormat, err := image.Decode(imageBuffer)
	if err != nil {
		return nil, err
	}

	originalWidth := inputImage.Bounds().Dx()

	var targetWidth uint
	switch imageSize {
	case "sm":
		targetWidth = 320
	case "md":
		targetWidth = 740
	case "lg":
		targetWidth = 1200
	default:
		targetWidth = uint(originalWidth)
	}

	if uint(originalWidth) > targetWidth {
		inputImage = resize.Resize(targetWidth, 0, inputImage, resize.Lanczos3)
	}

	var outputBuffer bytes.Buffer
	switch imageFormat {
	case "jpeg", "jpg":
		err = jpeg.Encode(&outputBuffer, inputImage, nil)
	case "png":
		err = png.Encode(&outputBuffer, inputImage)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", imageFormat)
	}

	if err != nil {
		return nil, err
	}

	return outputBuffer.Bytes(), nil
}
