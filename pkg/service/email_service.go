package service

type IEmailRepository interface {
	Save(email string) error
	AllEmails() ([]string, error)
}

type EmailService struct {
	repository IEmailRepository
}

func NewEmailService(emailRepository IEmailRepository) *EmailService {
	return &EmailService{
		repository: emailRepository,
	}
}

func (s *EmailService) Save(email string) error {
	return s.repository.Save(email)
}

func (s *EmailService) AllEmails() ([]string, error) {
	return s.repository.AllEmails()
}
