package routes

import (
	"crud/shopDB/controllers"
	"crud/shopDB/middleware"
	"crud/shopDB/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine { //changes to function name and working func
	// Create a new router
	// n := gin.Default()
	// // Add a welcome route
	// n.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Welcome To This Website")
	// })

	// temp, err := template.ParseGlob("templates/*")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // Simply calling execute parses the first file in the directory
	// err = temp.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	n := gin.Default()
	n.LoadHTMLFiles("templates/index.html")

	Shops := make([]models.Shop, 0)
	Shops = append(Shops, models.Shop{
		ProductName: "ProductName 1",
		Category:    "Category 1",
		Aviliable:   "Aviliable 1",
		Quantity:    "Quantity 1",
	})
	Shops = append(Shops, models.Shop{
		ProductName: "ProductName 1",
		Category:    "Category 1",
		Aviliable:   "Aviliable 1",
		Quantity:    "Quantity 1",
	})

	n.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Shops": Shops,
		})
	})

	// Create a new group for the API
	api := n.Group("/api")
	{
		// Create a new group for the public routes
		public := api.Group("/public")
		{
			// Add the login route
			public.POST("/login", controllers.Login)
			// Add the signup route
			public.POST("/signup", controllers.Signup) //not it's a POST Method
			// add the shop and Details Router
			public.POST("/shop", controllers.CreateShop)
			//ShopById
			public.GET("/shop/:id", controllers.GetShopById)
			public.GET("/all", controllers.GetShop)
			public.PATCH("/shop/:id", controllers.UpdateShop)
			public.DELETE("/shop/:id", controllers.Deleted)
		}
		// Add the signup route
		protected := api.Group("/protected").Use(middleware.Authz())
		{
			// Add the profile route
			protected.GET("/profile", controllers.Profile)
		}
	}
	// Return the router
	return n

}
