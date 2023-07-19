package filesystem

import (
	"fmt"
	"os"
	"strings"

	"github.com/vorobeiDev/crypto-client/internal/domain/user"
)

func (fs *Repository) Save(user *user.User) error {
	file, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, user.Email)

	return err
}

func (fs *Repository) AllEmails() ([]string, error) {
	content, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil, err
	}

	fileString := string(content)
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

func (fs *Repository) EmailExist(email string) (bool, error) {
	emails, err := fs.AllEmails()
	if err != nil {
		return false, err
	}

	for _, storedEmail := range emails {
		if storedEmail == email {
			return true, nil
		}
	}

	return false, nil
}
