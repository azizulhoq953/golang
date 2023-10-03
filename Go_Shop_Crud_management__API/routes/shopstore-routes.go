package routes

import (
	"crud/shopDB/controllers"

	"github.com/gorilla/mux"
)

var RegisterShopStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/shop", controllers.CreateShop).Methods("POST")
	router.HandleFunc("/shop", controllers.GetShop).Methods("GET")
	router.HandleFunc("/shop/{shopId}", controllers.GetShopById).Methods("GET")
	router.HandleFunc("/shop/{shopId}", controllers.UpdateShop).Methods("PUT")
	router.HandleFunc("/shop/{shopId}", controllers.DeleteShop).Methods("DELETE")
}
