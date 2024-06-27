package models

import (
	"gorm.io/gorm"
	"phonebook/internal/modules/user/models"
)

type Phone struct {
	gorm.Model
	Number string      `gorm:"unique" json:"phone"`
	User   models.User `json:"user"`
	UserId int         `json:"user_id"`
}
