package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenithBasit/jwt-authentication/intializers"
	"github.com/zenithBasit/jwt-authentication/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email and pass from req  body
	var body struct {
		Email    string
		Password string
	}


    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to read body",
        })
        return
    }
	// Hashing the pass
	hash, err:= bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	// create user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := intializers.DB.Create(&user)
	if result.Error !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
