package models

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/models"
)

type Address struct {
	Street     string  `json:"street"`
	Number     string  `json:"number"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Cep        string  `json:"cep"`
	Country    string  `json:"country"`
	Complement *string `json:"complement"`
}

func NewAddress(r models.AddressRequest) (Address, []*ErrorMessage) {
	err := validateAddress(r)
	if len(err) > 0 {
		return Address{}, err
	}
	return Address{
		Street:     r.Street,
		Number:     r.Number,
		City:       r.City,
		State:      r.State,
		Cep:        r.Cep,
		Country:    r.Country,
		Complement: r.Complement,
	}, nil
}

func validateAddress(request models.AddressRequest) []*ErrorMessage {
	return []*ErrorMessage{}
}
