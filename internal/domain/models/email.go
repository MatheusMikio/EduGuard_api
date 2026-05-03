package models

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/models"
	"github.com/asaskevich/govalidator"
)

type Email struct {
	Email string `gorm:"not null"`
}

func NewEmail(email models.EmailRequest) (Email, *ErrorMessage) {
	if err := validateEmail(email.Email); err != nil {
		return Email{}, err
	}
	return Email{
		Email: email.Email,
	}, nil

}

func validateEmail(email string) *ErrorMessage {
	valid := govalidator.IsEmail(email)
	if valid == false {
		return CreateErrorMessage("Email", "Inválido")
	}
	return nil
}
