package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = initPostGreSQL()
	if err != nil {
		return fmt.Errorf("Erro ao inicializar o DB: %s", err)
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetPepper() string {
	return os.Getenv("PASSWORD_PEPPER")
}
