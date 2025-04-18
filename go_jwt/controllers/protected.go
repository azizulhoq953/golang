package controllers

import (
	"github.com/azizulhoq953/jwtToken/database"
	"github.com/azizulhoq953/jwtToken/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile is a controller function that retrieves the user profile from the database
// based on the email provided in the authorization middleware.
// It returns a 404 status code if the user is not found,
// and a 500 status code if an error occurs while retrieving the user profile.

func Profile(c *gin.Context) {
	// Initialize a user model
	var user models.User
	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	// Query the database for the user
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)
	// If the user is not found, return a 404 status code
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}
	// If an error occurs while retrieving the user profile, return a 500 status code
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}
	// Set the user's password to an empty string
	user.Password = ""
	// Return the user profile with a 200 status code
	c.JSON(200, user)
}
