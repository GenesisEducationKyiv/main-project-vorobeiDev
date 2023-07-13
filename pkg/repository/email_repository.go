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

func NewEmailRepository(filePath string) *EmailRepository {
	return &EmailRepository{
		filePath: filePath,
	}
}

var ErrEmailExists = errors.New("email already exists")
var ErrInvalidEmail = errors.New("invalid email address")

func (r *EmailRepository) Save(email string) error {
	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if !r.validateEmail(email) {
		return fmt.Errorf("%w: %s", ErrInvalidEmail, email)
	}

	emails, err := r.AllEmails()
	if err != nil {
		return err
	}

	if r.isEmailExists(email, emails) {
		return fmt.Errorf("%w: %s", ErrEmailExists, email)
	}

	_, err = fmt.Fprintln(file, email)

	return err
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

func (r *EmailRepository) isEmailExists(email string, emails []string) bool {
	for _, s := range emails {
		if strings.Contains(s, email) {
			return true
		}
	}

	return false
}

func (r *EmailRepository) validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
