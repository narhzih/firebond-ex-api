package firebondmongo

import (
	"context"
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	ErrFiatRateToSymbolNotFound = fmt.Errorf("could not find the specified fiat rate for symbol")
	ErrNoDocuments              = fmt.Errorf("no documents found")
)

type rateActions struct {
	Collection *mongo.Collection
	Logger     zerolog.Logger
}

func NewRateActions(mongoDatabase *mongo.Database, logger zerolog.Logger) repository.RateRepository {
	collection := mongoDatabase.Collection("rate")
	return rateActions{Collection: collection, Logger: logger}
}

// GetCryptoRatesBySymbol fetches the rates of a crypto currency against any fiat
func (act rateActions) GetCryptoRatesBySymbol(symbol interface{}) (models.Rate, error) {
	var rate models.Rate
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	data := act.Collection.FindOne(ctx, bson.M{"symbol": symbol})
	err := data.Decode(&rate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Rate{}, ErrNoDocuments
		}
		return models.Rate{}, err
	}

	return rate, nil
}

// CreateCryptoToFiatRateData creates the crypto-fiat rate data of a crypto
func (act rateActions) CreateCryptoToFiatRateData(data models.Rate) (models.Rate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := act.Collection.InsertOne(ctx, data)
	if err != nil {
		return models.Rate{}, err
	}
	rate, err := act.GetCryptoRatesBySymbol(data.Symbol)
	if err != nil {
		return models.Rate{}, err
	}
	return rate, nil
}

func (act rateActions) CreateFiatRateRecordForSymbol(symbol, fiatSymbol, fiatValue string) (models.Rate, error) {
	//1. Check if this fiat is already in the fiat list

	rate, err := act.GetCryptoRatesBySymbol(symbol)
	if err != nil {
		return models.Rate{}, err
	}
	//2. If it's already there, just return it
	for key, _ := range rate.FiatPrices {
		if key == fiatSymbol {
			return rate, nil
		}
	}
	rate.FiatPrices[fiatSymbol] = fiatValue
	// update the value in db
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = act.Collection.UpdateOne(ctx, bson.M{"symbol": symbol}, bson.M{"fiatPrices": rate.FiatPrices})
	if err != nil {
		return models.Rate{}, nil
	}
	// At this point, we know it's not there
	return rate, nil
}

func (act rateActions) GetFiatRateRecordForSymbol(symbol, fiatSymbol string) (models.Rate, error) {

	rate, err := act.GetCryptoRatesBySymbol(symbol)
	if err != nil {
		return models.Rate{}, err
	}
	for key, _ := range rate.FiatPrices {
		if key == fiatSymbol {
			act.Logger.Info().Msg("found the rate we're looking for ")
			return rate, nil
		}
	}

	return models.Rate{}, ErrFiatRateToSymbolNotFound
}
