package service

import (
	"fmt"
	"net/smtp"
	"os"
)

type EmailService struct {
	auth smtp.Auth
}

func NewEmailService() *EmailService {
	auth := smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com")

	return &EmailService{
		auth: auth,
	}
}

func (service *EmailService) SendEmail(toEmail string, btcRate float64) error {
	from := os.Getenv("EMAIL_FROM")
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: BTC Rate\n\nCurrent BTC rate is %f UAH", from, toEmail, btcRate)

	err := smtp.SendMail("smtp.example.com:25", service.auth, from, []string{toEmail}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
