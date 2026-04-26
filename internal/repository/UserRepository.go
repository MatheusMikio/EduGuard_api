package repository

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
	"gorm.io/gorm"
)

type IUserRepositorty interface {
	base.ICrudRepository[schemas.User]
}

type UserRepository struct {
	base.CrudRepository[schemas.User]
}

func NewUserRepository(db *gorm.DB) IUserRepositorty {
	return &UserRepository{
		CrudRepository: base.CrudRepository[schemas.User]{
			Db: db,
		},
	}
}
