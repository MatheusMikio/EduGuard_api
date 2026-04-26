package base

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"github.com/MatheusMikio/eduGuard_api/internal/repository/base"
)

type ICrudService[TRequest any, TUpdate any, TResponse any] interface {
	GetAllAdmin(page uint64, size uint64) ([]*TResponse, models.ErrorMessage)
	GetAll(page uint64, size uint64, schoolId uint) ([]*TResponse, models.ErrorMessage)
	GetById(id uint) (*TResponse, models.ErrorMessage)
	Create(entity *TRequest) []*models.ErrorMessage
	Update(entity *TUpdate) []*models.ErrorMessage
	Delete(id uint) *models.Address
}

type CrudServiceConfig[TEntity any, TRequest any, TUpdate any, TResponse any] struct {
}

type CrudService[TEntity any, TRequest any, TUpdate any, TResponse any] struct {
	Repository base.ICrudRepository[TEntity]
	Config     CrudServiceConfig[TEntity, TRequest, TUpdate, TResponse]
}

func NewCrudService[TEntity any, TRequest any, TUpdate any, TResponse any](
	repository base.ICrudRepository[TEntity], config CrudServiceConfig[TEntity, TRequest, TUpdate, TResponse]) ICrudService[TRequest, TUpdate, TResponse] {
	return &CrudService[TEntity, TRequest, TUpdate, TResponse]{
		Repository: repository,
		Config:     config,
	}
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) GetAllAdmin(page uint64, size uint64) ([]*TResponse, models.ErrorMessage) {
	panic("unimplemented")
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) GetAll(page uint64, size uint64, schoolId uint) ([]*TResponse, models.ErrorMessage) {
	panic("unimplemented")
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) GetById(id uint) (*TResponse, models.ErrorMessage) {
	panic("unimplemented")
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) Create(entity *TRequest) []*models.ErrorMessage {
	panic("unimplemented")
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) Update(entity *TUpdate) []*models.ErrorMessage {
	panic("unimplemented")
}

func (cs *CrudService[TEntity, TRequest, TUpdate, TResponse]) Delete(id uint) *models.Address {
	panic("unimplemented")
}
