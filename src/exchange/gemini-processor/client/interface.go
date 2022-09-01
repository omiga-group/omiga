package client

import "github.com/omiga-group/omiga/src/exchange/gemini-processor/models"

type ApiClient interface {
	GetMarkets() (models.MarketsMap, error)
}
