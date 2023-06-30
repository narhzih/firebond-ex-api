package firebondmongo

import "firebond-ex-api.com/db/models"

var createCryptoToFiatRateData = map[string]struct {
	inputRate models.Rate
	wantRate  models.Rate
	wantErr   error
}{
	"successfully created crypto to fiat rate data": {
		inputRate: models.Rate{
			Symbol: "BTC",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 31000.5,
			},
		},
		wantRate: models.Rate{
			Symbol: "BTC",
			FiatPrices: map[string]interface{}{
				"USD": 30000.5,
				"GBP": 31000.5,
			},
		},
		wantErr: nil,
	},
}
