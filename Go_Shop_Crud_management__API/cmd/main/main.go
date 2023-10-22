package main

import (
	"crud/shopDB/controllers"
	"crud/shopDB/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Application Is Running")
	// r := mux.NewRouter()
	// routes.RegisterShopStoreRoutes(r)
	r := setupRouter()

	r.Run(":8080")
	// http.Handle("/", n)
	// http.ListenAndServe("localhost:8080", r)

}

func setupRouter() *gin.Engine {
	// Create a new router
	n := gin.Default()
	// Add a welcome route
	n.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
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
			public.POST("/signup", controllers.Signup)
			// add the shop and Details Router
			public.POST("/shop", controllers.CreateShop)
			//ShopById
			public.GET("/shop/:id", controllers.GetShopById)
			public.GET("/all", controllers.GetShop)
			public.PATCH("/shop/:id", controllers.UpdateShop)
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

// setupRouter()
// router.HandleFunc("/shop", controllers.CreateShop).Methods("POST")
// router.HandleFunc("/shop", controllers.GetShop).Methods("GET")
// router.HandleFunc("/shop/{shopId}", controllers.GetShopById).Methods("GET")
// router.HandleFunc("/shop/{shopId}", controllers.UpdateShop).Methods("PUT")
// router.HandleFunc("/shop/{shopId}", controllers.DeleteShop).Methods("DELETE")
// router.HandleFunc("/page", handle.ServeHTML).Methods("GET")
// //details pages
// router.HandleFunc("/detailse", controllers.CreateDetails).Methods("POST")
// router.HandleFunc("/detailse", controllers.GetAllDetails).Methods("GET")
// router.HandleFunc("/detailse/{detailsId}", controllers.GetDetailsById).Methods("GET")
// router.HandleFunc("/detailse/{detailsId}", controllers.UpdateDetails).Methods("PUT")
// router.HandleFunc("/detailse/{detailsId}", controllers.DeleteDetails).Methods("DELETE")
// router.HandleFunc("/page", handle.ServeHTML).Methods("GET")
