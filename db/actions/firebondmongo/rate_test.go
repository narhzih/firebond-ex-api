package firebondmongo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDemoActions_CreateDemoData(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	for name, tc := range createCryptoToFiatRateData {
		t.Run(name, func(t *testing.T) {
			db := newTestDb(t)
			ra := NewRateActions(db, logger)
			gotRate, gotErr := ra.CreateCryptoToFiatRateData(tc.inputRate)
			assert.Equal(t, tc.wantErr, gotErr)
			if nil == gotErr {
				assert.Equal(t, tc.wantRate.Symbol, gotRate.Symbol)
			}
		})
	}
}
