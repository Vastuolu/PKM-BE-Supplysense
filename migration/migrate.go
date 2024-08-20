package main

import (
	"fmt"
	"supplysense/database"
	"supplysense/migration/migrations"
)

func main(){
	database.DBConnect()
	db := database.DB
	var choose string
	fmt.Println("1. Migrate \n2. Rollback")
	fmt.Scanln(&choose)
	if choose == "1"{
		migrations.Up(db)
	} else if choose == "2"{
		migrations.Down(db)
	}
}