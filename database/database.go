// package database

// import (
// 	"backend/config"
// 	// "backend/models"
// 	"fmt"
// 	"log"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func InitDB() {
// 	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 	// 	config.GetEnv("DB_HOST"),
// 	// 	config.GetEnv("DB_USER"),
// 	// 	config.GetEnv("DB_PASSWORD"),
// 	// 	config.GetEnv("DB_NAME"),
// 	// 	config.GetEnv("DB_PORT"),
// 	// )																														

// 	dbUrl := config.GetEnv("DB_URL")

// 	var err error
// 	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	fmt.Println("Connected to the database!")

	 
// 	// err = db.AutoMigrate(&models.User{})
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to auto-migrate User model: %v", err)
// 	// }
// }


package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	log.Println("DB connected....")
}	

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return DB
}