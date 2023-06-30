package repository

import "firebond-ex-api.com/db/models"

type RateRepository interface {
	CreateCryptoToFiatRateData(data models.Rate) (models.Rate, error)
}
