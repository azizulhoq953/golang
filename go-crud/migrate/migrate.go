package migrate

import (
	"Go_3-Month/golang/go-crud/initializers"
	"Go_3-Month/golang/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
