package repository_test

import (
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/vorobeiDev/crypto-client/pkg/repository"
)

func TestCreateFile(t *testing.T) {
	envErr := godotenv.Load("../../.env.test")
	if envErr != nil {
		log.Fatal("Error loading .env file.", envErr)
	}

	emailRepo := repository.NewEmailRepository()

	if emailRepo.IsFileExists() {
		t.Fatalf("The file was expected not to exist but was found to exist")
	}

	saveErr := emailRepo.Save("test@example.com")
	if saveErr != nil {
		t.Fatalf("An error was received while saving the email address: %v", saveErr)
	}

	if !emailRepo.IsFileExists() {
		t.Fatalf("The file was expected to exist but found to be missing")
	}

	defer func() {
		if removeErr := emailRepo.RemoveFile(); removeErr != nil {
			t.Fatalf(removeErr.Error())
		}
	}()
}

func TestGetAllEmails(t *testing.T) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file.", err)
	}

	expectedEmails := []string{"test1@example.com", "test2@example.com", "test3@example.com"}

	emailRepo := repository.NewEmailRepository()

	for _, email := range expectedEmails {
		if saveErr := emailRepo.Save(email); saveErr != nil {
			t.Fatalf("An error with email saving!")
		}
	}

	emails, err := emailRepo.GetAllEmails()
	if err != nil {
		t.Fatalf("An error was received while retrieving the list of email addresses: %v", err)
	}

	expectedCount := 3
	if len(emails) != expectedCount {
		t.Errorf(
			"Incorrect number of email addresses received. Expected: %d, Received: %d",
			expectedCount,
			len(emails),
		)
	}

	for i, email := range emails {
		if email != expectedEmails[i] {
			t.Errorf("Invalid email address received. Expected: %s, Received: %s", expectedEmails[i], email)
		}
	}

	defer func() {
		if removeErr := emailRepo.RemoveFile(); removeErr != nil {
			t.Fatalf(removeErr.Error())
		}
	}()
}
