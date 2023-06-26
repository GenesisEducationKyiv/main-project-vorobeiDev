package service

type Services struct {
	CurrencyService *CurrencyService
	EmailService    *EmailService
}

func NewServices(currencyService *CurrencyService, emailService *EmailService) *Services {
	return &Services{
		CurrencyService: currencyService,
		EmailService:    emailService,
	}
}
