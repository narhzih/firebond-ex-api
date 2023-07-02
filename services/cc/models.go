package cc

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
