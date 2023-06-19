package service

import "net/mail"

type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (service *ValidationService) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	
	return err == nil
}
