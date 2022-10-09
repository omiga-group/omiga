package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/omiga-group/omiga/src/exchange/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/models"
)

type ApiClient interface {
	GetMarkets() (models.MarketsMap, error)
}

type apiResult[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
}

func NewFtxApiClient(cfg configuration.FtxConfig) ApiClient {
	httpClient := createDefaultHttpClient(cfg.Timeout)
	return ftxApiClient{
		http:    httpClient,
		baseUrl: cfg.ApiUrl,
	}
}

type ftxApiClient struct {
	http    http.Client
	baseUrl string
}

func (ftx ftxApiClient) GetMarkets() (models.MarketsMap, error) {
	resp, err := ftx.http.Get(ftx.baseUrl + "/markets")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res apiResult[[]models.Market]
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	} else if !res.Success {
		return nil, fmt.Errorf("ftx api was not successful")
	}

	mm := models.MarketsMap{}
	for _, m := range res.Result {
		if !m.Enabled || m.Type != models.MarketTypeSpot {
			continue
		}

		mm[m.Name] = m
	}

	return mm, nil
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
