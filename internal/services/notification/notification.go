package notification

import (
	"fmt"
	"net/smtp"
)

type Service struct {
	auth      smtp.Auth
	emailFrom string
}

func NewNotificationService(emailFrom string) *Service {
	auth := smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com")

	return &Service{
		auth:      auth,
		emailFrom: emailFrom,
	}
}

func (s *Service) Send(toEmail string, btcRate float64) error {
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
