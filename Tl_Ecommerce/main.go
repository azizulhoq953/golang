package main

import (
	"azizulhoq953/golang/Tl_Ecommerce/controllers"
	"azizulhoq953/golang/Tl_Ecommerce/database"
	"azizulhoq953/golang/Tl_Ecommerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	productData := database.ProductData(database.Client, "Products")
	userData := database.UserData(database.Client, "Users")
	// addressData := database.AddressData(database.Client, "Addresses")
	// orderData := database.OrderData(database.Client, "Orders")
	// paymentData := database.PaymentData(database.Client, "Payments")

	// app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	app := controllers.NewApplication(productData, userData)
	// discountedPrice := models.Product(CalculateDiscountedPrice())
	//, , paymentData,
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}
