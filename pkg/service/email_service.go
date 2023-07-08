package service

import (
	"fmt"
	"github.com/vorobeiDev/crypto-client/pkg/repository"
	"net/smtp"
	"os"
)

type EmailService struct {
	auth            smtp.Auth
	emailFrom       string
	emailRepository *repository.EmailRepository
}

func NewEmailService(emailRepository *repository.EmailRepository) *EmailService {
	emailFrom := os.Getenv("EMAIL_FROM")
	auth := smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com")

	return &EmailService{
		auth:            auth,
		emailFrom:       emailFrom,
		emailRepository: emailRepository,
	}
}

func (s *EmailService) SaveEmail(email string) error {
	return s.emailRepository.Save(email)
}

func (s *EmailService) GetAllEmails() ([]string, error) {
	return s.emailRepository.AllEmails()
}

func (s *EmailService) SendEmail(toEmail string, btcRate float64) error {
	msg := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: BTC Rate\n\nCurrent BTC rate is %f UAH",
		s.emailFrom,
		toEmail,
		btcRate,
	)

	err := smtp.SendMail(
		"smtp.example.com:25",
		s.auth,
		s.emailFrom,
		[]string{toEmail},
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}
