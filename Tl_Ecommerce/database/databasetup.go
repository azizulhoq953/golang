package database

// import (
// 	// "github.com/jinzhu/gorm"
// 	"fmt"
// 	"log"

// 	"github.com/joho/godotenv"
// 	// _ "github.com/lib/pq"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	// _ "github.com/lib/pq"
// )

// var (
// 	GlobalDB *gorm.DB
// )

// func Connect() (err error) {
// 	// Read the environment variables from the .env file
// 	config, err := godotenv.Read()
// 	if err != nil {
// 		log.Fatal("Error reading .env file")
// 	}
// 	print("this is config", config)
// 	// Create the data source name (DSN) using the environment variables
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",

// 		config["host"],
// 		config["port"],
// 		config["user"],
// 		config["password"],
// 		config["dbname"],
// 		// config["DB_HOST"],
// 		// config["DB_PORT"],
// 		// config["DB_USERNAME"],
// 		// config["DB_PASSWORD"],
// 		// config["DB_NAME"],
// 	)
// 	// Create the connection and store it in the GlobalDB variable
// 	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return
// 	}
// 	print(dsn)
// 	return
// }

// //
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// 	"gorm.io/driver/postgres"
	// 	"gorm.io/gorm"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://azizulhoq953:azizulhoq953@cluster0.ajm4cxn.mongodb.net/?retryWrites=true&w=majority")) //mongodb://development:testpassword@localhost:27017
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}
	fmt.Println("Successfully Connected to the mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return collection

}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}

func AddressData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}

func OrderData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}

func PaymentData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}
