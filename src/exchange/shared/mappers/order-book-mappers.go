package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
)

func FromOrderBookToEventOrderBook(src models.OrderBook) orderbookv1.OrderBook {
	order := orderbookv1.OrderBook{
		ExchangeId:      src.ExchangeId,
		Time:            src.Time,
		BaseCurrency:    fromCurrencyToEventCurrency(src.BaseCurrency),
		CounterCurrency: fromCurrencyToEventCurrency(src.CounterCurrency),
	}

	order.Bids = slices.Map(
		src.Bids,
		func(bid models.OrderBookEntry) orderbookv1.OrderBookEntry {
			return fromOrderBookEntryToEventOrderBookEntry(bid)
		})

	order.Asks = slices.Map(
		src.Asks,
		func(ask models.OrderBookEntry) orderbookv1.OrderBookEntry {
			return fromOrderBookEntryToEventOrderBookEntry(ask)
		})

	return order
}

func fromOrderBookEntryToEventOrderBookEntry(src models.OrderBookEntry) orderbookv1.OrderBookEntry {
	return orderbookv1.OrderBookEntry{
		Quantity: fromQuantityToEventQuantity(src.Quantity),
		Price:    fromMoneyToEventMoney(src.Price),
	}
}

func fromQuantityToEventQuantity(src models.Quantity) orderbookv1.Quantity {
	return orderbookv1.Quantity{
		Amount: src.Amount,
		Scale:  src.Scale,
	}
}

func fromMoneyToEventMoney(src models.Money) orderbookv1.Money {
	return orderbookv1.Money{
		Quantity: fromQuantityToEventQuantity(src.Quantity),
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