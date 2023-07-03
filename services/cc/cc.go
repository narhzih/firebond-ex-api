package cc

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strings"
	"time"
)

var ExchangeApiErrStringIncompatibleCoinPair = "Binance market does not exist for this coin pair"
var ExchangeApiErrIncompatibleCoinPair = fmt.Errorf("binance market does not exist for this coin pair")

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
	// api - // https://min-api.cryptocompare.com/data/v4/all/exchange?e=Binance
	// docs -https://min-api.cryptocompare.com/documentation
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

func (e *ExchangeApiConn) GetRatesForFsymsAndTsyms(fsym, tsyms string) (map[string]interface{}, error) {
	// api - // https://min-api.cryptocompare.com/data/price
	// docs -https://min-api.cryptocompare.com/documentation
	reqUrl := fmt.Sprintf("%v/price?fsym=%v&tsyms=%v&api_key=%v", e.ApiUrl, fsym, tsyms, e.ApiKey)
	res, err := e.doRequest(reqUrl, nil)
	if err != nil {
		return map[string]interface{}{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var returnedData ExchangeApiFsymsToTsymsResponse
	err = json.Unmarshal(body, &returnedData.Error)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(strings.TrimSpace(returnedData.Error.Response)) == 0 {
		// This means there's no error
		err = json.Unmarshal(body, &returnedData.Data)
		if err != nil {
			return map[string]interface{}{}, err
		}

		return returnedData.Data, nil
	}
	return returnedData.Data, fmt.Errorf("%v", returnedData.Error.Message)
}

func (e *ExchangeApiConn) GetSymbolToFiatHistory(symbol, fiat string) (ExchangeApiHistoryResponse, error) {
	// api - // https://min-api.cryptocompare.com/data/v2/histohour
	// docs -https://min-api.cryptocompare.com/documentation
	var returnedData ExchangeApiHistoryResponse
	now := time.Now()
	reqUrl := fmt.Sprintf("%v/v2/histohour?fsym=%v&tsym=%v&limit=24&e=Binance&toTs=%v&api_key=%v", e.ApiUrl, symbol, fiat, now.Unix(), e.ApiKey)
	res, err := e.doRequest(reqUrl, nil)
	if err != nil {
		return returnedData, nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return returnedData, err
	}
	err = json.Unmarshal(body, &returnedData)
	if err != nil {
		return returnedData, err
	}
	if returnedData.Response == "Success" {
		var newData []ExchangeApiHistoryResponseDataSH

		for _, data := range returnedData.Data.Data {
			data.ReadableTime = time.Unix(data.Time, 0)
			newData = append(newData, data)
		}
		returnedData.Data.Data = newData
		return returnedData, nil
	}

	if strings.Contains(returnedData.Message, ExchangeApiErrStringIncompatibleCoinPair) {
		return returnedData, ExchangeApiErrIncompatibleCoinPair
	}
	return returnedData, fmt.Errorf("%v", returnedData.Message)
}
