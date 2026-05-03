package schemas

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/room"
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Floor    string `gorm:"not null"`
	SchoolID uint   `gorm:"not null"`
	School   School `gorm:"foreignKey:'SchoolID'"`
	Cameras  []Camera
}

func NewRoom(r room.Request) (Room, []*models.ErrorMessage) {
	if err := validateRoom(r); err != nil {
		return Room{}, err
	}
	return Room{}, nil
}

func validateRoom(request room.Request) []*models.ErrorMessage {
	return []*models.ErrorMessage{}
}
