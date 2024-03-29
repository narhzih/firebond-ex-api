package cc

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var exApiConn *ExchangeApiConn
var logger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()

func TestMain(m *testing.M) {
	// load env
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	exApiConn = NewExchangeApiConn(os.Getenv("CRYPTO_COMPARE_API_KEY"), logger)
	code := m.Run()
	os.Exit(code)
}

func TestExchangeApiConn_GetSupportedCryptoToFiatPairsForBinance(t *testing.T) {
	for name, tc := range getSupportedCryptoToFiatPairsForBinanceTestCases {
		t.Run(name, func(t *testing.T) {
			gotRes, gotErr := exApiConn.GetSupportedCryptoToFiatPairsForBinance()
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, "Success", gotRes.Response)
			}
		})
	}
}

func TestExchangeApiConn_GetRatesForFsymsAndTsyms(t *testing.T) {
	for name, tc := range getRatesForFsymsAndTsymsTestCases {
		t.Run(name, func(t *testing.T) {
			gotRes, gotErr := exApiConn.GetRatesForFsymsAndTsyms(tc.inputFsyms, tc.inputTsyms)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				for fiat, _ := range gotRes {
					assert.Equal(t, tc.inputTsyms, fiat)
				}
			}
		})
	}
}

func TestExchangeApiConn_GetSymbolToFiatHistory(t *testing.T) {
	for name, tc := range getSymbolToFiatHistoryTestCases {
		t.Run(name, func(t *testing.T) {
			gotRes, gotErr := exApiConn.GetSymbolToFiatHistory(tc.inputSymbol, tc.inputFiat)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, 25, len(gotRes.Data.Data))
			}
		})
	}
}
