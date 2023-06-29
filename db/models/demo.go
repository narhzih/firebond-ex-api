package models

type Demo struct {
	ID       string `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	FullName string `json:"fullName" bson:"fullName"`
}
