package repository

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
	"gorm.io/gorm"
)

type ICameraRepository interface {
	base.IBaseRepository[schemas.Camera]
}

type CameraRepository struct {
	base.CrudRepository[schemas.Camera]
}

func NewCameraRepository(db *gorm.DB) ICameraRepository {
	return &CameraRepository{
		CrudRepository: base.CrudRepository[schemas.Camera]{
			Db: db,
		},
	}
}
