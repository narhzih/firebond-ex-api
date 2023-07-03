package firebondmongo

import (
	"context"
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/db/repository"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// GetFiatRateRecordForSymbol checks if a fiat rate exists on a symbol. If it does, it returns it, else, it throws an error
func (act rateActions) GetFiatRateRecordForSymbol(symbol, fiatSymbol string) (models.Rate, error) {
	rate, err := act.GetCryptoRatesBySymbol(symbol)
	if err != nil {
		return models.Rate{}, err
	}
	for key, _ := range rate.FiatPrices {
		if key == fiatSymbol {
			return rate, nil
		}
	}

	return models.Rate{}, ErrFiatRateToSymbolNotFound
}

func (act rateActions) GetAllRates() ([]models.Rate, error) {
	var rates []models.Rate
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cursor, err := act.Collection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Rate{}, ErrNoDocuments
		}
	}
	if err = cursor.All(ctx, &rates); err != nil {
		return []models.Rate{}, err
	}
	return rates, nil
}

// UpSert Tries to create an exchange rate record or update it if  it already exists
func (act rateActions) UpSert(data models.Rate) error {
	rate := models.Rate{
		Symbol:     data.Symbol,
		FiatPrices: data.FiatPrices,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	_, err := act.Collection.UpdateOne(ctx, bson.M{"symbol": rate.Symbol}, bson.M{"$set": bson.M{"symbol": rate.Symbol, "fiatPrices": rate.FiatPrices}}, opts)
	if err != nil {
		return err
	}

	return nil
}
