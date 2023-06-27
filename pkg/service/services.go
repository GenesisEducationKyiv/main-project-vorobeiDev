package service

type Services struct {
	CurrencyService ICurrencyService
	EmailService    IEmailService
}

func NewServices(currencyService ICurrencyService, emailService IEmailService) *Services {
	return &Services{
		CurrencyService: currencyService,
		EmailService:    emailService,
	}
}
