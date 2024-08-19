package database

import (
	"fmt"
	"log"
	"supplysense/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB{
	config.LoadConfig()
	dbconfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
	config.GetEnv("DATABASE_HOST"),
	config.GetEnv("DATABASE_USERNAME"),
	config.GetEnv("DATABASE_PASSWORD"),
	config.GetEnv("DATABASE_NAME"),
	config.GetEnv("DATABASE_PORT"),	
	)
	db, err := gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if(err != nil){
		log.Fatalf("Error: Connect to database failed")
	}
	fmt.Print("⇨ Successfully connected to database")
	return db
}