package repository

type IEmailRepository interface {
	Save(email string) error
	AllEmails() ([]string, error)
}

type Repositories struct {
	EmailRepository IEmailRepository
}

func NewRepositories(emailRepository IEmailRepository) *Repositories {
	return &Repositories{
		EmailRepository: emailRepository,
	}
}
