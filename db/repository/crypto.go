package repository

import "firebond-ex-api.com/db/models"

type CryptoRepository interface {
	CreateCryptoExchangeData(data models.Crypto) (models.Crypto, error)
}
