package config

import (
	// "github.com/jinzhu/gorm"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/lib/pq"
)

var (
	GlobalDB *gorm.DB
)

func Connect() (err error) {
	// Read the environment variables from the .env file
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	print("this is config", config)
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

// func Connect() {
// 	//"postgres", "host=localhost port=5432 user=postgres password=new_password dbname=owndb ")

// 	GlobalDB, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=new_password dbname=lin_database")
// 	if err != nil {
// 		panic(err)
// 	}
// 	GlobalDB = GlobalDB
// }

// func GetDB() *gorm.DB {
// 	return GlobalDB
// }
