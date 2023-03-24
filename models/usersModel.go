package models

import (
	"time"

	"github.com/TirthShingala/the-user/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName" validate:"required,min=2,max=100"`
	LastName  string             `json:"lastName" validate:"required,min=2,max=100"`
	Password  string             `json:"password" validate:"required,min=6"`
	Email     string             `json:"email" validate:"required,email"`
	Phone     string             `json:"phone" validate:"required"`
	Role      string             `json:"role" validate:"required"`
	IsBlocked bool               `json:"isBlocked"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
