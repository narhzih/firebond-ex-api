package firebondMongo

import (
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type cryptoActions struct {
	Collection *mongo.Collection
	Logger     zerolog.Logger
}

func NewCryptoActions(mongoDatabase *mongo.Database, logger zerolog.Logger) repository.CryptoRepository {
	collection := mongoDatabase.Collection("crypto")
	return cryptoActions{Collection: collection, Logger: logger}
}

func (act cryptoActions) CreateCryptoExchangeData(data models.Crypto) (models.Crypto, error) {
	// create a new exchange data
	
	return models.Crypto{}, nil
}
