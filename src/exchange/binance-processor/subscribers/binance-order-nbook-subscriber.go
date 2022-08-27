package subscribers

import (
	"context"
	"sort"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"go.uber.org/zap"
)

type BinanceOrderBookSubscriber interface {
}

type BinanceOrderBookEntry struct {
	Time time.Time
	Bid  *binance.Bid
	Ask  *binance.Ask
}

type binanceOrderBookSubscriber struct {
	logger           *zap.SugaredLogger
	symbol           string
	purgeTime        time.Duration
	binanceOrderBook []BinanceOrderBookEntry
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	symbolConfig configuration.SymbolConfig) (BinanceOrderBookSubscriber, error) {

	binance.UseTestnet = binanceConfig.UseTestnet

	purgeTime, err := time.ParseDuration(symbolConfig.PurgeTime)
	if err != nil {
		return nil, err
	}

	instance := &binanceOrderBookSubscriber{
		logger:           logger,
		symbol:           symbolConfig.Symbol,
		binanceOrderBook: make([]BinanceOrderBookEntry, 0),
		purgeTime:        purgeTime,
	}

	go instance.run(ctx)

	return instance, nil
}

func (bobs *binanceOrderBookSubscriber) run(ctx context.Context) {
	for {
		bobs.connectAndSubscribe(ctx)

		if ctx.Err() == context.Canceled {
			return
		}
	}
}

func (bobs *binanceOrderBookSubscriber) connectAndSubscribe(ctx context.Context) {
	_, stopChannel, err := binance.WsDepthServe100Ms(
		bobs.symbol,
		bobs.wsDepthHandler,
		bobs.wsErrorHandler)
	if err != nil {
		bobs.logger.Errorf("Failed to connect to Binance websocket. Error: %v", err)

		return
	}

	for {
		time.Sleep(1 * time.Second)

		if ctx.Err() == context.Canceled {
			stopChannel <- struct{}{}

			return
		}
	}
}

func (bobs *binanceOrderBookSubscriber) wsDepthHandler(event *binance.WsDepthEvent) {
	if event == nil {
		bobs.logger.Warnf(
			"Binance websocket returned nil event for symbol %s",
			bobs.symbol)

		return
	}

	entryTime := time.UnixMilli(event.Time)

	asks := slices.Map(event.Asks, func(ask binance.Ask) BinanceOrderBookEntry {
		return BinanceOrderBookEntry{
			Time: entryTime,
			Ask:  &ask,
			Bid:  nil,
		}
	})

	bids := slices.Map(event.Bids, func(bid binance.Bid) BinanceOrderBookEntry {
		return BinanceOrderBookEntry{
			Time: entryTime,
			Ask:  nil,
			Bid:  &bid,
		}
	})

	bobs.binanceOrderBook = slices.Concat(bobs.binanceOrderBook, asks, bids)

	purgeTime := time.Now().Add(-1 * bobs.purgeTime)

	bobs.binanceOrderBook = slices.Filter(bobs.binanceOrderBook, func(orderBookEntry BinanceOrderBookEntry) bool {
		return orderBookEntry.Time.After(purgeTime)
	})

	sort.SliceStable(bobs.binanceOrderBook, func(i, j int) bool {
		return bobs.binanceOrderBook[i].Time.Before(bobs.binanceOrderBook[j].Time)
	})
}

func (bobs *binanceOrderBookSubscriber) wsErrorHandler(err error) {
	bobs.logger.Errorf(
		"Binance websocket returned error for symbol %s. Error: %v",
		bobs.symbol,
		err)
}
