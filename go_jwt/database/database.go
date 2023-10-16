package database

import (
	"fmt"
	"log"

	// "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

// GlobalDB is a global db object that will be used across different packages
// var GlobalDB *gorm.DB

var (
	// db *gorm.DB
	GlobalDB *gorm.DB
)

// func Connect() {
// 	//"postgres", "host=localhost port=5432 user=postgres password=new_password dbname=owndb ")

// 	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=new_password dbname=jwt_db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	GlobalDB = d
// }

// func InitDatabase() *gorm.DB {
// 	return GlobalDB
// }

// InitDatabase creates a mysql db connection and stores it in the GlobalDB variable
// It reads the environment variables from the .env file and uses them to create the connection
// It returns an error if the connection fails
func InitDatabase() (err error) {
	// Read the environment variables from the .env file
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	// Create the data source name (DSN) using the environment variables
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",

		config["host"],
		config["port"],
		config["user"],
		config["password"],
		config["dbname"],
		// config["DB_HOST"],
		// config["DB_PORT"],
		// config["DB_USERNAME"],
		// config["DB_PASSWORD"],
		// config["DB_NAME"],
	)
	// Create the connection and store it in the GlobalDB variable
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	print(dsn)
	return
}
