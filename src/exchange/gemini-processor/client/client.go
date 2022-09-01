package client

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/models"
)

type apiResult[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
}

func NewGeminiApiClient(cfg configuration.GeminiConfig) ApiClient {
	httpClient := createDefaultHttpClient(cfg.Timeout)
	return geminiApiClient{
		http:    httpClient,
		baseUrl: cfg.ApiUrl,
	}
}

type geminiApiClient struct {
	http    http.Client
	baseUrl string
}

func (ftx geminiApiClient) GetMarkets() (models.MarketsMap, error) {
	marketNames, err := getMarketNames()
	if err != nil {
		return nil, err
	}

	marketsMap := make(models.MarketsMap)
	for _, name := range marketNames {
		market, err := ftx.getMarketDetails(name)
		if err != nil {
			return nil, err
		}
		marketsMap[name] = market
	}

	return marketsMap, nil
}

func (ftx geminiApiClient) getMarketNames() (models.MarketNames, error) {
	resp, err := ftx.http.Get(ftx.baseUrl + "/symbols")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mns models.MarketNames
	err = json.Unmarshal(body, &mns)
	if err != nil {
		return nil, err
	}
}

func (ftx geminiApiClient) getMarketDetails(name string) (models.Markets, error) {
	resp, err := ftx.http.Get(ftx.baseUrl + "/symbols/details/" + name)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m models.Market
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func createDefaultHttpClient(timeout int) http.Client {
	dur := time.Second * time.Duration(timeout)
	return http.Client{
		Timeout: dur,
		Transport: &http.Transport{
			Dial:                (&net.Dialer{Timeout: dur}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
}
