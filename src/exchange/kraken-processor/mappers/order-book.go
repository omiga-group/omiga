package mappers

import (
	"strconv"
	"strings"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/kraken-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/shared/enterprise/decimal"
)

func KrakenOrderBookToOrderBook(
	baseCurrency exchangeModels.Currency,
	counterCurrency exchangeModels.Currency,
	orderBook []models.KrakenOrderBookEntry) exchangeModels.OrderBook {
	asks := slices.Filter(orderBook, func(entry models.KrakenOrderBookEntry) bool {
		return entry.Ask != nil
	})

	bids := slices.Filter(orderBook, func(entry models.KrakenOrderBookEntry) bool {
		return entry.Bid != nil
	})

	convertedAsks := slices.Map(asks, func(entry models.KrakenOrderBookEntry) exchangeModels.OrderBookEntry {
		timePeices := strings.Split(string(entry.Ask.Time), ".")
		timePeice1, _ := strconv.ParseInt(timePeices[0], 10, 64)
		timePeice2, _ := strconv.ParseInt(timePeices[1], 10, 64)

		orderbookEntry := exchangeModels.OrderBookEntry{
			Time: time.Unix(timePeice1, timePeice2),
		}

		if decimal, err := decimal.StringToDecimal(string(entry.Ask.Volume)); err != nil {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		if decimal, err := decimal.StringToDecimal(string(entry.Ask.Price)); err != nil {
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

	convertedBids := slices.Map(bids, func(entry models.KrakenOrderBookEntry) exchangeModels.OrderBookEntry {
		timePeices := strings.Split(string(entry.Bid.Time), ".")
		timePeice1, _ := strconv.ParseInt(timePeices[0], 10, 64)
		timePeice2, _ := strconv.ParseInt(timePeices[1], 10, 64)

		orderbookEntry := exchangeModels.OrderBookEntry{
			Time: time.Unix(timePeice1, timePeice2),
		}

		if decimal, err := decimal.StringToDecimal(string(entry.Bid.Volume)); err != nil {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Scale: -1,
			}
		} else {
			orderbookEntry.Quantity = exchangeModels.Quantity{
				Amount: decimal.Amount,
				Scale:  decimal.Scale,
			}
		}

		if decimal, err := decimal.StringToDecimal(string(entry.Bid.Price)); err != nil {
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
