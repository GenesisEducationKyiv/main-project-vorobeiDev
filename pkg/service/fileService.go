package service

import (
	"errors"
	"os"
	"strings"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

var ErrEmailExists = errors.New("email already exists")

func (service *FileService) WriteToFile(email string) error {
	fileName := os.Getenv("DB_FILE_NAME")
	if !service.isFileExists(fileName) {
		_, err := os.Create(fileName)
		if err != nil {
			return err
		}
	}

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileString := string(fileData)
	if strings.Contains(fileString, email) {
		return ErrEmailExists
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(email + "\n")

	return err
}

func (service *FileService) isFileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}

func (service *FileService) ReadFromFile() ([]string, error) {
	fileName := os.Getenv("DB_FILE_NAME")
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	fileString := string(fileData)
	emails := strings.Split(fileString, "\n")

	return emails, nil
}
