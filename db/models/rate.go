package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rate struct {
	ID         primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Symbol     string                 `json:"symbol" bson:"symbol"`
	FiatPrices map[string]interface{} `json:"fiatPrices" bson:"fiatPrices"`
}
