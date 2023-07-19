package user

import (
	"errors"
	"net/mail"
)

type User struct {
	Email string `json:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}

func (d *User) ValidateUser() error {
	if !d.isEmailValid(d.Email) {
		return errors.New("invalid request")
	}

	return nil
}

func (d *User) isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
