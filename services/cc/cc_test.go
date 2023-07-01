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

func TestExchangeApiConn_GetSymbolToFiatRate(t *testing.T) {

	for name, tc := range getSymbolToFiatRateTestCases {
		t.Run(name, func(t *testing.T) {
			gotRes, gotErr := exApiConn.GetSymbolToFiatRate(tc.cryptoSymbol, tc.fiatSymbol)
			assert.Equal(t, tc.wantErr, gotErr)
			if gotErr == nil {
				assert.Equal(t, "", gotRes.Error.Response)
			}
		})
	}

}
