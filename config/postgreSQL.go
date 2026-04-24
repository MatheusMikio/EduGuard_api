package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostGreSQL() (*gorm.DB, error) {
	_ = godotenv.Load()
	_ = godotenv.Load(filepath.Join("..", "..", ".env"))

	dbConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sskmode=disabled",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// USAR EM DEV
	// err := db.AutoMigrate()

	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
