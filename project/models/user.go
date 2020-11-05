package models

import "time"

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"createdAt"`
}
