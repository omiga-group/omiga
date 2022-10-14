package mappers

import (
	currencyrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/currency"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	coingeckov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingecko/v3"
)

func FromCoingeckoCoinToCurrency(
	coin coingeckov3.Coin) models.Currency {
	return models.Currency{
		Symbol: coin.Symbol,
		Name:   coin.Name,
		Type:   currencyrepo.TypeDIGITAL,
	}
}
