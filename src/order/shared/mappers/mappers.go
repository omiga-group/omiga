package mappers

import (
	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
)

func FromSubmitOrderInputToOrder(src shared.SubmitOrderInput) models.Order {
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

	order.PreferredExchanges = make([]models.Exchange, 0)
	for _, preferredExchange := range src.PreferredExchanges {
		order.PreferredExchanges = append(order.PreferredExchanges, fromExchangeInputToExchange(preferredExchange))
	}

	return order
}

func fromMoneyInputToCurrency(src *shared.MoneyInput) models.Money {
	return models.Money{
		Amount:   src.Amount,
		Scale:    src.Scale,
		Currency: fromCurrencyInputToCurrency(src.Currency),
	}
}

func fromCurrencyInputToCurrency(src *shared.CurrencyInput) models.Currency {
	return models.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: src.MaxPrecision,
		Digital:      src.Digital,
	}
}

func fromExchangeInputToExchange(src *shared.ExchangeInput) models.Exchange {
	return models.Exchange{
		Id: src.ID,
	}
}

func FromOrderToEventOrder(src models.Order) orderv1.Order {
	order := orderv1.Order{
		Id:           src.Id,
		OrderDetails: fromOrderDetailsToEventOrderDetails(src.OrderDetails),
	}

	order.PreferredExchanges = make([]orderv1.Exchange, 0)
	for _, preferredExchange := range src.PreferredExchanges {
		order.PreferredExchanges = append(order.PreferredExchanges, fromExchangeToEventExchange(preferredExchange))
	}

	return order
}

func fromOrderDetailsToEventOrderDetails(src models.OrderDetails) orderv1.OrderDetails {
	return orderv1.OrderDetails{
		BaseCurrency:    fromCurrencyToEventCurrency(src.BaseCurrency),
		CounterCurrency: fromCurrencyToEventCurrency(src.CounterCurrency),
		Type:            orderv1.OrderType(src.Type),
		Side:            orderv1.OrderSide(src.Side),
		Quantity:        fromMoneyToEventMoney(src.Quantity),
		Price:           fromMoneyToEventMoney(src.Price),
	}
}

func fromMoneyToEventMoney(src models.Money) orderv1.Money {
	return orderv1.Money{
		Amount:   src.Amount,
		Scale:    src.Scale,
		Currency: fromCurrencyToEventCurrency(src.Currency),
	}
}

func fromCurrencyToEventCurrency(src models.Currency) orderv1.Currency {
	return orderv1.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: src.MaxPrecision,
		Digital:      src.Digital,
	}
}

func fromExchangeToEventExchange(src models.Exchange) orderv1.Exchange {
	return orderv1.Exchange{
		Id: src.Id,
	}
}

func FromEventOrderBookToOrderBook(src orderbookv1.OrderBook) models.OrderBook {
	order := models.OrderBook{
		BaseCurrency:    fromEventCurrencyToCurrency(src.BaseCurrency),
		CounterCurrency: fromEventCurrencyToCurrency(src.CounterCurrency),
	}

	order.Bids = make([]models.OrderBookEntry, 0)
	for _, bid := range src.Bids {
		order.Bids = append(order.Bids, fromEventOrderBookEntryToOrderBookEntry(bid))
	}

	order.Asks = make([]models.OrderBookEntry, 0)
	for _, ask := range src.Asks {
		order.Asks = append(order.Asks, fromEventOrderBookEntryToOrderBookEntry(ask))
	}

	return order
}

func fromEventOrderBookEntryToOrderBookEntry(src orderbookv1.OrderBookEntry) models.OrderBookEntry {
	return models.OrderBookEntry{
		Quantity: fromEventMoneyToMoney(src.Quantity),
		Price:    fromEventMoneyToMoney(src.Price),
	}
}

func fromEventMoneyToMoney(src orderbookv1.Money) models.Money {
	return models.Money{
		Amount:   src.Amount,
		Scale:    src.Scale,
		Currency: fromEventCurrencyToCurrency(src.Currency),
	}
}

func fromEventCurrencyToCurrency(src orderbookv1.Currency) models.Currency {
	return models.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: src.MaxPrecision,
		Digital:      src.Digital,
	}
}
