package mappers

import (
	"github.com/omiga-group/omiga/src/exchange/coingeko/models"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
)

func FromCoingekoCoinToCoin(
	coin coingekov3.Coin) models.Coin {
	return models.Coin{
		Symbol: coin.Symbol,
		Name:   coin.Name,
	}
}
