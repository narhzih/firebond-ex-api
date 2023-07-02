package helpers

import (
	"firebond-ex-api.com/db/models"
	"firebond-ex-api.com/services/cc"
)

func TransformExchangeApiResponseDataToRateModel(exApiRes cc.ExchangeApiResponse) []models.Rate {
	var rates []models.Rate
	pairs := exApiRes.Data.Data.Exchanges.Binance.Pairs
	for symbol, symbolData := range pairs {
		rateModel := models.Rate{
			Symbol:     symbol,
			FiatPrices: map[string]interface{}{},
		}
		for fiat, _ := range symbolData.Tsyms {
			// default everything to zero
			rateModel.FiatPrices[fiat] = 0
		}
		rates = append(rates, rateModel)
	}
	return rates
}
