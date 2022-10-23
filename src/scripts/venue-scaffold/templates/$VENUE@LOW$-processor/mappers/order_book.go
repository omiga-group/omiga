package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/shared/enterprise/decimal"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/venue/shared/models"
)

func $VENUE@PAS$OrderBookToOrderBook(
	baseCurrency exchangeModels.OrderCurrency,
	counterCurrency exchangeModels.OrderCurrency,
	orderBook []models.$VENUE@PAS$OrderBookEntry) exchangeModels.OrderBook {
	asks := slices.Filter(orderBook, func(entry models.$VENUE@PAS$OrderBookEntry) bool {
		return entry.Ask != nil
	})

	bids := slices.Filter(orderBook, func(entry models.$VENUE@PAS$OrderBookEntry) bool {
		return entry.Bid != nil
	})

	convertedAsks := slices.Map(asks, func(entry models.$VENUE@PAS$OrderBookEntry) exchangeModels.OrderBookEntry {
		orderbookEntry := exchangeModels.OrderBookEntry{
			Time: entry.Time,
		}

		if decimal, err := decimal.StringToDecimal(entry.Ask.Quantity); err != nil {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		if decimal, err := decimal.StringToDecimal(entry.Ask.Price); err != nil {
			orderbookEntry.Price = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Price = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		return orderbookEntry
	})

	convertedBids := slices.Map(bids, func(entry models.$VENUE@PAS$OrderBookEntry) exchangeModels.OrderBookEntry {
		orderbookEntry := exchangeModels.OrderBookEntry{
			Time: entry.Time,
		}

		if decimal, err := decimal.StringToDecimal(entry.Bid.Quantity); err != nil {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		if decimal, err := decimal.StringToDecimal(entry.Bid.Price); err != nil {
			orderbookEntry.Price = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Price = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		return orderbookEntry
	})

	return exchangeModels.OrderBook{
		BaseCurrency:    baseCurrency,
		CounterCurrency: counterCurrency,
		Asks: slices.Filter(convertedAsks, func(entry exchangeModels.OrderBookEntry) bool {
			return entry.Quantity.Scale != -1 && entry.Price.Scale != -1
		}),
		Bids: slices.Filter(convertedBids, func(entry exchangeModels.OrderBookEntry) bool {
			return entry.Quantity.Scale != -1 && entry.Price.Scale != -1
		}),
	}
}
