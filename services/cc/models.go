package cc

import "time"

// ExchangeApiResponse is the struct that represents a response to
// https://min-api.cryptocompare.com/data/v4/all/exchange?e=Binance
type ExchangeApiResponse struct {
	CoolDown int64 `json:"CoolDown"`
	Data     struct {
		Data struct {
			Exchanges struct {
				Binance struct {
					IsActive  bool `json:"isActive"`
					IsTopTier bool `json:"isTopTier"`
					Pairs     map[string]struct {
						Tsyms map[string]struct {
							HistoMinuteEnd   string `json:"histo_minute_end"`
							HistoMinuteStart string `json:"histo_minute_start"`
						} `json:"tsyms"`
					} `json:"pairs"`
				} `json:"Binance"`
			} `json:"exchanges"`
		} `json:"Data"`
	} `json:"Data"`
	HasWarning bool        `json:"HasWarning"`
	Message    string      `json:"Message"`
	RateLimit  interface{} `json:"RateLimit"`
	Response   string      `json:"Response"`
	Type       int64       `json:"Type"`
}

// ExchangeApiFsymsToTsymsError describes the error data returned for
// https://
type ExchangeApiFsymsToTsymsError struct {
	CoolDown   int64       `json:"CoolDown"`
	HasWarning bool        `json:"HasWarning"`
	Message    string      `json:"Message"`
	RateLimit  interface{} `json:"RateLimit"`
	Response   string      `json:"Response"`
	Type       int64       `json:"Type"`
}

type ExchangeApiFsymsToTsymsData struct {
	Data interface{} `json:"Data"`
}
type ExchangeApiFsymsToTsymsResponse struct {
	Error ExchangeApiFsymsToTsymsError `json:"Error"`
	Data  map[string]interface{}       `json:"Data"`
}

type ExchangeApiHistoryResponse struct {
	HasWarning bool                           `json:"HasWarning"`
	Message    string                         `json:"Message"`
	RateLimit  interface{}                    `json:"RateLimit"`
	Response   string                         `json:"Response"`
	Type       int64                          `json:"Type"`
	Data       ExchangeApiHistoryResponseData `json:"Data"`
}
type ExchangeApiHistoryResponseData struct {
	Aggregated bool                               `json:"Aggregated"`
	Data       []ExchangeApiHistoryResponseDataSH `json:"Data"`
}

type ExchangeApiHistoryResponseDataSH struct {
	Close        float64   `json:"close"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Open         float64   `json:"open"`
	Time         int64     `json:"time"`
	ReadableTime time.Time `json:"readableTime"`
}
