package controllers

import (
	"net/http"

	"github.com/mkumar2307/go-restapi/config"
	"github.com/mkumar2307/go-restapi/models"

	"github.com/gin-gonic/gin"
)

// GET /users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// POST /users
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GET /users/:id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
