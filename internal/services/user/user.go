package user

import (
	"fmt"
	"github.com/vorobeiDev/crypto-client/internal/domain/user"
)

type IUserRepository interface {
	Save(user *user.User) error
	EmailExist(email string) (bool, error)
	AllEmails() ([]string, error)
}

type Service struct {
	userRepository IUserRepository
}

func NewUserService(repository IUserRepository) *Service {
	return &Service{
		userRepository: repository,
	}
}

func (s *Service) CreateNewUser(user *user.User) error {
	exist, err := s.userRepository.EmailExist(user.Email)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("%w: %s", ErrUserExists, user.Email)
	}

	return s.userRepository.Save(user)
}

func (s *Service) AllEmails() ([]string, error) {
	return s.userRepository.AllEmails()
}
