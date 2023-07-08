package repository_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/vorobeiDev/crypto-client/pkg/repository"
)

const DefaultTestEmail = "test@example.com"

func TestMain(m *testing.M) {
	if envErr := godotenv.Load("../../.env.test"); envErr != nil {
		log.Fatal("Error loading .env file.", envErr)
	}

	exitCode := m.Run()

	os.Remove(os.Getenv("DB_FILE_NAME"))
	os.Exit(exitCode)
}

func TestEmailRepository_Save(t *testing.T) {
	r := repository.NewEmailRepository()

	if err := r.Save(DefaultTestEmail); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	err := r.Save(DefaultTestEmail)
	if err == nil || !errors.Is(err, repository.ErrEmailExists) {
		t.Fatalf("expected '%v' error, got %v", repository.ErrEmailExists, err)
	}

	err = r.Save("testemailexamplecom")
	if err == nil || !errors.Is(err, repository.ErrInvalidEmail) {
		t.Fatalf("expected '%v' error, got %v", repository.ErrInvalidEmail, err)
	}
}

func TestEmailRepository_AllEmails(t *testing.T) {
	r := repository.NewEmailRepository()

	err := r.Save(DefaultTestEmail)
	if err == nil || !errors.Is(err, repository.ErrEmailExists) {
		t.Fatalf("expected '%v' error, got %v", repository.ErrEmailExists, err)
	}

	emails, err := r.AllEmails()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !contains(emails, DefaultTestEmail) {
		t.Fatalf("expected all emails includes, %v, but thay not. All emails: %v", DefaultTestEmail, emails)
	}
}

func contains(s []string, searchString string) bool {
	for _, v := range s {
		if v == searchString {
			return true
		}
	}

	return false
}
