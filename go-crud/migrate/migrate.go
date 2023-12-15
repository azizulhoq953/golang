package migrate

import (
	"Go_3-Month/golang/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// func main() {
// 	initializers.DB.AutoMigrate(&models.Post{})
// }
