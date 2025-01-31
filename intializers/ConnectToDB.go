package intializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetToDB() {
	var err error

	dsn := os.Getenv("DB")
	if dsn == "" {
		log.Fatal("Database URL not found in .env")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
}
