package migrations

import (
	"fmt"
	"supplysense/internal/User/model"
	"gorm.io/gorm"
)

func Up(db *gorm.DB){
	err := db.AutoMigrate(&model.User{})
	if err != nil{
		fmt.Errorf("Failed to migrate: %v", err)
	}
	fmt.Println("Migrate 'users' Success")
}

func Down(db *gorm.DB){
	if !db.Migrator().HasTable("users") {
		fmt.Println("Table 'users' does not exist")
		return
	}
	
	err := db.Migrator().DropTable("users")
	if err != nil{
		fmt.Errorf("Failed to migrate: %v", err)
	}	
	fmt.Println("Table 'users' successfully deleted")
}