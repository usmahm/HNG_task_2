package controllers

import (
	"fmt"
	"net/http"
	"person-app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetUser(c *gin.Context) {
	user_id := c.Param("user_id")
	fmt.Println(user_id)
	c.JSON(http.StatusOK, gin.H{"success": "Success GET"})
}

func CreateUser(c *gin.Context) {
	var new_user models.User

	if err := c.BindJSON(&new_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if validation_error := validate.Struct(new_user); validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	// Add to db

	fmt.Println(new_user)
	c.JSON(http.StatusOK, new_user)
}

func UpdateUser(c *gin.Context) {
	user_id := c.Param("user_id")
	fmt.Println(user_id)
	c.JSON(http.StatusOK, gin.H{"success": "Success PATCH"})
}

func DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")
	fmt.Println(user_id)
	c.JSON(http.StatusOK, gin.H{"success": "Success DEL"})
}
