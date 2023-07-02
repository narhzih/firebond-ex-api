package firebondmongo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRateActions_CreateCryptoToFiatRateData(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range createCryptoToFiatRateDataTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRate, gotErr := ra.CreateCryptoToFiatRateData(tc.inputRate)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
				assert.NotEmpty(t, gotRate.ID)
			}
		})
	}
}

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
			if nil != gotErr {
				assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
			}
		})
	}
}

func TestRateActions_CreateFiatRateRecordForSymbol(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range createFiatRateRecordForSymbolTestCases {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRate, gotErr := ra.CreateFiatRateRecordForSymbol(tc.inputSymbol, tc.inputFiatSymbol, tc.inputFiatValue)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil != gotErr {
				assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
				assert.Equal(t, tc.wantRate.FiatPrices[tc.inputFiatSymbol], gotRate.FiatPrices[tc.inputFiatSymbol])
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
				if nil != gotErr {
					assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
					assert.Equal(t, tc.wantRate.FiatPrices[tc.inputFiatSymbol], gotRate.FiatPrices[tc.inputFiatSymbol])
				}
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
