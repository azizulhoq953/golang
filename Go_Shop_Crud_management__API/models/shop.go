package models

import (
	"crud/shopDB/pkg/config"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var GlobalDB *gorm.DB

// var db *gorm.DB

type Shop struct {
	gorm.Model
	ProductName string `gorm:"" json:"productname"`
	Category    string `json:"category"`
	Aviliable   string `json:"aviliable"`
	Quantity    string `json:"quantity"`
}

// Error implements error.
func (*Shop) Error() string {
	panic("unimplemented")
}

type RegistrationForm struct {
	gorm.Model
	// ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct { //mustCamelcase model name
	gorm.Model
	Details         string `gorm:"" json:"Details"`
	From            string `json:"From"`
	Country         string `json:"Country"`
	ImportedCompany string `json:"ImportedCompany"`
}

// var details = &Shop.details {
// 	// gorm.Model
// 	userName  string `gorm:"" json:"username"`
// 	Category  string `json:"category"`
// 	Aviliable string `json:"aviliable"`
// 	Quantity  string `json:"quantity"`
// }

func init() {
	config.Connect()
	// db = config.GetDB()
	// template.Must(template.New("templates").Parse("html"))

	config.GlobalDB.AutoMigrate(&Shop{})
	config.GlobalDB.AutoMigrate(&Users{})
	config.GlobalDB.AutoMigrate(&RegistrationForm{})

	// db.AutoMigrate(&Shop{})
	// db.AutoMigrate(&Users{})
	// db.AutoMigrate(&RegistrationForm{})

}

func (Regf *RegistrationForm) CreateUserRecord() error {
	result := config.GlobalDB.Create(&Regf)

	if result.Error != nil {
		return result.Error
	}
	return nil

}

// HashPassword encrypts user password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func (user *RegistrationForm) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func (user *RegistrationForm) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

//RegistrationForm
// func (r *RegistrationForm) CreatRegistration() *RegistrationForm {
// 	db.NewRecord(r)
// 	db.Create(&r)
// 	return r
// }

func GetAllRegistration() []RegistrationForm {
	var RegistrationForms []RegistrationForm
	config.GlobalDB.Find(&RegistrationForms)
	return RegistrationForms
}

func Login(Email string) (*RegistrationForm, *gorm.DB) {
	var getemaiPass RegistrationForm
	GlobalDB := GlobalDB.Where("EMAIL=?", Email).Find(&getemaiPass)
	return &getemaiPass, GlobalDB
}

// result := config.GlobalDB.Create(&Regf)

// if result.Error != nil {
// 	return result.Error
// }
// return nil

func (s *Shop) CreatShop() error {
	result := config.GlobalDB.Create(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil
	// 	// GlobalDB.NewRecord(s)
	// 	// GlobalDB.Create(&s)
	// return s
}

func DeleteShop(ID int64) Shop {
	var shop Shop
	GlobalDB.Where("ID=?", ID).Delete(shop)
	return shop
}

//user_products
func (i *Users) CreateDetails() *Users {
	GlobalDB.NewRecord(i)
	GlobalDB.Create(&i)
	return i
}

func GetAllDetails() []Users {
	var detailse []Users
	GlobalDB.Find(&detailse)
	return detailse
}

func GetDetailsById(Id int64) (*Users, *gorm.DB) {
	var getdetails Users
	GlobalDB := GlobalDB.Where("ID=?", Id).Find(&getdetails)
	return &getdetails, GlobalDB
}

func DeleteDetails(ID int64) Users {
	var detailse Users
	GlobalDB.Where("ID=?", ID).Delete(detailse)
	return detailse
}
