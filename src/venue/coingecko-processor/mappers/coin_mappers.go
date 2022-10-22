package mappers

import (
	coingeckov3 "github.com/omiga-group/omiga/src/venue/coingecko-processor/coingeckoclient/v3"
	currencyrepo "github.com/omiga-group/omiga/src/venue/shared/entities/currency"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func FromCoingeckoCoinToCurrency(
	coin coingeckov3.Coin) models.Currency {
	return models.Currency{
		Symbol: coin.Symbol,
		Name:   coin.Name,
		Type:   currencyrepo.TypeDIGITAL,
	}
}
