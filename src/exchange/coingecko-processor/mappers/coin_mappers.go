package mappers

import (
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	coingeckov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingecko/v3"
)

func FromCoingeckoCoinToCoin(
	coin coingeckov3.Coin) models.Coin {
	return models.Coin{
		Symbol: coin.Symbol,
		Name:   coin.Name,
	}
}
