package models

import (
	"github.com/azizulhoq953/jwtToken/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserRecord creates a user record in the database
// CreateUserRecord takes a pointer to a User struct and creates a user record in the database
// It returns an error if there is an issue creating the user record
func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil

}

// HashPassword encrypts user password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func (user *User) HashPassword(password string) error {
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
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
