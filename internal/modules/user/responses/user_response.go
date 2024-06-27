package responses

import (
	userModel "phonebook/internal/modules/user/models"
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUser(user userModel.User) User {
	return User{
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}