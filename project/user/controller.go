package user

import "example.com/myNetwork/models"

type Controller struct {
	users UserRepository
}

func NewController(users UserRepository) *Controller {
	return &Controller{users: users}
}

type UserRepository interface {
	Create(string) (*models.User, error)
}

func (c *Controller) SignUp(email string) (*models.User, error) {
	return c.users.Create(email)
}
