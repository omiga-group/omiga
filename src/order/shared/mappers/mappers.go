package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/order/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
)

func FromOrderToEventOrder(src models.Order) orderv1.Order {
	order := orderv1.Order{
		Id:           src.Id,
		OrderDetails: fromOrderDetailsToEventOrderDetails(src.OrderDetails),
	}

	order.PreferredExchanges = slices.Map(
		src.PreferredExchanges,
		func(preferredExchange models.Exchange) orderv1.Exchange {
			return fromExchangeToEventExchange(preferredExchange)
		})

	return order
}

func fromOrderDetailsToEventOrderDetails(src models.OrderDetails) orderv1.OrderDetails {
	return orderv1.OrderDetails{
		BaseCurrency:    fromCurrencyToEventCurrency(src.BaseCurrency),
		CounterCurrency: fromCurrencyToEventCurrency(src.CounterCurrency),
		ReservedType:    src.Type,
		Side:            src.Side,
		Quantity:        fromQuantityToEventQuantity(src.Quantity),
		Price:           fromQuantityToEventQuantity(src.Price),
	}
}

func fromQuantityToEventQuantity(src models.Quantity) orderv1.Quantity {
	return orderv1.Quantity{
		Amount: src.Amount,
		Scale:  src.Scale,
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

	order.Bids = slices.Map(
		src.Bids,
		func(bid orderbookv1.OrderBookEntry) models.OrderBookEntry {
			return fromEventOrderBookEntryToOrderBookEntry(bid)
		})

	order.Asks = slices.Map(
		src.Asks,
		func(ask orderbookv1.OrderBookEntry) models.OrderBookEntry {
			return fromEventOrderBookEntryToOrderBookEntry(ask)
		})

	return order
}

func fromEventOrderBookEntryToOrderBookEntry(src orderbookv1.OrderBookEntry) models.OrderBookEntry {
	return models.OrderBookEntry{
		Quantity: fromEventQuantityToQuantity(src.Quantity),
		Price:    fromEventQuantityToQuantity(src.Price),
	}
}

func fromEventQuantityToQuantity(src orderbookv1.Quantity) models.Quantity {
	return models.Quantity{
		Amount: src.Amount,
		Scale:  src.Scale,
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
