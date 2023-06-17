package service

import (
	"errors"
	"os"
	"strings"
)

const FileName = "emails.txt"

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

var ErrEmailExists = errors.New("email already exists")

func (service *FileService) WriteToFile(email string) error {
	if !service.isFileExists() {
		_, err := os.Create(FileName)
		if err != nil {
			return err
		}
	}

	fileData, err := os.ReadFile(FileName)
	if err != nil {
		return err
	}

	fileString := string(fileData)
	if strings.Contains(fileString, email) {
		return ErrEmailExists
	}

	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.WriteString(email + "\n")
	return err
}

func (service *FileService) isFileExists() bool {
	_, err := os.Stat(FileName)
	return !os.IsNotExist(err)
}

func (service *FileService) ReadFromFile() ([]string, error) {
	fileData, err := os.ReadFile(FileName)
	if err != nil {
		return nil, err
	}

	fileString := string(fileData)
	emails := strings.Split(fileString, "\n")
	return emails, nil
}
