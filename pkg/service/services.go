package service

type Services struct {
	CurrencyService   *CurrencyService
	EmailService      *EmailService
	FileService       *FileService
	ValidationService *ValidationService
}

func NewServices() *Services {
	return &Services{
		CurrencyService:   NewCurrencyService(),
		EmailService:      NewEmailService(),
		FileService:       NewFileService(),
		ValidationService: NewValidationService(),
	}
}
