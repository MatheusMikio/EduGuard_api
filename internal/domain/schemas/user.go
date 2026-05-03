package schemas

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/user"
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleUser string

const (
	Admin       RoleUser = "ADMIN"
	Coordinator RoleUser = "COORDENADOR"
	Reviewer    RoleUser = "VIEWER"
)

type User struct {
	gorm.Model
	Name          string       `gorm:"not null"`
	Email         models.Email `gorm:"embedded"`
	PasswordHash  string       ``
	Role          RoleUser     `gorm:"type:role_user;not null;default:'VIEWER'"`
	SchoolID      uint         `gorm:"not null"`
	School        School       `gorm:"foreignKey:SchoolID"`
	RefreshToken  *string      `gorm:"type:text"`
	TokenIssuedAt *time.Time
	PublicId      uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	ImageURL      *string   `gorm:"type:text"`
}

func NewUSer(r *user.Request, email models.Email) (*User, []*models.ErrorMessage) {
	if errors := validateUser(r); errors != nil {
		return &User{}, errors
	}

	return &User{
		Name:         r.Name,
		Email:        email,
		PasswordHash: r.Password,
		Role:         RoleUser(r.Role),
		SchoolID:     r.SchoolID,
		PublicId:     uuid.New(),
		ImageURL:     r.ImageURL,
	}, nil

}

func validateUser(request *user.Request) []*models.ErrorMessage {
	return []*models.ErrorMessage{}
}
