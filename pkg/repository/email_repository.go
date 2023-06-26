package repository

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strings"
)

type EmailRepository struct {
	filePath string
}

func NewEmailRepository() *EmailRepository {
	filePath := os.Getenv("DB_FILE_NAME")

	return &EmailRepository{
		filePath: filePath,
	}
}

var ErrEmailExists = errors.New("email already exists")

func (r *EmailRepository) Save(email string) error {
	if err := r.CreateFileIfNotExist(); err != nil {
		return err
	}

	emails, err := r.AllEmails()
	if err != nil {
		return err
	}

	if r.IsEmailExists(email, emails) {
		return fmt.Errorf("%w: %s", ErrEmailExists, email)
	}

	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, email)

	return err
}

func (r *EmailRepository) IsFileExists() bool {
	_, err := os.Stat(r.filePath)
	return !os.IsNotExist(err)
}

func (r *EmailRepository) AllEmails() ([]string, error) {
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	fileString := string(fileData)
	emails := strings.Split(fileString, "\n")

	var trimmedEmails []string

	for _, email := range emails {
		trimmedEmail := strings.TrimSpace(email)
		if trimmedEmail != "" {
			trimmedEmails = append(trimmedEmails, trimmedEmail)
		}
	}

	return trimmedEmails, nil
}

func (r *EmailRepository) CreateFileIfNotExist() error {
	isFileExists := r.IsFileExists()
	if !isFileExists {
		_, createErr := os.Create(r.filePath)
		if createErr != nil {
			return createErr
		}
	}

	return nil
}

func (r *EmailRepository) IsEmailExists(email string, emails []string) bool {
	for _, s := range emails {
		if strings.Contains(s, email) {
			return true
		}
	}

	return false
}

func (r *EmailRepository) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (r *EmailRepository) RemoveFile() error {
	err := os.Remove(r.filePath)
	if err != nil {
		return err
	}

	return nil
}
