package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password        *string            `json:"password"   validate:"required,min=6"`
	Email           *string            `json:"email"      validate:"email,required"`
	Phone           *string            `json:"phone"      validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint64            `json:"price" bson:"price"`
	Discount     *float64           `json:"discount" bson:"discount"`
	Category     *string            `json:"category" bson:"category"`
	Rating       *uint8             `json:"rating" bson:"rating"`
	Image        *string            `json:"image"`
}

func (p *Product) CalculateDiscountedPrice() *float64 {
	if p.Price == nil || p.Discount == nil {
		return nil // Cannot calculate discount without price or discount information
	}

	price := float64(*p.Price)
	discount := float64(*p.Discount)

	discountedPrice := price - (price * discount / 100.0)
	return &discountedPrice
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price"  bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image"  bson:"image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Orderered_At   time.Time          `json:"ordered_on"  bson:"ordered_on"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}

// func init() {
// 	database.Connect()
// 	// db = config.GetDB()
// 	// template.Must(template.New("templates").Parse("html"))

// 	database.GlobalDB.AutoMigrate(&User{})
// 	// config.GlobalDB.AutoMigrate(&Users{})
// 	database.GlobalDB.AutoMigrate(&Product{})
// 	database.GlobalDB.AutoMigrate(&ProductUser{})
// 	database.GlobalDB.AutoMigrate(&Address{})
// 	database.GlobalDB.AutoMigrate(&Order{})

// 	// db.AutoMigrate(&Shop{})
// 	// db.AutoMigrate(&Users{})
// 	// db.AutoMigrate(&RegistrationForm{})

// }

// func (Regf *User) CreateUserRecord() error {
// 	result := database.GlobalDB.Create(&Regf)

// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil

// }

// // HashPassword encrypts user password
// // HashPassword takes a string as a parameter and encrypts it using bcrypt
// // It returns an error if there is an issue encrypting the password
// func (userHas *User) HashPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return err
// 	}
// 	userHas.Password = string(bytes)
// 	return nil
// }

// // CheckPassword checks user password
// // CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// // It returns an error if there is an issue comparing the passwords
// func (userHas *User) CheckPassword(providedPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetAllRegistration() []RegistrationForm {
// 	var RegistrationForms []RegistrationForm
// 	database.GlobalDB.Find(&RegistrationForms)
// 	return RegistrationForms
// }

// func Login(Email string) (*RegistrationForm, *gorm.DB) {
// 	var getemaiPass RegistrationForm
// 	GlobalDB := GlobalDB.Where("EMAIL=?", Email).Find(&getemaiPass)
// 	return &getemaiPass, GlobalDB
// }

// func (s *Shop) CreatShop() error {
// 	result := database.GlobalDB.Create(&s)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil

// }
