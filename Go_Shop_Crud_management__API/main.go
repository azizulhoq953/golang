package main

import (
	routes "crud/shopDB/routers"
	"net/http"
)

func main() {
	println("Application Is Running")

	staticDir := "./templates"

	// Create a file server for the static directory
	fileServer := http.FileServer(http.Dir(staticDir))

	// Register the file server handler and serve on the root URL
	http.Handle("/", http.StripPrefix("/", fileServer))

	// Start the web server on port 8080
	// http.ListenAndServe(":8080", nil)

	n := routes.RouteSetup()

	n.Run(":8080")
	// http.Handle("/", n)
	// http.ListenAndServe("localhost:8080", r)

}
