package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstname"`
	LastName  string             `bson:"lastname"`
	Email     string             `bson:"email"`
}
