package repository

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
	"gorm.io/gorm"
)

type IEventRepository interface {
	base.ICrudRepository[schemas.Event]
}

type EventRepository struct {
	base.CrudRepository[schemas.Event]
}

func NewEventRepository(db *gorm.DB) IEventRepository {
	return &EventRepository{
		CrudRepository: base.CrudRepository[schemas.Event]{
			Db: db,
		},
	}
}
