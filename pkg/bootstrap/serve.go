package bootstrap

import (
	"phonebook/pkg/config"
	"phonebook/pkg/database"
	"phonebook/pkg/routing"
)

func Server() {

	//set config
	config.Set()

	//connect database
	database.Connect()

	routing.Set()
}
