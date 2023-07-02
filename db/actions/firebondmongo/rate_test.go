package firebondmongo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRateActions_GetCryptoRatesBySymbol(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range getCryptoRatesBySymbolTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRate, gotErr := ra.GetCryptoRatesBySymbol(tc.inputSymbol)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
			}
		})
	}
}

func TestRateActions_GetFiatRateRecordForSymbol(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range getFiatRateRecordForSymbolTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRate, gotErr := ra.GetFiatRateRecordForSymbol(tc.inputSymbol, tc.inputFiatSymbol)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil != gotErr {
				if nil == gotErr {
					assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
					assert.Equal(t, tc.wantRate.FiatPrices[tc.inputFiatSymbol], gotRate.FiatPrices[tc.inputFiatSymbol])
				}
			}
		})
	}
}

func TestRateActions_GetAllRates(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range getAllRatesTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRates, gotErr := ra.GetAllRates()
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, 2, len(gotRates))
				assert.Equal(t, "BTC", gotRates[0].Symbol)
			}
		})
	}

}

func TestRateActions_UpSert(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range upsertTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotErr := ra.UpSert(tc.inputRate)
			assert.Equal(t, tc.wantErr, gotErr)
		})
	}
}
