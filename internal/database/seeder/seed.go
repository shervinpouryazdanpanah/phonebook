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
		Name:     "admin",
		Password: string(hashedPassword),
	}

	db.Create(&user)
	log.Print("Successfully created user")

	phone := phoneModel.Phone{
		User:   user,
		Number: "09381584207",
	}

	db.Create(&phone)
	log.Print("Successfully created phone")

	phone2 := phoneModel.Phone{
		User:   user,
		Number: "09102587628",
	}

	db.Create(&phone2)
	log.Print("Successfully created phone2")

	phone3 := phoneModel.Phone{
		User:   user,
		Number: "09388405367",
	}

	db.Create(&phone3)
	log.Print("Successfully created phone2")

}
