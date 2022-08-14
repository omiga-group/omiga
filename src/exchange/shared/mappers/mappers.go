package mappers

import (
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
)

func FromOrderToEventOrder(src models.OrderBook) orderbookv1.OrderBook {
	order := orderbookv1.OrderBook{
		ExchangeId:      src.ExchangeId,
		BaseCurrency:    fromCurrencyToEventCurrency(src.BaseCurrency),
		CounterCurrency: fromCurrencyToEventCurrency(src.CounterCurrency),
	}

	order.Bids = make([]orderbookv1.OrderBookEntry, 0)
	for _, bid := range src.Bids {
		order.Bids = append(order.Bids, fromOrderBookEntryToEventOrderBookEntry(bid))
	}

	order.Asks = make([]orderbookv1.OrderBookEntry, 0)
	for _, ask := range src.Asks {
		order.Asks = append(order.Asks, fromOrderBookEntryToEventOrderBookEntry(ask))
	}

	return order
}

func fromOrderBookEntryToEventOrderBookEntry(src models.OrderBookEntry) orderbookv1.OrderBookEntry {
	return orderbookv1.OrderBookEntry{
		Quantity: fromMoneyToEventMoney(src.Quantity),
		Price:    fromMoneyToEventMoney(src.Price),
	}
}

func fromMoneyToEventMoney(src models.Money) orderbookv1.Money {
	return orderbookv1.Money{
		Amount:   src.Amount,
		Scale:    src.Scale,
		Currency: fromCurrencyToEventCurrency(src.Currency),
	}
}

func fromCurrencyToEventCurrency(src models.Currency) orderbookv1.Currency {
	return orderbookv1.Currency{
		Name:         src.Name,
		Code:         src.Code,
		MaxPrecision: src.MaxPrecision,
		Digital:      src.Digital,
	}
}
