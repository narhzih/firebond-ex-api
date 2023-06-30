package firebondmongo

import "firebond-ex-api.com/db/models"

var createCryptoToFiatRateDataTestCases = map[string]struct {
	inputRate models.Rate
	wantRate  models.Rate
	wantErr   error
}{
	"successfully created crypto to fiat rate data": {
		inputRate: models.Rate{
			Symbol: "HGH",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 31000.5,
			},
		},
		wantRate: models.Rate{
			Symbol: "HGH",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 31000.5,
			},
		},
		wantErr: nil,
	},
}

var getCryptoRatesBySymbolTestCases = map[string]struct {
	inputSymbol interface{}
	wantRate    models.Rate
	wantErr     error
}{
	"successfully fetched rates data by symbol": {
		inputSymbol: "BTC",
		wantRate: models.Rate{
			Symbol: "BTC",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			},
		},
		wantErr: nil,
	},
	"fails when a non-existent symbol isn't provided": {
		inputSymbol: "XXX",
		wantRate:    models.Rate{},
		wantErr:     ErrNoDocuments,
	},
}
