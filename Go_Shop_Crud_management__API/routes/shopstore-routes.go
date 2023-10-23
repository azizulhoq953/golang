package routes

// func setupRouter() *gin.Engine {
// 	// Create a new router
// 	r := gin.Default()
// 	// Add a welcome route
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(200, "Welcome To This Website")
// 	})
// 	// Create a new group for the API
// 	api := r.Group("/api")
// 	{
// 		// Create a new group for the public routes
// 		public := api.Group("/public")
// 		{
// 			// Add the login route
// 			public.POST("/login", controllers.Login)
// 			// Add the signup route
// 			public.POST("/signup", controllers.Signup)
// 		}
// 		// Add the signup route
// 		protected := api.Group("/protected").Use(middleware.Authz())
// 		{
// 			// Add the profile route
// 			protected.GET("/profile", controllers.Profile)
// 		}
// 	}
// 	// Return the router
// 	return r
// }

// var RegisterShopStoreRoutes = func(router *mux.Router) {
// 	// setupRouter()
// 	router.HandleFunc("/shop", controllers.CreateShop).Methods("POST")
// 	router.HandleFunc("/shop", controllers.GetShop).Methods("GET")
// 	router.HandleFunc("/shop/{shopId}", controllers.GetShopById).Methods("GET")
// 	router.HandleFunc("/shop/{shopId}", controllers.UpdateShop).Methods("PUT")
// 	router.HandleFunc("/shop/{shopId}", controllers.DeleteShop).Methods("DELETE")
// 	router.HandleFunc("/page", handle.ServeHTML).Methods("GET")
// 	//details pages
// 	router.HandleFunc("/detailse", controllers.CreateDetails).Methods("POST")
// 	router.HandleFunc("/detailse", controllers.GetAllDetails).Methods("GET")
// 	router.HandleFunc("/detailse/{detailsId}", controllers.GetDetailsById).Methods("GET")
// 	router.HandleFunc("/detailse/{detailsId}", controllers.UpdateDetails).Methods("PUT")
// 	router.HandleFunc("/detailse/{detailsId}", controllers.DeleteDetails).Methods("DELETE")
// 	router.HandleFunc("/page", handle.ServeHTML).Methods("GET")

// 	//RegistrationForm
// 	// router.HandleFunc("/user", controllers.CreatRegistration).Methods("POST")
// 	// router.HandleFunc("/userall", controllers.GetAllRegistration).Methods("GET")
// 	// router.HandleFunc("/user/login", controllers.Login).Methods("GET")
// }
