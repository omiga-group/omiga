package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func CoinbaseProductsToTradingPairs(products []coinbasepro.Product) []models.TradingPair {
	return slices.Map(products, func(product coinbasepro.Product) models.TradingPair {
		return models.TradingPair{
			Symbol:  product.DisplayName,
			Base:    product.BaseCurrency,
			Counter: product.QuoteCurrency,
		}
	})
}
