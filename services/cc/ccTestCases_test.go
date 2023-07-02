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
