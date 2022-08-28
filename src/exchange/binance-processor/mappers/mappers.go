package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
)

func FromBinanceOrderBookToModelOrderBook(
	baseCurrency exchangeModels.Currency,
	counterCurrency exchangeModels.Currency,
	orderBook []models.BinanceOrderBookEntry) exchangeModels.OrderBook {
	asks := slices.Filter(orderBook, func(entry models.BinanceOrderBookEntry) bool {
		return entry.Ask != nil
	})

	bids := slices.Filter(orderBook, func(entry models.BinanceOrderBookEntry) bool {
		return entry.Bid != nil
	})

	convertedAsks := slices.Map(asks, func(ask models.BinanceOrderBookEntry) exchangeModels.OrderBookEntry {
		quantity, amount, err := ask.Ask.Parse()
		if err != nil {
			quantity = -1
			amount = -1
		}

		return exchangeModels.OrderBookEntry{
			Quantity: exchangeModels.Money{
				Amount:   quantity,
				Scale:    1,
				Currency: baseCurrency,
			},
			Price: exchangeModels.Money{
				Amount:   amount,
				Scale:    1,
				Currency: counterCurrency,
			},
		}
	})

	convertedBids := slices.Map(bids, func(bid models.BinanceOrderBookEntry) exchangeModels.OrderBookEntry {
		quantity, amount, err := bid.Bid.Parse()
		if err != nil {
			quantity = -1
			amount = -1
		}

		return exchangeModels.OrderBookEntry{
			Quantity: exchangeModels.Money{
				Amount:   quantity,
				Scale:    1,
				Currency: baseCurrency,
			},
			Price: exchangeModels.Money{
				Amount:   amount,
				Scale:    1,
				Currency: counterCurrency,
			},
		}
	})

	return exchangeModels.OrderBook{
		BaseCurrency:    baseCurrency,
		CounterCurrency: counterCurrency,
		Asks: slices.Filter(convertedAsks, func(entry exchangeModels.OrderBookEntry) bool {
			return entry.Quantity.Amount != -1 && entry.Price.Amount != -1
		}),
		Bids: slices.Filter(convertedBids, func(entry exchangeModels.OrderBookEntry) bool {
			return entry.Quantity.Amount != -1 && entry.Price.Amount != -1
		}),
	}
}
