package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Books struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string	`json:"title" bson:"title"`
	Price int `json:"price" bson:"price"`
	Category string `json:"category" bson:"category"`
}

type UpdateBody struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string	`json:"title" bson:"title"`
	Price int `json:"price" bson:"price"` // value that has to be modified
}