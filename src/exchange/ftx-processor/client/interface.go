package client

import "github.com/omiga-group/omiga/src/exchange/ftx-processor/models"

type ApiClient interface {
	GetMarkets() (models.MarketsMap, error)
}
