package container

import "gorm.io/gorm"

type Container struct{}

func NewContainer(db *gorm.DB) *Container {
	//Repositorios

	// Serviços
	return &Container{}
}
