package main

import (
	"firebond-ex-api.com/cmd/api/helpers"
	"firebond-ex-api.com/cmd/api/internal"
)

func periodicallyUpdateDatabase(app internal.Application) error {
	//1. Get all the supported coin pairs
	supportedRates, err := app.Services.CC.GetSupportedCryptoToFiatPairsForBinance()
	if err != nil {
		return err
	}
	//2. Transform them into rateModel format
	transformedSupportedRates := helpers.TransformExchangeApiResponseDataToRateModel(supportedRates)

	//3. For each symbol, get rate data and update each symbol
	for _, rate := range transformedSupportedRates {
		fsym := rate.Symbol
		var tsyms string
		for fiat, _ := range rate.FiatPrices {
			tsyms += fiat + ","
		}
		symbolToFiatRates, err := app.Services.CC.GetRatesForFsymsAndTsyms(fsym, tsyms)
		if err != nil {
			return err
		}
		rate.FiatPrices = symbolToFiatRates
		// now create new or update if it doesn't exist
		_ = app.Repositories.Rate.UpSert(rate)
	}
	return nil
}
