package repository

import "firebond-ex-api.com/db/models"

type RateRepository interface {
	GetCryptoRatesBySymbol(symbol interface{}) (models.Rate, error)
	CreateCryptoToFiatRateData(data models.Rate) (models.Rate, error)
	CreateFiatRateRecordForSymbol(symbol, fiatSymbol string, fiatValue float64) (models.Rate, error)
	GetFiatRateRecordForSymbol(symbol, fiatSymbol string) (models.Rate, error)
	UpSert(data models.Rate) error
}
