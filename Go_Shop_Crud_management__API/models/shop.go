package models

import (
	"crud/shopDB/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Shop struct {
	gorm.Model
	ProductName string `gorm:"" json:"productname"`
	Category    string `json:"category"`
	Aviliable   string `json:"Aviliable"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Shop{})
}

func (s *Shop) CreatShop() *Shop {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllShop() []Shop {
	var Shops []Shop
	db.Find(&Shops)
	return Shops
}

func GetShopById(Id int64) (*Shop, *gorm.DB) {
	var getShop Shop
	db := db.Where("ID=?", Id).Find(&getShop)
	return &getShop, db
}

func DeleteShop(ID int64) Shop {
	var shop Shop
	db.Where("ID=?", ID).Delete(shop)
	return shop
}
