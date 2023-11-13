package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	host := "localhost"
	username := "postgres"
	password := "mmcloud"
	databaseName := "gin_api_2"
	port := "5433"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&Book{})

	//DB = database
}
