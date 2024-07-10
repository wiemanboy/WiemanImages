package service

import "WiemanCDN/src/data"

type FileService struct {
	fileRepository data.FileRepository
}

func NewFileService(fileRepository data.FileRepository) *FileService {
	return &FileService{
		fileRepository: fileRepository,
	}
}

func (service *FileService) GetFile(objectKey string) ([]byte, error) {
	return service.fileRepository.GetFile(objectKey)
}
