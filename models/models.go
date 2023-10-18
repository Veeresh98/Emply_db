package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
}
