package controllers

import (
	"crud/shopDB/models"
	"crud/shopDB/utilis"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewShop models.Shop

func GetShop(w http.ResponseWriter, r *http.Request) {
	newShop := models.GetAllShop()
	res, _ := json.Marshal(newShop)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetShopById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shopId := vars["shopId"]
	ID, err := strconv.ParseInt(shopId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	shopDetails, _ := models.GetShopById(ID)
	res, _ := json.Marshal(shopDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateShop(w http.ResponseWriter, r *http.Request) {
	CreateShop := &models.Shop{}
	utilis.ParseBody(r, CreateShop)
	s := CreateShop.CreatShop()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// w, _ = template.ParseFiles("tutorial.html")

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

func UpdateShop(w http.ResponseWriter, r *http.Request) {
	var updateShop = &models.Shop{}
	utilis.ParseBody(r, updateShop)
	vars := mux.Vars(r)
	shopId := vars["shopId"]
	ID, err := strconv.ParseInt(shopId, 0, 0)
	if err != nil {
		fmt.Println("error whhile parsing Update ")
	}
	shopDetails, db := models.GetShopById(ID)
	if updateShop.ProductName != "" {
		shopDetails.ProductName = updateShop.ProductName

	}
	if updateShop.Category != "" {
		shopDetails.Category = updateShop.Category

	}
	if updateShop.Quantity != "" {
		shopDetails.Quantity = updateShop.Quantity
	}
	db.Save(&shopDetails)

	res, _ := json.Marshal(shopDetails)
	w.Header().Set("content Type ", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
