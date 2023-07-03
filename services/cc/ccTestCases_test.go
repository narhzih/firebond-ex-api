package cc

var getSupportedCryptoToFiatPairsForBinanceTestCases = map[string]struct {
	cryptoSymbol string
	fiatSymbol   string
	wantErr      error
}{
	"successfully fetched symbols and their supported fiats": {
		wantErr: nil,
	},
}

var getRatesForFsymsAndTsymsTestCases = map[string]struct {
	inputFsyms string
	inputTsyms string
	wantErr    error
}{
	"successfully fetched fsyms and tsyms data": {
		inputFsyms: "BTC",
		inputTsyms: "EUR",
		wantErr:    nil,
	},
}

var getSymbolToFiatHistoryTestCases = map[string]struct {
	inputSymbol string
	inputFiat   string
	wantErr     error
}{
	"successfully fetched history for coin pair": {
		inputSymbol: "BTC",
		inputFiat:   "EUR",
		wantErr:     nil,
	},
	"return error for incompatible coin-pair": {
		inputSymbol: "BTC",
		inputFiat:   "USD",
		wantErr:     ExchangeApiErrIncompatibleCoinPair,
	},
}
