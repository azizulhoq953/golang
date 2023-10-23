package controllers

import (
	"crud/shopDB/authjwt"
	"crud/shopDB/models"
	"crud/shopDB/pkg/config"
	"crud/shopDB/utilis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//Resigtrtion
// func CreatRegistration(w http.ResponseWriter, r *http.Request) {

// 	CreatRegistration := &models.RegistrationForm{}

// 	fmt.Println(r.Body)

// 	utilis.ParseBody(r, CreatRegistration)
// 	s := CreatRegistration.CreatRegistration()
// 	Reg, _ := json.Marshal(s)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(Reg)

// }

//public data
// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
	Email    string `json:"email" `
	Password string `json:"password"`
}

// type getId struct {
// 	id int64 `json:"id" `
// }

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

//end
// func GetAllRegistration(w http.ResponseWriter, r *http.Request) {
// 	newRegis := models.GetAllRegistration()
// 	reg, _ := json.Marshal(newRegis)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(reg)
// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	email := vars["email"]

// 	LoginDetails, _ := models.Login(email)
// 	res, _ := json.Marshal(LoginDetails)

// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// 	fmt.Print(r)

// vars := mux.Vars(r)
// // password := vars["pass"]
// // Password, err := strconv.ParseInt(password, 0, 0)
// // PasswordDetails, _ := models.Login(Password)
// // pas, _ := json.Marshal(PasswordDetails)

// //email
// email := vars["email"]
// // Email, err := strconv.ParseInt(email, 0, 0)
// // if err != nil {
// // 	fmt.Println("User Name or password Is Incorrect")
// // }

// LoginDetails, _ := models.RegistrationForm.Email()

// comb := LoginDetails == email
// if err != nil {
// 	fmt.println("userName password Wrong")
// }

// log, _ := json.Marshal(comb)
// // pass,_ := models.Login()
// w.Header().Set("Content-Type", "pkglication/json")
// w.WriteHeader(http.StatusOK)
// w.Write(log)
// // w.Write(pas)
// //password verify

// }
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

// func GetAllShop(c *gin.Context) {
// 	mapd := config.GlobalDB
// 	var Shops []models.Shop
// 	_, err := mapd.Select(&Shops, "select * from Shop")
// 	if err == nil {
// 		c.JSON(200, Shops)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }

func GetShop(all *gin.Context) {
	var Shops []models.Shop
	config.GlobalDB.Find(&Shops)
	all.JSON(http.StatusOK, gin.H{"AllData": Shops})
}

// func GetShop(w http.ResponseWriter, r *http.Request) {
// 	newShop := models.GetAllShop()
// 	res, _ := json.Marshal(newShop)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

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

	// vars := mux.Vars(r)
	// shopId := vars["shopId"]
	// ID, err := strconv.ParseInt(shopId, 0, 0)
	// if err != nil {
	// 	fmt.Println("error while parsing")
	// }

	// shopDetails, _ := models.GetShopById(ID)
	// res, _ := json.Marshal(shopDetails)
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

// func Signup(c *gin.Context) {

// }

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
	// err = user.HashPassword(user.Password)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(500, gin.H{
	// 		"Error": "Error Hashing Password",
	// 	})
	// 	c.Abort()
	// 	return
	// }
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

	// tmpl := template.Must(template.ParseFiles("templates/base.gohtml"))
	// if r.Method != http.MethodPost {
	// 	tmpl.Execute(w, nil)
	// 	return
	// }

	// CreateShop := &models.Shop{}
	// utilis.ParseBody(r, CreateShop)
	// s := CreateShop.CreatShop()
	// res, _ := json.Marshal(s)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)

	// CreateDetails := &models.details{}
	// utilis.ParseBody(r, CreateDetails)
	// i := CreateDetails.CreateDetails()
	// req, _ := json.Marshal(i)
	// w.WriteHeader(http.StatusOK)
	// w.Write(req)

	// w, _ = template.ParseFiles("tutorial.html")
	// handle.RenderTemplate(w, "base.gohtml")

}

func DeleteShop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shopId := vars["shopId"]
	ID, err := strconv.ParseInt(shopId, 0, 0)
	if err != nil {
		fmt.Println("error whhile parsing Delete ")
	}

	shop := models.DeleteShop(ID)
	res, _ := json.Marshal(shop)

	w.Header().Set("content Type ", "pkglication/json")
	w.Write(res)
}

// func UpdateShop(w http.ResponseWriter, r *http.Request) {
// 	var updateShop = &models.Shop{}
// 	utilis.ParseBody(r, updateShop)
// 	vars := mux.Vars(r)
// 	shopId := vars["shopId"]
// 	ID, err := strconv.ParseInt(shopId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error whhile parsing Update ")
// 	}
// 	shopDetails, db := models.GetShopById(ID)
// 	if updateShop.ProductName != "" {
// 		shopDetails.ProductName = updateShop.ProductName

// 	}
// 	if updateShop.Category != "" {
// 		shopDetails.Category = updateShop.Category

// 	}
// 	if updateShop.Quantity != "" {
// 		shopDetails.Quantity = updateShop.Quantity
// 	}
// 	db.Save(&shopDetails)

// 	res, _ := json.Marshal(shopDetails)
// 	w.Header().Set("content Type ", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

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

//User Detais from shop

func CreateDetails(w http.ResponseWriter, r *http.Request) {
	CreateDetails := &models.Users{}
	utilis.ParseBody(r, CreateDetails)
	s := CreateDetails.CreateDetails()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllDetails(w http.ResponseWriter, r *http.Request) {
	newDetails := models.GetAllDetails()
	req, _ := json.Marshal(newDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}

func GetDetailsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DetailsId := vars["detailsId"]
	ID, err := strconv.ParseInt(DetailsId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing Id")
	}

	Details, _ := models.GetDetailsById(ID)
	req, _ := json.Marshal(Details)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}

func DeleteDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detailsId := vars["detailsId"]
	ID, err := strconv.ParseInt(detailsId, 0, 0)
	if err != nil {
		fmt.Println("error whhile parsing Delete ")
	}

	details := models.DeleteShop(ID)
	req, _ := json.Marshal(details)

	w.Header().Set("content Type ", "pkglication/json")
	w.Write(req)
}

func UpdateDetails(w http.ResponseWriter, r *http.Request) {
	var updateDetails = &models.Users{}
	utilis.ParseBody(r, updateDetails)
	vars := mux.Vars(r)
	detailsId := vars["shopId"]
	ID, err := strconv.ParseInt(detailsId, 0, 0)
	if err != nil {
		fmt.Println("error whhile parsing Update ")
	}
	detailsUpdate, db := models.GetDetailsById(ID)
	if updateDetails.Details != "" {
		detailsUpdate.Details = updateDetails.Details

	}
	if updateDetails.From != "" {
		detailsUpdate.From = updateDetails.From

	}
	if updateDetails.Country != "" {
		detailsUpdate.Country = updateDetails.Country
	}
	if updateDetails.ImportedCompany != "" {
		detailsUpdate.Country = updateDetails.ImportedCompany
	}
	db.Save(&detailsUpdate)

	req, _ := json.Marshal(detailsUpdate)
	w.Header().Set("content Type ", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}
