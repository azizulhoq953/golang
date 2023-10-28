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
	// config.GlobalDB.AutoMigrate(&Users{})
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

func (s *Shop) CreatShop() error {
	result := config.GlobalDB.Create(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
