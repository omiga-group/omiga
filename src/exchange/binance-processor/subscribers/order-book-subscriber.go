package subscribers

import (
	"context"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/exchange/shared/services"
	"go.uber.org/zap"
)

type BinanceOrderBookSubscriber interface {
}

type binanceOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	symbolConfig       configuration.SymbolConfig
	symbol             string
	orderBookPublisher publishers.OrderBookPublisher
	coinHelper         services.CoinHelper
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	symbolConfig configuration.SymbolConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	coinHelper services.CoinHelper) (BinanceOrderBookSubscriber, error) {

	binance.UseTestnet = binanceConfig.UseTestnet

	instance := &binanceOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		symbolConfig:       symbolConfig,
		orderBookPublisher: orderBookPublisher,
		coinHelper:         coinHelper,
		symbol:             symbolConfig.Symbol1 + symbolConfig.Symbol2,
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
			bobs.symbolConfig)

		return
	}

	entryTime := time.UnixMilli(event.Time)

	asks := slices.Map(event.Asks, func(ask binance.Ask) models.BinanceOrderBookEntry {
		return models.BinanceOrderBookEntry{
			Symbol: bobs.symbol,
			Time:   entryTime,
			Ask:    &ask,
			Bid:    nil,
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

	binanceOrderBook := slices.Concat(asks, bids)

	symbol1 := strings.ToLower(bobs.symbolConfig.Symbol1)
	symbol2 := strings.ToLower(bobs.symbolConfig.Symbol2)
	coins, err := bobs.coinHelper.GetCoinsNames(bobs.ctx, []string{symbol1, symbol2})
	if err != nil {
		bobs.logger.Errorf("Failed to fetch coin names. Error: %v", err)

		return
	}

	baseCoinName := coins[symbol1]
	counterCoinName := coins[symbol2]

	orderBook := mappers.FromBinanceOrderBookToOrderBook(
		exchangeModels.Currency{
			Name:         baseCoinName,
			Code:         bobs.symbolConfig.Symbol1,
			MaxPrecision: 1,
			Digital:      true,
		},
		exchangeModels.Currency{
			Name:         counterCoinName,
			Code:         bobs.symbolConfig.Symbol2,
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
		"Binance websocket returned error for symbol %s. Error: %v",
		bobs.symbolConfig,
		err)
}
