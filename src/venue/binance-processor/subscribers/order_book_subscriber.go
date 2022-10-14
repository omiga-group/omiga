package subscribers

import (
	"context"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/binance-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/binance-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/services"
	"go.uber.org/zap"
)

type BinanceOrderBookSubscriber interface {
	Close()
}

type binanceOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	pair               string
	orderBookPublisher publishers.OrderBookPublisher
	coinHelper         services.CurrencyHelper
	symbol1            string
	symbol2            string
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pairConfig configuration.PairConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	coinHelper services.CurrencyHelper) (BinanceOrderBookSubscriber, error) {

	pairs := strings.Split(pairConfig.Pair, "/")
	symbol1 := strings.ToLower(pairs[0])
	symbol2 := strings.ToLower(pairs[1])

	instance := &binanceOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		orderBookPublisher: orderBookPublisher,
		coinHelper:         coinHelper,
		pair:               strings.Replace(pairConfig.Pair, "/", "", -1),
		symbol1:            symbol1,
		symbol2:            symbol2,
	}

	go instance.run()

	return instance, nil
}

func (bobs *binanceOrderBookSubscriber) Close() {
	bobs.orderBookPublisher.Close()
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
		bobs.pair,
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
			"Binance websocket returned nil event for pair %s",
			bobs.pair)

		return
	}

	entryTime := time.UnixMilli(event.Time)

	asks := slices.Map(event.Asks, func(ask binance.Ask) models.BinanceOrderBookEntry {
		return models.BinanceOrderBookEntry{
			Symbol: bobs.pair,
			Time:   entryTime,
			Ask:    &ask,
			Bid:    nil,
		}
	})

	bids := slices.Map(event.Bids, func(bid binance.Bid) models.BinanceOrderBookEntry {
		return models.BinanceOrderBookEntry{
			Symbol: bobs.pair,
			Time:   entryTime,
			Ask:    nil,
			Bid:    &bid,
		}
	})

	binanceOrderBook := slices.Concat(asks, bids)

	coins, err := bobs.coinHelper.GetCoinsNames(bobs.ctx, []string{bobs.symbol1, bobs.symbol2})
	if err != nil {
		bobs.logger.Errorf("Failed to fetch coin names. Error: %v", err)

		return
	}

	baseCoinName := coins[bobs.symbol1]
	counterCoinName := coins[bobs.symbol2]

	orderBook := mappers.BinanceOrderBookToOrderBook(
		exchangeModels.OrderCurrency{
			Name:         baseCoinName,
			Code:         bobs.symbol1,
			MaxPrecision: 1,
			Digital:      true,
		},
		exchangeModels.OrderCurrency{
			Name:         counterCoinName,
			Code:         bobs.symbol2,
			MaxPrecision: 1,
			Digital:      true,
		},
		binanceOrderBook,
	)

	orderBook.ExchangeId = "binance"

	if err := bobs.orderBookPublisher.Publish(
		bobs.ctx,
		orderBook.ExchangeId,
		orderBook); err != nil {
		bobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)

		return
	}
}

func (bobs *binanceOrderBookSubscriber) wsErrorHandler(err error) {
	bobs.logger.Errorf(
		"Binance websocket returned error for pair %s. Error: %v",
		bobs.pair,
		err)
}
