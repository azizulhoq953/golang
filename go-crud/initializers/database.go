package initializers

import (
	"Go_3-Month/golang/go-crud/models"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open("host=localhost user=postgres password=new_password dbname=crud_DB port=5432 sslmode=disable"))

	if err != nil {
		log.Fatal("Failed DataBase Conncet")
	}
	DB.AutoMigrate(&models.Post{})

}
