package repository

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
	"gorm.io/gorm"
)

type ISchoolRepository interface {
	base.ICrudRepository[schemas.School]
}

type SchoolRepository struct {
	base.CrudRepository[schemas.School]
}

func NewSchoolRepository(db *gorm.DB) ISchoolRepository {
	return &SchoolRepository{
		CrudRepository: base.CrudRepository[schemas.School]{
			Db: db,
		},
	}
}
