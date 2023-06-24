package repository_test

import (
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	tmpFileName := "./test_emails.txt"
	t.Setenv("DB_FILE_NAME", tmpFileName)

	emailRepo := repository.NewEmailRepository()

	if emailRepo.IsFileExists() {
		t.Fatalf("The file was expected not to exist but was found to exist")
	}

	err := emailRepo.Save("test@example.com")

	if err != nil {
		t.Fatalf("An error was received while saving the email address: %v", err)
	}

	if !emailRepo.IsFileExists() {
		t.Fatalf("The file was expected to exist but found to be missing")
	}

	os.Remove(tmpFileName)
}

func TestGetAllEmails(t *testing.T) {
	tmpFileName := "./test_emails.txt"
	t.Setenv("DB_FILE_NAME", tmpFileName)

	file, err := os.Create(tmpFileName)

	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	defer file.Close()

	_, err = file.WriteString("test1@example.com\ntest2@example.com\ntest3@example.com")
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	emailRepo := repository.NewEmailRepository()

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

	expectedEmails := []string{"test1@example.com", "test2@example.com", "test3@example.com"}
	for i, email := range emails {
		if email != expectedEmails[i] {
			t.Errorf("Invalid email address received. Expected: %s, Received: %s", expectedEmails[i], email)
		}
	}

	defer func() {
		file.Close()

		errSync := file.Sync()
		if errSync != nil {
			t.Errorf("Failed to sync file: %v", errSync)
		}

		errRemove := os.Remove(tmpFileName)
		if errRemove != nil {
			t.Errorf("Error deleting temporary file: %v", errRemove)
		}
	}()
}
