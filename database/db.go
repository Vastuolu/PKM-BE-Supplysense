package database

import (
	"fmt"
	"log"
	"supplysense/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	config.LoadEnv()
	dbconfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
	config.GetEnv("DATABASE_HOST"),
	config.GetEnv("DATABASE_USERNAME"),
	config.GetEnv("DATABASE_PASSWORD"),
	config.GetEnv("DATABASE_NAME"),
	config.GetEnv("DATABASE_PORT"),	
	)
	DB, err = gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if(err != nil){
		log.Fatalf("Error: Connect to database failed")
	}
	fmt.Println("⇨ Successfully connected to database")
}