package repository_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/repository"
)

const defaultTestEmail = "test@example.com"
const invalidTestEmail = "testemailexamplecom"

func TestMain(m *testing.M) {
	if envErr := godotenv.Load("../../.env.test"); envErr != nil {
		log.Fatal("Error loading .env file.", envErr)
	}

	exitCode := m.Run()

	os.Remove(os.Getenv("DB_FILE_NAME"))
	os.Exit(exitCode)
}

func TestEmailRepository_Save(t *testing.T) {
	filePath := os.Getenv("DB_FILE_NAME")
	r := repository.NewEmailRepository(filePath)

	if err := r.Save(defaultTestEmail); err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	err := r.Save(defaultTestEmail)
	if err == nil || !errors.Is(err, repository.ErrEmailExists) {
		t.Fatalf("Expected '%v' error, got %v", repository.ErrEmailExists, err)
	}

	err = r.Save(invalidTestEmail)
	if err == nil || !errors.Is(err, repository.ErrInvalidEmail) {
		t.Fatalf("Expected '%v' error, got %v", repository.ErrInvalidEmail, err)
	}
}

func TestEmailRepository_AllEmails(t *testing.T) {
	filePath := os.Getenv("DB_FILE_NAME")
	r := repository.NewEmailRepository(filePath)

	err := r.Save(defaultTestEmail)
	if err == nil || !errors.Is(err, repository.ErrEmailExists) {
		t.Fatalf("Expected '%v' error, got %v", repository.ErrEmailExists, err)
	}

	emails, err := r.AllEmails()
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	if len(emails) != 1 || emails[0] != defaultTestEmail {
		t.Errorf("Expected single email '%s', but got '%s'", defaultTestEmail, emails[0])
	}
}
