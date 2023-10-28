package controllers

import (
	"crud/shopDB/authjwt"
	"crud/shopDB/models"

	"crud/shopDB/pkg/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
	Email    string `json:"email" `
	Password string `json:"password"`
}

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

func Signup(c *gin.Context) {
	var user models.RegistrationForm
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs ",
		})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"Error": "Error Hashing Password",
		})
		c.Abort()
		return
	}
	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Register",
	})
}

//login
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.RegistrationForm
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	result := config.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := authjwt.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedtoken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedtoken,
	}
	c.JSON(200, tokenResponse)
}

func Profile(c *gin.Context) {
	// Initialize a user model
	var user models.RegistrationForm
	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	// Query the database for the user
	result := config.GlobalDB.Where("email = ?", email.(string)).First(&user)
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

//End Registration

var NewShop models.Shop

func GetShop(all *gin.Context) {
	var Shops []models.Shop
	config.GlobalDB.Find(&Shops)
	all.JSON(http.StatusOK, gin.H{"AllData": Shops})
}

func GetShopById(i *gin.Context) {

	var shop models.Shop

	if err := config.GlobalDB.Where("id = ?", i.Param("id")).First(&shop).Error; err != nil {
		i.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	i.JSON(http.StatusOK, gin.H{"data": shop})
	// Set the user's password to an empty string
	// user.Password = ""
	// // Return the user ID with a 200 status code
	// c.JSON(200, user)

}

func CreateShop(c *gin.Context) {

	var user models.Shop
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs ",
		})
		c.Abort()
		return
	}

	err = user.CreatShop()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Shop Created",
	})

}

//update with gin framework
func UpdateShop(c *gin.Context) {
	// Get model if exist
	var shop models.Shop
	if err := config.GlobalDB.Where("id = ?", c.Param("id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Shop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.GlobalDB.Model(&shop).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

//DELETED DATA USING GIN
func Deleted(d *gin.Context) {
	var shop models.Shop
	if err := config.GlobalDB.Where("id = ?", d.Param("id")).First(&shop).Error; err != nil {
		d.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.GlobalDB.Delete(&shop)

	d.JSON(http.StatusOK, gin.H{"Deleted": true})
}
