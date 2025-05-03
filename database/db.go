package database

import (
	"back/calculate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := DB.AutoMigrate(&calculate.Calculations{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
