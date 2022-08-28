package mappers

import (
	"github.com/life4/genesis/slices"
	graphqlmodels "github.com/omiga-group/omiga/src/order/order-api/graphql/models"
	"github.com/omiga-group/omiga/src/order/shared/models"
)

func FromSubmitOrderInputToOrder(src graphqlmodels.SubmitOrderInput) models.Order {
	order := models.Order{
		OrderDetails: models.OrderDetails{
			BaseCurrency:    fromCurrencyInputToCurrency(src.OrderDetails.BaseCurrency),
			CounterCurrency: fromCurrencyInputToCurrency(src.OrderDetails.CounterCurrency),
			Type:            models.OrderType(src.OrderDetails.Type),
			Side:            models.OrderSide(src.OrderDetails.Side),
			Quantity:        fromMoneyInputToCurrency(src.OrderDetails.Quantity),
			Price:           fromMoneyInputToCurrency(src.OrderDetails.Price),
		},
	}

	order.PreferredExchanges = slices.Map(
		src.PreferredExchanges,
		func(preferredExchange *graphqlmodels.ExchangeInput) models.Exchange {
			return fromExchangeInputToExchange(preferredExchange)
		})

	return order
}

func fromMoneyInputToCurrency(src *graphqlmodels.MoneyInput) models.Money {
	return models.Money{
		Amount:   src.Amount,
		Scale:    src.Scale,
		Currency: fromCurrencyInputToCurrency(src.Currency),
	}
}

func fromCurrencyInputToCurrency(src *graphqlmodels.CurrencyInput) models.Currency {
	return models.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: src.MaxPrecision,
		Digital:      src.Digital,
	}
}

func fromExchangeInputToExchange(src *graphqlmodels.ExchangeInput) models.Exchange {
	return models.Exchange{
		Id: src.ID,
	}
}
