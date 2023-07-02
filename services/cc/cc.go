package cc

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strings"
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

func (e *ExchangeApiConn) doRequest(reqUrl string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		e.Logger.Err(err).Msg("unable to build request")
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		e.Logger.Err(err).Msg("Unable to fetch response after building request")
		return nil, err
	}

	return res, nil
}

func (e *ExchangeApiConn) GetSupportedCryptoToFiatPairsForBinance() (ExchangeApiResponse, error) {
	reqUrl := fmt.Sprintf("%v/v4/all/exchanges?e=Binance&api_key=%v", e.ApiUrl, e.ApiKey)
	res, err := e.doRequest(reqUrl, nil)
	if err != nil {
		return ExchangeApiResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExchangeApiResponse{}, err
	}

	var returnedData ExchangeApiResponse
	err = json.Unmarshal(body, &returnedData)
	if err != nil {
		return ExchangeApiResponse{}, err
	}

	if strings.TrimSpace(returnedData.Response) == "Success" {
		// This means there wasn't any error in the data
		// now proceed to parse the data
		err = json.Unmarshal(body, &returnedData.Data)
		if err != nil {
			return ExchangeApiResponse{}, err
		}

		return returnedData, nil
	}
	//return ExchangeApiResponse{}, fmt.Errorf("%v", returnedData.Error.Message)
	e.Logger.Info().Msg(fmt.Sprintf("%v", returnedData.Data))
	return ExchangeApiResponse{}, fmt.Errorf("%v", returnedData.Message)

}

func (e *ExchangeApiConn) GetSymbolHistory() error {
	//TODO implement me
	panic("implement me")
}
