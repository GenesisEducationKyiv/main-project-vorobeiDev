package filesystem_test

import (
	"github.com/vorobeiDev/crypto-client/internal/domain/user"
	"github.com/vorobeiDev/crypto-client/internal/repositories/filesystem"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const (
	expectedEmail1 = "test1@example.com"
	expectedEmail2 = "test2@example.com"
	expectedEmail3 = "test3@example.com"
)

func TestMain(m *testing.M) {
	if envErr := godotenv.Load("../../../.env.test"); envErr != nil {
		log.Fatal("Error loading .env file.", envErr)
	}

	exitCode := m.Run()

	os.Remove(os.Getenv("DB_FILE_NAME"))
	os.Exit(exitCode)
}

func TestSave(t *testing.T) {
	filePath := os.Getenv("DB_FILE_NAME")
	repo := filesystem.NewFSRepository(filePath)
	validUser := user.NewUser(expectedEmail1)

	if err := repo.Save(validUser); err != nil {
		t.Fatalf("Save returned an error: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("File reading error: %v", err)
	}

	expectedContent := expectedEmail1 + "\n"
	if string(content) != expectedContent {
		t.Errorf("Unexpected file content. Expected: %q, got: %q", expectedContent, string(content))
	}
}

func TestAllEmails(t *testing.T) {
	repo := filesystem.NewFSRepository(os.Getenv("DB_FILE_NAME"))
	expectedEmails := []string{expectedEmail1, expectedEmail2, expectedEmail3}

	for _, email := range expectedEmails {
		if err := repo.Save(user.NewUser(email)); err != nil {
			t.Fatalf("Save returned an error: %v", err)
		}
	}

	emails, err := repo.AllEmails()
	if err != nil {
		t.Errorf("AllEmails returned an error: %v", err)
	}

	if len(emails) != len(expectedEmails) {
		t.Errorf("Unexpected number of emails. Expected: %d, got: %d", len(expectedEmails), len(emails))
	}

	for i, email := range emails {
		if email != expectedEmails[i] {
			t.Errorf("Unexpected email at index %d. Expected: %q, got: %q", i, expectedEmails[i], email)
		}
	}
}

func TestEmailExist(t *testing.T) {
	repo := filesystem.NewFSRepository(os.Getenv("DB_FILE_NAME"))
	expectedEmails := []string{expectedEmail1, expectedEmail2}

	for _, email := range expectedEmails {
		if err := repo.Save(user.NewUser(email)); err != nil {
			t.Fatalf("Save returned an error: %v", err)
		}
	}

	exists, err := repo.EmailExist(expectedEmail1)
	if err != nil {
		t.Errorf("EmailExist returned an error: %v", err)
	}

	if !exists {
		t.Error("Expected email to exist, but it was not found")
	}

	exists, err = repo.EmailExist(expectedEmail3)
	if err != nil {
		t.Errorf("EmailExist returned an error: %v", err)
	}

	if exists {
		t.Error("Expected email to not exist, but it was found")
	}
}
