package main

import routes "crud/shopDB/routers"

func main() {
	println("Application Is Running")
	// r := mux.NewRouter()
	// routes.RegisterShopStoreRoutes(r)

	// r := setupRouter()
	n := routes.RouteSetup()

	n.Run(":8080")
	// http.Handle("/", n)
	// http.ListenAndServe("localhost:8080", r)

}
