package base

import "gorm.io/gorm"

type IBaseRepository[T any] interface {
	GetAll(page uint64, size uint64) ([]*T, error)
	GetById(id uint64) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type CrudRepository[T any] struct {
	Db *gorm.DB
}

func (b *CrudRepository[T]) GetAll(page uint64, size uint64) ([]*T, error) {
	var entities []T
	offSet := int((page - 1) * size)
	if err := b.Db.Offset(offSet).Limit(int(size)).Find(&entities).Error; err != nil {
		return nil, err
	}

	result := make([]*T, len(entities))
	for i := range entities {
		result[i] = &entities[i]
	}

	return result, nil
}

func (b *CrudRepository[T]) GetById(id uint64) (*T, error) {
	var entity T
	if err := b.Db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (b *CrudRepository[T]) Create(entity *T) error {
	if err := b.Db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (b *CrudRepository[T]) Update(entity *T) error {
	if err := b.Db.Save(entity).Error; err != nil {
		return err
	}

	return nil
}

func (b *CrudRepository[T]) Delete(entity *T) error {
	if err := b.Db.Delete(entity).Error; err != nil {
		return err
	}

	return nil
}
