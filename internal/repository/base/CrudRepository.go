package base

import "gorm.io/gorm"

type ICrudRepository[T any] interface {
	GetAllAdmin(page uint64, size uint64) ([]*T, error)
	GetAll(page uint64, size uint64, schoolId uint) ([]*T, error)
	GetById(id uint64) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type CrudRepository[T any] struct {
	Db *gorm.DB
}

func (cr *CrudRepository[T]) GetAllAdmin(page uint64, size uint64) ([]*T, error) {
	var entities []*T
	offSet := int((page - 1) * size)
	if err := cr.Db.Offset(offSet).Limit(int(size)).Find(entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (cr *CrudRepository[T]) GetAll(page uint64, size uint64, schoolId uint) ([]*T, error) {
	var entities []*T
	offSet := int((page - 1) * size)
	if err := cr.Db.Where("school_id = ?", schoolId).Offset(offSet).Limit(int(size)).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil

}

func (cr *CrudRepository[T]) GetById(id uint64) (*T, error) {
	var entity T
	if err := cr.Db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (cr *CrudRepository[T]) Create(entity *T) error {
	if err := cr.Db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CrudRepository[T]) Update(entity *T) error {
	if err := cr.Db.Save(entity).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CrudRepository[T]) Delete(entity *T) error {
	if err := cr.Db.Delete(entity).Error; err != nil {
		return err
	}

	return nil
}
