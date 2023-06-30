package firebondmongo

import (
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type rateActions struct {
	Collection *mongo.Collection
	Logger     zerolog.Logger
}

func NewRateActions(mongoDatabase *mongo.Database, logger zerolog.Logger) repository.RateRepository {
	collection := mongoDatabase.Collection("crypto")
	return rateActions{Collection: collection, Logger: logger}
}

func (act rateActions) CreateCryptoToFiatRateData(data models.Rate) (models.Rate, error) {
	// create a new exchange data

	return models.Rate{}, nil
}
