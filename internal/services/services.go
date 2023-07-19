package services

import (
	user2 "github.com/vorobeiDev/crypto-client/internal/domain/user"
	"github.com/vorobeiDev/crypto-client/internal/repositories/filesystem"
	"github.com/vorobeiDev/crypto-client/internal/services/currency"
	"github.com/vorobeiDev/crypto-client/internal/services/notification"
	"github.com/vorobeiDev/crypto-client/internal/services/user"
	"os"
)

type IFSRepository interface {
	Save(user *user2.User) error
	EmailExist(email string) (bool, error)
	AllEmails() ([]string, error)
}

type Repositories struct {
	filesRepo IFSRepository
}

type Services struct {
	UserService         *user.Service
	CurrencyService     *currency.Service
	NotificationService *notification.Service
}

func NewServices() *Services {
	repositories := createRepository()

	return &Services{
		UserService:         user.NewUserService(repositories.filesRepo),
		CurrencyService:     currency.NewCurrencyService(os.Getenv(os.Getenv("EMAIL_FROM"))),
		NotificationService: notification.NewNotificationService(os.Getenv("")),
	}
}

func createRepository() *Repositories {
	return &Repositories{
		filesRepo: filesystem.NewFSRepository(os.Getenv("DB_FILE_NAME")),
	}
}
