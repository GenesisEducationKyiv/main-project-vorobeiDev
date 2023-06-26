package repository

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strings"
)

type EmailRepository struct {
	fileName string
}

func NewEmailRepository() *EmailRepository {
	fileName := os.Getenv("DB_FILE_NAME")

	return &EmailRepository{
		fileName: fileName,
	}
}

var ErrEmailExists = errors.New("email already exists")

func (r *EmailRepository) Save(email string) error {
	if !r.isFileExists() {
		_, err := os.Create(r.fileName)
		if err != nil {
			return err
		}
	}

	emails, err := r.AllEmails()
	if err != nil {
		return err
	}

	if r.IsEmailExists(email, emails) {
		return fmt.Errorf("%w: %s", ErrEmailExists, email)
	}

	file, err := os.OpenFile(r.fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(email + "\n")

	return err
}

func (r *EmailRepository) isFileExists() bool {
	_, err := os.Stat(r.fileName)
	return !os.IsNotExist(err)
}

func (r *EmailRepository) AllEmails() ([]string, error) {
	fileData, err := os.ReadFile(r.fileName)
	if err != nil {
		return nil, err
	}

	fileString := string(fileData)
	emails := strings.Split(fileString, "\n")

	return emails, nil
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
