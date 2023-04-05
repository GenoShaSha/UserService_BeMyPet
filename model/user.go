package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"firstName,omitempty"`
	LastName    string             `bson:"lastName,omitempty"`
	DateOfBirth string             `bson:"dateOfBirth,omitempty"`
	Email       string             `bson:"email,omitempty"`
}
