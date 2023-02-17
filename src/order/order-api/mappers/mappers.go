package mappers

import (
	"github.com/life4/genesis/slices"
	graphqlmodels "github.com/omiga-group/omiga/src/order/order-api/graphql/models"
	"github.com/omiga-group/omiga/src/order/shared/models"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
)

func FromSubmitOrderInputToOrder(src graphqlmodels.SubmitOrderInput) models.Order {
	order := models.Order{
		OrderDetails: models.OrderDetails{
			BaseCurrency:    fromCurrencyInputToCurrency(src.OrderDetails.BaseCurrency),
			CounterCurrency: fromCurrencyInputToCurrency(src.OrderDetails.CounterCurrency),
			Type:            orderv1.ValuesToOrderType[src.OrderDetails.Type],
			Side:            orderv1.ValuesToOrderSide[src.OrderDetails.Side],
			Quantity:        fromQuantityInputToQuantity(src.OrderDetails.Quantity),
			Price:           fromQuantityInputToQuantity(src.OrderDetails.Price),
		},
	}

	order.PreferredExchanges = slices.Map(
		src.PreferredExchanges,
		func(preferredExchange *graphqlmodels.ExchangeInput) models.Exchange {
			return fromExchangeInputToExchange(preferredExchange)
		})

	return order
}

func fromQuantityInputToQuantity(src *graphqlmodels.QuantityInput) models.Quantity {
	return models.Quantity{
		Amount: int64(src.Amount),
		Scale:  int32(src.Scale),
	}
}

func fromCurrencyInputToCurrency(src *graphqlmodels.CurrencyInput) models.Currency {
	return models.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: int32(src.MaxPrecision),
		Digital:      src.Digital,
	}
}

func fromExchangeInputToExchange(src *graphqlmodels.ExchangeInput) models.Exchange {
	return models.Exchange{
		Id: src.ID,
	}
}
