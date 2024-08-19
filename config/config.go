package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadConfig(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error Load Env %v", err)
	}
}

func GetEnv(key string) string {
	var val string = os.Getenv(key)
	if (val == "") {
		log.Fatalf("Env Key not found")
	} 
	return val
}