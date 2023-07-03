package repository

import "firebond-ex-api.com/db/models"

type RateRepository interface {
	GetAllRates() ([]models.Rate, error)
	GetCryptoRatesBySymbol(symbol interface{}) (models.Rate, error)
	GetFiatRateRecordForSymbol(symbol, fiatSymbol string) (models.Rate, error)
	UpSert(data models.Rate) error
}
