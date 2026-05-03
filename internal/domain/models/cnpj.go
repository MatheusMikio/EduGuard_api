package models

import (
	dto "github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/models"
	"github.com/MatheusMikio/eduGuard_api/internal/utils"
	c "github.com/jfelipearaujo/cpfcnpj/cnpj"
)

type Cnpj struct {
	Cnpj string `gorm:"not null"`
}

func NewCnpj(cnpj dto.CnpjRequest) (Cnpj, *ErrorMessage) {
	if err := validateCnpj(cnpj.Cnpj); err != nil {
		return Cnpj{}, err
	}
	return Cnpj{
		Cnpj: cnpj.Cnpj,
	}, nil
}

func validateCnpj(cnpj string) *ErrorMessage {
	if err := validateFormat(cnpj); err != nil {
		return err
	}

	if err := utils.ValidateCnpjExists(cnpj); err != nil {
		return CreateErrorMessage("Cnpj", err.Error())
	}

	return nil
}

func validateFormat(cnpj string) *ErrorMessage {
	svc := c.New()
	if err := svc.IsValid(cnpj); err != nil {
		return CreateErrorMessage("Cnpj", "Inválido")
	}
	return nil
}
