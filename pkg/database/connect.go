package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"phonebook/pkg/config"
)

func Connect() {
	configuration := config.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.DB.Username,
		configuration.DB.Password,
		configuration.DB.Host,
		configuration.DB.Port,
		configuration.DB.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
