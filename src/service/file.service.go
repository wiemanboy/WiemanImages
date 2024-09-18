package service

import (
	"WiemanImages/src/data"
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
)

type FileService struct {
	fileRepository data.FileRepository
}

func NewFileService(fileRepository data.FileRepository) FileService {
	return FileService{
		fileRepository: fileRepository,
	}
}

func (service *FileService) GetFile(objectKey string, imageSize string) ([]string, []byte, error) {
	fileList, _ := service.fileRepository.ListFiles(objectKey)
	if len(fileList) > 1 {
		return fileList, nil, nil
	}
	imageBytes, err := service.fileRepository.GetFile(objectKey)

	if err != nil {
		return nil, nil, err
	}

	inputImage, imageFormat, err := readBytes(imageBytes)
	if err != nil {
		return nil, nil, err
	}

	encodedImage, err := encodeImage(resizeImage(inputImage, imageSize), imageFormat)
	return fileList, encodedImage, err
}

func (service *FileService) CreateFile(objectKey string, fileContent []byte) error {
	return service.fileRepository.SaveFile(objectKey, fileContent)
}

func readBytes(imageBytes []byte) (image.Image, string, error) {
	imageBuffer := bytes.NewReader(imageBytes)
	return image.Decode(imageBuffer)
}

func resizeImage(inputImage image.Image, imageSize string) image.Image {
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
		return resize.Resize(targetWidth, 0, inputImage, resize.Lanczos3)
	}
	return inputImage
}

func encodeImage(inputImage image.Image, format string) ([]byte, error) {
	var outputBuffer bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err := jpeg.Encode(&outputBuffer, inputImage, nil)
		return outputBuffer.Bytes(), err
	case "png":
		err := png.Encode(&outputBuffer, inputImage)
		return outputBuffer.Bytes(), err
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
}
