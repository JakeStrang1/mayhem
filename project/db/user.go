package db

import (
	"time"

	"example.com/myNetwork/models"
	"github.com/globalsign/mgo/bson"
)

type UserRepository struct {
	client     ClientInterface
	collection string
}

func NewUserRepository(client ClientInterface) *UserRepository {
	return &UserRepository{client: client, collection: "users"}
}

type ClientInterface interface {
	Create(string, interface{}, interface{}) error
	EnsureUniqueIndex(string, []string) error
}

func (r *UserRepository) Create(email string) (*models.User, error) {
	user := models.User{
		ID:        bson.NewObjectId().Hex(),
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := r.client.EnsureUniqueIndex(r.collection, []string{"email"})
	if err != nil {
		return nil, err
	}
	err = r.client.Create(r.collection, user, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
