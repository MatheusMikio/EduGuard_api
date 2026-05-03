package schemas

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/school"
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name           string      `gorm:"not null"`
	Cnpj           models.Cnpj `gorm:"embedded"`
	models.Address `gorm:"embedded"`
	PublicID       uuid.UUID `gorm:"not null"`
	Rooms          []Room
	Users          []User
	Cameras        []Camera
}

func NewSchool(r school.Request) (School, []*models.ErrorMessage) {
	if err := validateSchool(r); err != nil {
		return School{}, err
	}
	return School{}, nil
}

func validateSchool(request school.Request) []*models.ErrorMessage {
	return []*models.ErrorMessage{}
}
