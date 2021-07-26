package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Data string             `bson:"data" json:"data"`
}
