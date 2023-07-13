package service

import "golang.org/x/net/context"

type ICurrencyService interface {
	GetPrice(ctx context.Context, from string, to string) (float64, error)
}

type IEmailSenderService interface {
	Send(toEmail string, btcRate float64) error
}

type IFileService interface {
	Save(email string) error
	AllEmails() ([]string, error)
}

type Services struct {
	CurrencyService    ICurrencyService
	EmailSenderService IEmailSenderService
	EmailService       IFileService
}

func NewServices(
	currencyService ICurrencyService,
	emailService IEmailSenderService,
	fileService IFileService,
) *Services {
	return &Services{
		CurrencyService:    currencyService,
		EmailSenderService: emailService,
		EmailService:       fileService,
	}
}
