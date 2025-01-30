package intializers

import "github.com/zenithBasit/jwt-authentication/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}