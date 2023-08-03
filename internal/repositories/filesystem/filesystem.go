package filesystem

import (
	"log"
	"os"
)

type Repository struct {
	filePath string
}

func NewFSRepository(filePath string) *Repository {
	createFileDB(filePath)

	return &Repository{
		filePath: filePath,
	}
}

func createFileDB(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(filePath)
			if createErr != nil {
				log.Fatal(createErr)
				return
			}
			defer file.Close()
		}
	}
}
