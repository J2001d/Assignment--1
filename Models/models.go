package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID       primitive.ObjectID `json:"_id" bson:"id,omitempty"`
	Rating   string             `json:"rating" bson:"rating,omitempty"`
	Title    string             `json:"title" bson:"title,omitempty"`
	Director string             `json:"director" bson:"director,omitempty"`
}
