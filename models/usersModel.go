package models

import (
	"context"
	"fmt"
	"time"

	"github.com/TirthShingala/the-user/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	var indexModel = mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
	}

	var name, err = UserCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println("Index creation error", err)
	}

	fmt.Println("Name of Index Created: " + name)
}

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	FirstName     string             `json:"firstName" validate:"required,min=2,max=100"`
	LastName      string             `json:"lastName" validate:"required,min=2,max=100"`
	Password      string             `json:"password" validate:"required,min=6"`
	Email         string             `json:"email" validate:"required,email"`
	Phone         string             `json:"phone" validate:"required"`
	Role          string             `json:"role" validate:"required"`
	IsBlocked     bool               `json:"isBlocked"`
	ProfilePicUrl string             `json:"profilePicUrl"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
