package services

import (
	"context"
	"sort"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"go.uber.org/zap"
)

type BinanceOrderBookEntry struct {
	Symbol string
	Time   time.Time
	Bid    *binance.Bid
	Ask    *binance.Ask
}

type OrderBookAggregator interface {
	UpdateOrderBook(
		symbol string,
		orderBook []BinanceOrderBookEntry)
}

type orderBookInput struct {
	symbol    string
	orderBook []BinanceOrderBookEntry
}

type orderBookAggregator struct {
	ctx                        context.Context
	logger                     *zap.SugaredLogger
	orderBookChannel           chan orderBookInput
	aggregatedBinanceOrderBook []BinanceOrderBookEntry
	orderBookPublisher         publishers.OrderBookPublisher
}

func NewOrderBookAggregator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	orderBookPublisher publishers.OrderBookPublisher) (OrderBookAggregator, error) {
	instance := orderBookAggregator{
		ctx:                        ctx,
		logger:                     logger,
		orderBookChannel:           make(chan orderBookInput),
		aggregatedBinanceOrderBook: make([]BinanceOrderBookEntry, 0),
		orderBookPublisher:         orderBookPublisher,
	}

	go instance.processIncomingOrderBookEntries()

	return &instance, nil
}

func (oba *orderBookAggregator) UpdateOrderBook(
	symbol string,
	orderBook []BinanceOrderBookEntry) {
	oba.orderBookChannel <- orderBookInput{
		symbol:    symbol,
		orderBook: orderBook,
	}
}

func (oba *orderBookAggregator) processIncomingOrderBookEntries() {
	for {
		select {
		case <-oba.ctx.Done():
		case orderBookInput := <-oba.orderBookChannel:
			oba.aggregatedBinanceOrderBook = slices.Filter(oba.aggregatedBinanceOrderBook, func(orderBookEntry BinanceOrderBookEntry) bool {
				return orderBookEntry.Symbol != orderBookInput.symbol
			})

			oba.aggregatedBinanceOrderBook = slices.Concat(oba.aggregatedBinanceOrderBook, orderBookInput.orderBook)

			sort.SliceStable(oba.aggregatedBinanceOrderBook, func(i, j int) bool {
				return oba.aggregatedBinanceOrderBook[i].Time.Before(oba.aggregatedBinanceOrderBook[j].Time)
			})
		}

		if oba.ctx.Err() == context.Canceled {
			break
		}
	}

}
