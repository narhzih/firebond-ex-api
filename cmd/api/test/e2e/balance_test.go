package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetWalletBalance(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	t.Run("successfully fetched wallet balance", func(t *testing.T) {
		fetchRes := struct {
			Message string `json:"message"`
		}{}
		req, err := http.NewRequest(http.MethodGet, "/v1/balance/0x00000000219ab540356cbb839cbe05303d7705fa", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusOK, res.Code)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("could not parse response body")
		}
		err = json.Unmarshal(resBody, &fetchRes)
		if err != nil {
			t.Fatal(fmt.Sprintf(err.Error()))
		}
		assert.Equal(t, "Balance fetched successfully", fetchRes.Message)
	})

}
