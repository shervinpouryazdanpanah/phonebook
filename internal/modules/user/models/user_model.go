package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(255)" json:"name"`
	Username string `gorm:"unique;varchar(255)" json:"username"`
	Password string `gorm:"->:false;<-:create" json:"password"`
}
