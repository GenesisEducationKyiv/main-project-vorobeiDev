package service

type Services struct {
	CurrencyService *CurrencyService
	EmailService    *EmailService
}

func NewServices() *Services {
	return &Services{
		CurrencyService: NewCurrencyService(),
		EmailService:    NewEmailService(),
	}
}
