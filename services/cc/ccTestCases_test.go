package cc

var getSymbolToFiatRateTestCases = map[string]struct {
	cryptoSymbol string
	fiatSymbol   string
	wantErr      error
}{
	"successfully fetch exchange rate of fiat-symbol": {
		cryptoSymbol: "BTC",
		fiatSymbol:   "USD",
		wantErr:      nil,
	},
	"report error for wrong symbol": {
		cryptoSymbol: "XZQA",
		fiatSymbol:   "USD",
		wantErr:      ErrNonExistentCoinPairError,
	},
	"report error for wrong fiat": {
		cryptoSymbol: "BTC",
		fiatSymbol:   "XXXX",
		wantErr:      ErrNonExistentCoinPairError,
	},
}
