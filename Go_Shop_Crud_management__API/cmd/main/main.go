package main

import (
	"crud/shopDB/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterShopStoreRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)

}
