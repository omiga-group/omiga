package subscribers

import (
	"context"
	"sort"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/models"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/services"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"go.uber.org/zap"
)

type BinanceOrderBookSubscriber interface {
}

type binanceOrderBookSubscriber struct {
	ctx                                                          context.Context
	logger                                                       *zap.SugaredLogger
	symbol                                                       string
	purgeTime                                                    time.Duration
	orderBookPublisher                                           publishers.OrderBookPublisher
	binanceOrderBook                                             []models.BinanceOrderBookEntry
	baseCoinCode, baseCoinName, counterCoinCode, counterCoinName string
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	symbolConfig configuration.SymbolConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	symbolEnricher services.SymbolEnricher) (BinanceOrderBookSubscriber, error) {

	binance.UseTestnet = binanceConfig.UseTestnet

	purgeTime, err := time.ParseDuration(symbolConfig.PurgeTime)
	if err != nil {
		return nil, err
	}

	baseCoinCode, baseCoinName, counterCoinCode, counterCoinName, err := symbolEnricher.GetCoinPair(symbolConfig.Symbol)
	if err != nil {
		return nil, err
	}

	instance := &binanceOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		symbol:             symbolConfig.Symbol,
		binanceOrderBook:   make([]models.BinanceOrderBookEntry, 0),
		purgeTime:          purgeTime,
		orderBookPublisher: orderBookPublisher,
		baseCoinCode:       baseCoinCode,
		baseCoinName:       baseCoinName,
		counterCoinCode:    counterCoinCode,
		counterCoinName:    counterCoinName,
	}

	go instance.run()

	return instance, nil
}

func (bobs *binanceOrderBookSubscriber) run() {
	for {
		bobs.connectAndSubscribe()

		if bobs.ctx.Err() == context.Canceled {
			return
		}
	}
}

func (bobs *binanceOrderBookSubscriber) connectAndSubscribe() {
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

		if bobs.ctx.Err() == context.Canceled {
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

	asks := slices.Map(event.Asks, func(ask binance.Ask) models.BinanceOrderBookEntry {
		return models.BinanceOrderBookEntry{
			Time: entryTime,
			Ask:  &ask,
			Bid:  nil,
		}
	})

	bids := slices.Map(event.Bids, func(bid binance.Bid) models.BinanceOrderBookEntry {
		return models.BinanceOrderBookEntry{
			Symbol: bobs.symbol,
			Time:   entryTime,
			Ask:    nil,
			Bid:    &bid,
		}
	})

	bobs.binanceOrderBook = slices.Concat(bobs.binanceOrderBook, asks, bids)

	purgeTime := time.Now().Add(-1 * bobs.purgeTime)

	bobs.binanceOrderBook = slices.Filter(bobs.binanceOrderBook, func(orderBookEntry models.BinanceOrderBookEntry) bool {
		return orderBookEntry.Time.After(purgeTime)
	})

	sort.SliceStable(bobs.binanceOrderBook, func(i, j int) bool {
		return bobs.binanceOrderBook[i].Time.Before(bobs.binanceOrderBook[j].Time)
	})

	orderBook := mappers.FromBinanceOrderBookToModelOrderBook(
		exchangeModels.Currency{
			Name:         bobs.baseCoinName,
			Code:         bobs.baseCoinCode,
			MaxPrecision: 1,
			Digital:      true,
		},
		exchangeModels.Currency{
			Name:         bobs.counterCoinName,
			Code:         bobs.counterCoinCode,
			MaxPrecision: 1,
			Digital:      true,
		},
		bobs.binanceOrderBook,
	)

	if err := bobs.orderBookPublisher.Publish(
		bobs.ctx,
		"binance",
		orderBook); err != nil {
		bobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)
	}
}

func (bobs *binanceOrderBookSubscriber) wsErrorHandler(err error) {
	bobs.logger.Errorf(
		"Binance websocket returned error for symbol %s. Error: %v",
		bobs.symbol,
		err)
}
