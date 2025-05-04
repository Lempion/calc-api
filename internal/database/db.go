package database

import (
	"back/internal/calculationService"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := DB.AutoMigrate(&calculationService.Calculations{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return DB, nil
}
