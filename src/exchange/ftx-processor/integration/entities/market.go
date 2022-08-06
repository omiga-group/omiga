package entities

import "encoding/json"

type Market struct {
	Name                  string  `json:"name"`
	BaseCurrency          string  `json:"baseCurrency"`
	QuoteCurrency         string  `json:"quoteCurrency"`
	QuoteVolume24H        float64 `json:"quoteVolume24h"`
	Change1H              float64 `json:"change1h"`
	Change24H             float64 `json:"change24h"`
	ChangeBod             float64 `json:"changeBod"`
	HighLeverageFeeExempt bool    `json:"highLeverageFeeExempt"`
	MinProvideSize        float64 `json:"minProvideSize"`
	Type                  string  `json:"type"`
	Underlying            string  `json:"underlying"`
	Enabled               bool    `json:"enabled"`
	Ask                   float64 `json:"ask"`
	Bid                   float64 `json:"bid"`
	Last                  float64 `json:"last"`
	PostOnly              bool    `json:"postOnly"`
	Price                 float64 `json:"price"`
	PriceIncrement        float64 `json:"priceIncrement"`
	SizeIncrement         float64 `json:"sizeIncrement"`
	Restricted            bool    `json:"restricted"`
	VolumeUsd24H          float64 `json:"volumeUsd24h"`
	LargeOrderThreshold   float64 `json:"largeOrderThreshold"`
	IsEtfMarket           bool    `json:"isEtfMarket"`
}

func MarketsFromResponse(data []byte) ([]Market, error) {
	var response struct {
		Success bool     `json:"success"`
		Result  []Market `json:"result"`
	}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}
