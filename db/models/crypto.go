package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FiatPrice struct {
	Price  float64 `json:"price" bson:"price"`
	Symbol string  `json:"symbol" bson:"symbol"`
}

type Crypto struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Price      float64            `json:"price" bson:"price"`
	Symbol     string             `json:"symbol" bson:"symbol"`
	FiatPrices []FiatPrice        `json:"fiatPrices" bson:"fiatPrices"`
}
