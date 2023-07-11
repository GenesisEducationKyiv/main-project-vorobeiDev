package service

import "golang.org/x/net/context"

type ICurrencyService interface {
	GetPrice(ctx context.Context, from string, to string) (float64, error)
}

type IEmailSenderService interface {
	Send(toEmail string, btcRate float64) error
}

type Services struct {
	CurrencyService ICurrencyService
	EmailSender     IEmailSenderService
}

func NewServices(currencyService ICurrencyService, emailService IEmailSenderService) *Services {
	return &Services{
		CurrencyService: currencyService,
		EmailSender:     emailService,
	}
}
