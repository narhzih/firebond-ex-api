package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetAllRates(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	t.Run("successfully fetched all rates", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusOK, res.Code)
	})
}

func TestGetRate(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	t.Run("successfully fetched the symbol to fiat rate", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/BTC/EUR", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusOK, res.Code)
	})

	t.Run("returns error for incompatible coin pair", func(t *testing.T) {
		fetchRes := struct {
			Message string `json:"message"`
			Err     string `json:"err"`
		}{}
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/BTC/GGGH", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusBadRequest, res.Code)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("could not parse response body")
		}
		err = json.Unmarshal(resBody, &fetchRes)
		if err != nil {
			t.Fatal(fmt.Sprintf(err.Error()))
		}
		assert.Equal(t, "BTC-GGGH coin pair is not supported on Binance Market", fetchRes.Err)
		assert.Equal(t, "An error occurred", fetchRes.Message)
	})
}

func TestSymbolRates(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	t.Run("successfully fetched rates for specified symbol", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/BTC", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusOK, res.Code)
	})

	t.Run("returns error for symbol not in database", func(t *testing.T) {
		fetchRes := struct {
			Message string `json:"message"`
			Err     string `json:"err"`
		}{}
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/GGG", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusBadRequest, res.Code)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("could not parse response body")
		}
		err = json.Unmarshal(resBody, &fetchRes)
		if err != nil {
			t.Fatal(fmt.Sprintf(err.Error()))
		}
		assert.Equal(t, "No exchange record for this symbol. Check that you entered the correct symbol, or try again in 10 minutes", fetchRes.Err)
		assert.Equal(t, "An error occurred", fetchRes.Message)
	})
}

func TestGetHistory(t *testing.T) {
	if testing.Short() {
		t.Skip(skipMessage)
	}

	t.Run("successfully fetched history for symbol-fiat", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/history/BTC/EUR", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusOK, res.Code)
	})

	t.Run("returns error for incompatible pair", func(t *testing.T) {
		fetchRes := struct {
			Message string `json:"message"`
			Err     string `json:"err"`
		}{}
		req, err := http.NewRequest(http.MethodGet, "/v1/rates/history/BTC/GGGG", nil)
		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		res := executeRequest(req)
		checkResponseCode(t, http.StatusBadRequest, res.Code)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("could not parse response body")
		}
		err = json.Unmarshal(resBody, &fetchRes)
		if err != nil {
			t.Fatal(fmt.Sprintf(err.Error()))
		}
		assert.Equal(t, "An error occurred", fetchRes.Message)
	})
}
