package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// postgres dialect from gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var db *gorm.DB

// User table
type User struct {
	gorm.Model
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname, omitempty"`
	Email     string `json:"email, omitempty"`
	Password  string `json:"password, omitempty"`

	HelpRequestHelpers []HelpRequestHelper `gorm:"foreignkey:UserID"`
	Incidents          []Incident          `gorm:"foreignkey:UserID"`
}

// Incident table
type Incident struct {
	gorm.Model
	Longitude   float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
	Description string  `json:"description"`
	UserID      uint    `json:"user_id"` // pointer back to the User table
}

// HelpRequest table
type HelpRequest struct {
	gorm.Model
	Longitude          float32             `json:"longitude"`
	Latitude           float32             `json:"latitude"`
	Description        string              `json:"description"`
	Status             string              `json:"status"` // string: open/ closed/ in-progress
	Tags               pq.StringArray      `gorm:"type:varchar(64)[]"`
	HelpRequestHelpers []HelpRequestHelper `gorm:"foreignkey:UserID"`
}

// HelpRequestHelper table
type HelpRequestHelper struct {
	gorm.Model
	UserID        uint   `json:"user_id"` // pointer back to the User table
	HelpRequestID uint   `json:"help_request_id"`
	Description   string `json:"description"`
}

func init() {
	fmt.Println("init in model.go")

	_ = godotenv.Load()

	host := os.Getenv("HOST")
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTableIfExists(&User{}, &Incident{}, &HelpRequest{}, &HelpRequestHelper{})
	// Migrate the schema
	db.AutoMigrate(&User{}, &Incident{}, &HelpRequest{}, &HelpRequestHelper{})
	db.Create(&User{Firstname: "Diep", Lastname: "Truong", Email: "ngocdieptruong11@gmail.com", Password: "123"})
}

func TestFunction() {
	fmt.Println("running testFunction in model")
}

// func main() {

// 	_ = godotenv.Load()

// 	host := os.Getenv("HOST")
// 	user := os.Getenv("DBUSER")
// 	dbname := os.Getenv("DBNAME")
// 	password := os.Getenv("PASSWORD")
// 	port := 5432

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := gorm.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})

// 	// Create
// 	db.Create(&Product{Code: "L1212", Price: 1000})

// 	// Read
// 	var product Product
// 	db.First(&product, 1)                   // find product with id 1
// 	db.First(&product, "code = ?", "L1212") // find product with code l1212

// 	// Update - update product's price to 2000
// 	db.Model(&product).Update("Price", 2000)

// 	// Delete - delete product
// 	db.Delete(&product)
// }
