package schemas

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/event"
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"gorm.io/gorm"
)

type EventType string

const (
	Empurrao   EventType = "Empurrão"
	CercoGrupo EventType = "Cerco em grupo"
	Confronto  EventType = "Confronto"
	Queda      EventType = "Queda forçada"
)

type EventStatus string

const (
	EventStatusPending   EventStatus = "PENDENTE"
	EventStatusReviewed  EventStatus = "EM REVISÃO"
	EventStatusCompleted EventStatus = "CONCLUIDO"
)

type Event struct {
	gorm.Model
	SchoolID     uint        `gorm:"not null"`
	School       *School     `gorm:"foreignKey:SchoolID"`
	CameraID     uint        `gorm:"not null"`
	Camera       *Camera     `gorm:"foreignKey:CameraID"`
	Type         EventType   `gorm:"type:event_type;not null"`
	RiskScore    float64     `gorm:"not null"`
	Status       EventStatus `gorm:"type:event_status;not null;default:'PENDENTE'"`
	StartedAt    time.Time   `gorm:"not null"`
	EndedAt      time.Time   `gorm:"not null"`
	VideoClipURL string      `gorm:"not null;type:text"`
	ReviewedByID *uint
	ReviewedBy   *User `gorm:"foreignKey:ReviewedByID"`
	ReviewedAt   *time.Time
	Notes        *string
}

func NewEvent(r event.Request) (Event, []*models.ErrorMessage) {
	if err := validateEvent(r); err != nil {
		return Event{}, err
	}
	return Event{}, nil
}

func validateEvent(request event.Request) []*models.ErrorMessage {
	return []*models.ErrorMessage{}
}
