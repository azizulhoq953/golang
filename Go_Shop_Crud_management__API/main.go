package main

import (
	routes "crud/shopDB/routers"
)

func main() {
	println("Application Is Running")
	// r := mux.NewRouter()
	// routes.RegisterShopStoreRoutes(r)

	// r := gin.Default()
	// r.LoadHTMLFiles("templates/index.html")

	// Shops := make([]models.Shop, 0)
	// Shops = append(Shops, models.Shop{
	// 	ProductName: "ProductName 1",
	// 	Category:    "Category 1",
	// 	Aviliable:   "Aviliable 1",
	// 	Quantity:    "Quantity 1",
	// })
	// Shops = append(Shops, models.Shop{
	// 	ProductName: "ProductName 1",
	// 	Category:    "Category 1",
	// 	Aviliable:   "Aviliable 1",
	// 	Quantity:    "Quantity 1",
	// })

	// r.GET("/shops", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"Shops": Shops,
	// 	})
	// })

	// r := setupRouter()
	n := routes.RouteSetup()

	n.Run(":8080")
	// http.Handle("/", n)
	// http.ListenAndServe("localhost:8080", r)

}
