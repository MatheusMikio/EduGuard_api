package schemas

import (
	"time"

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
	EventStatusReviewd   EventStatus = "EM REVISÃO"
	EventStatusCompleted EventStatus = "CONCLUIDO"
)

type Event struct {
	gorm.Model
	SchoolID     uint        `gorm:"not null"`
	School       *School     `gorm:"foreignKey:SchoolID"`
	CameraID     uint        `gorm:"not null"`
	Camera       *Camera     `gorm:"foreignKey:CameraID"`
	Type         EventType   `gorm:"not null"`
	RiskScore    float64     `gorm:"not null"`
	Status       EventStatus `gorm:"not null;default:'PENDENTE'"`
	StartedAt    time.Time   `gorm:"not null"`
	EndedAt      time.Time   `gorm:"not null"`
	VideoClipURL string      `gorm:"not null;type:text"`
	ReviewedByID uint        `gorm:"not null"`
	ReviewedBy   *User       `gorm:"foreignKey:ReviewedByID"`
	ReviewedAt   *time.Time
	Notes        *string
}
