package subscribers

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"go.uber.org/zap"
)

type BinanceOrderBookSubscriber interface {
}

type binanceOrderBookSubscriber struct {
	logger *zap.SugaredLogger
	symbol string
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	symbol string) (BinanceOrderBookSubscriber, error) {

	binance.UseTestnet = binanceConfig.UseTestnet

	instance := &binanceOrderBookSubscriber{
		logger: logger,
		symbol: symbol,
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

	bobs.logger.Info(*event)
}

func (bobs *binanceOrderBookSubscriber) wsErrorHandler(err error) {
	bobs.logger.Errorf(
		"Binance websocket returned error for symbol %s. Error: %v",
		bobs.symbol,
		err)
}
