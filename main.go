package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zenithBasit/jwt-authentication/controllers"
	"github.com/zenithBasit/jwt-authentication/intializers"
	"github.com/zenithBasit/jwt-authentication/middleware"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnetToDB()
	intializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
