package repository

import "firebond-ex-api.com/db/models"

type RateRepository interface {
	GetCryptoRatesBySymbol(symbol interface{}) (models.Rate, error)
	CreateCryptoToFiatRateData(data models.Rate) (models.Rate, error)
}
