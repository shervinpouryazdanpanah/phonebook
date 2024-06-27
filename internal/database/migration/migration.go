package migration

import (
	"fmt"
	phoneModel "phonebook/internal/modules/phone/models"
	userModel "phonebook/internal/modules/user/models"
	"phonebook/pkg/database"
)

func Migration() {
	db := database.Connection()
	err := db.AutoMigrate(&userModel.User{}, &phoneModel.Phone{})
	if err != nil {
		panic(err)
	}
	fmt.Println("migration success")
}
