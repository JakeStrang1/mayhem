package user

import (
	"time"

	"example.com/myNetwork/models"
)

type User struct {
	ID        string    `json:"id" sign_up:"isdefault"`
	Email     string    `json:"email" sign_up:"required,email"`
	CreatedAt time.Time `json:"createdAt" sign_up:"isdefault"`
}

func FromDBUser(dbUser models.User) User {
	return User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
	}
}
