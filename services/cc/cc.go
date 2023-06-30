package cc

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type ExchangeApiConn struct {
	ApiKey string
	ApiUrl string
	Logger zerolog.Logger
}

func NewExchangeApiConn(apiKey string, logger zerolog.Logger) *ExchangeApiConn {
	return &ExchangeApiConn{
		ApiKey: apiKey,
		ApiUrl: "https://min-api.cryptocompare.com/data",
		Logger: logger,
	}
}

func (e *ExchangeApiConn) GetSymbolToFiatRate(cryptoSymbol, fiatSymbol string) (interface{}, error) {
	reqUrl := fmt.Sprintf("%v/price?fsym=%v&tsym=%v&api_key=%v", e.ApiUrl, cryptoSymbol, fiatSymbol, e.ApiKey)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		e.Logger.Err(err).Msg("unable to build request")
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		e.Logger.Err(err).Msg("Unable to fetch response while building request")
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (e *ExchangeApiConn) GetSymbolRateForSupportedFiats() error {
	//TODO implement me
	panic("implement me")
}

func (e *ExchangeApiConn) GetSymbolHistory() error {
	//TODO implement me
	panic("implement me")
}
