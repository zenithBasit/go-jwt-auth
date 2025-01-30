package intializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetToDB() {
	var err error

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil{
		panic("Failed to connect to DB")
	}
}
