package firebondmongo

import "firebond-ex-api.com/db/models"

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

var getFiatRateRecordForSymbolTestCases = map[string]struct {
	inputSymbol     string
	inputFiatSymbol string
	wantRate        models.Rate
	wantErr         error
}{
	"successfully fetched fiat-symbol rate record": {
		inputSymbol:     "BTC",
		inputFiatSymbol: "USD",
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
	"failed to fetch fiat-symbol rate record due to invalid symbol": {
		inputSymbol:     "XXX",
		inputFiatSymbol: "USD",
		wantRate:        models.Rate{},
		wantErr:         ErrNoDocuments,
	},
	"failed to fetch fiat-symbol rate record due to invalid fiat symbol": {
		inputSymbol:     "BTC",
		inputFiatSymbol: "XXX",
		wantRate:        models.Rate{},
		wantErr:         ErrFiatRateToSymbolNotFound,
	},
}

var upsertTestCases = map[string]struct {
	inputRate models.Rate
	wantRate  models.Rate
	wantErr   error
}{
	"successfully creates a new rate record that doesn't already exist": {
		inputRate: models.Rate{
			Symbol: "ATX",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			},
		},
		wantRate: models.Rate{
			Symbol: "ATX",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			},
		},
		wantErr: nil,
	},
	"successfully updates a document that already exists": {
		inputRate: models.Rate{
			Symbol: "BTC",
			FiatPrices: map[string]interface{}{
				"USD": 38000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			},
		},
		wantRate: models.Rate{
			Symbol: "BTC",
			FiatPrices: map[string]interface{}{
				"USD": 38000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			},
		},
		wantErr: nil,
	},
}

var getAllRatesTestCases = map[string]struct {
	wantRates []models.Rate
	wantErr   error
}{
	"successfully fetched all rates": {
		wantRates: []models.Rate{
			{Symbol: "BTC", FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			}},
			{Symbol: "ETH", FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 38000.0,
				"EUR": 35000.0,
			}},
		},
		wantErr: nil,
	},
}
