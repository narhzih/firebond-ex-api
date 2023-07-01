package cc

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strings"
)

var (
	ErrNonExistentCoinPairError        = fmt.Errorf("market does not exist for specified coin pair")
	ConstMarketDoesNotExistForCoinPair = "cccagg_or_exchange market does not exist for this coin pair "
)

type ExchangeApiConn struct {
	ApiKey string
	ApiUrl string
	Logger zerolog.Logger
}

type ExchangeApiError struct {
	CoolDown   int64       `json:"CoolDown"`
	Data       interface{} `json:"Data"`
	HasWarning bool        `json:"HasWarning"`
	Message    string      `json:"Message"`
	RateLimit  interface{} `json:"RateLimit"`
	Response   string      `json:"Response"`
	Type       int64       `json:"Type"`
}

type ExchangeApiData interface {
}

type ExchangeApiResponse struct {
	Error ExchangeApiError `json:"error"`
	Data  ExchangeApiData  `json:"data"`
}

func NewExchangeApiConn(apiKey string, logger zerolog.Logger) *ExchangeApiConn {
	return &ExchangeApiConn{
		ApiKey: apiKey,
		ApiUrl: "https://min-api.cryptocompare.com/data",
		Logger: logger,
	}
}

func (e *ExchangeApiConn) GetSymbolToFiatRate(cryptoSymbol, fiatSymbol string) (ExchangeApiResponse, error) {
	reqUrl := fmt.Sprintf("%v/price?fsym=%v&tsyms=%v&api_key=%v", e.ApiUrl, cryptoSymbol, fiatSymbol, e.ApiKey)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		e.Logger.Err(err).Msg("unable to build request")
		return ExchangeApiResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		e.Logger.Err(err).Msg("Unable to fetch response while building request")
		return ExchangeApiResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExchangeApiResponse{}, err
	}

	// first check for error
	var returnedData ExchangeApiResponse
	err = json.Unmarshal(body, &returnedData.Error)
	if err != nil {
		return ExchangeApiResponse{}, err
	}

	if len(strings.TrimSpace(returnedData.Error.Response)) == 0 {
		// This means there wasn't any error in the data
		// now proceed to parse the data
		err = json.Unmarshal(body, &returnedData.Data)
		if err != nil {
			return ExchangeApiResponse{}, err
		}

		return returnedData, nil
	}

	// Check for the type of error returned
	if strings.Contains(returnedData.Error.Message, ConstMarketDoesNotExistForCoinPair) {
		return ExchangeApiResponse{}, ErrNonExistentCoinPairError
	}
	return ExchangeApiResponse{}, fmt.Errorf("%v", returnedData.Error.Message)
}

func (e *ExchangeApiConn) GetSymbolRateForSupportedFiats() error {
	//TODO implement me
	panic("implement me")
}

func (e *ExchangeApiConn) GetSymbolHistory() error {
	//TODO implement me
	panic("implement me")
}
