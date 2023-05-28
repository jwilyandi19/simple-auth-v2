package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Username  string             `bson:"username" json:"username" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	IsAdmin   bool               `bson:"is_admin" json:"is_admin" validate:"required"`
}
