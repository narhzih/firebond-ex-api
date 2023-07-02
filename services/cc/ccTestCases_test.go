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
