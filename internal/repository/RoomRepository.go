package repository

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
	"gorm.io/gorm"
)

type IRoomRepository interface {
	base.ICrudRepository[schemas.Room]
}

type RoomRepository struct {
	base.CrudRepository[schemas.Room]
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &RoomRepository{
		CrudRepository: base.CrudRepository[schemas.Room]{
			Db: db,
		},
	}
}
