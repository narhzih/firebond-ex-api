package firebondmongo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDemoActions_CreateDemoData(t *testing.T) {
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
