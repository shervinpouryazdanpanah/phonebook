package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	phoneModel "phonebook/internal/modules/phone/models"
	userModel "phonebook/internal/modules/user/models"
	"phonebook/pkg/database"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{
		Username: "admin",
		Password: string(hashedPassword),
	}

	phone := phoneModel.Phone{
		User:  user,
		Phone: "09381584207",
	}

	db.Create(&user)
	log.Print("Successfully created user")
	db.Create(&phone)
	log.Print("Successfully created phone")

}
