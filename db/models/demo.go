package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Demo struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	FullName string             `json:"fullName" bson:"fullName"`
}
