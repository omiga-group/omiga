package subscribers

import (
	"context"

	krakenwebsocket "github.com/aopoltorzhicky/go_kraken/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/kraken-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/exchange/shared/services"
	"go.uber.org/zap"
)

type KrakenOrderBookSubscriber interface {
}

type krakenOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	pairs              []configuration.PairConfig
	orderBookPublisher publishers.OrderBookPublisher
}

func NewKrakenOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	krakenConfig configuration.KrakenConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	symbolEnricher services.SymbolEnricher) (KrakenOrderBookSubscriber, error) {

	instance := &krakenOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		pairs:              krakenConfig.OrderBook.Pairs,
		orderBookPublisher: orderBookPublisher,
	}

	go instance.run()

	return instance, nil
}

func (kobs *krakenOrderBookSubscriber) run() {
	kraken := krakenwebsocket.NewKraken(krakenwebsocket.ProdBaseURL)
	if err := kraken.Connect(); err != nil {
		kobs.logger.Errorf("Error connecting to Kraken websocket. Error:  %v", err)

		return
	}

	pairs := slices.Map(kobs.pairs, func(pair configuration.PairConfig) string {
		return pair.Pair
	})

	if err := kraken.SubscribeBook(pairs, 100); err != nil {
		kobs.logger.Errorf("Failed to subscribe to Kraken websocket to receive order books. Error: %v", err)

		return
	}

	defer func() {
		err := kraken.UnsubscribeBook(pairs, 100)
		if err != nil {
			kobs.logger.Errorf("Failed to unsubscribe from Kraken orer book. Error: %v", err)
		}
	}()

	for {
		select {
		case <-kobs.ctx.Done():
			if err := kraken.Close(); err != nil {
				kobs.logger.Errorf("Failed to close Kraken websocket client. Error: %v", err)
			}
		case update := <-kraken.Listen():
			switch data := update.Data.(type) {
			case krakenwebsocket.OrderBookUpdate:
				kobs.logger.Infof("----Ticker of %s----", update.Pair)
				kobs.logger.Infof("----Ticker of %s----", data.Asks)
				kobs.logger.Infof("----Ticker of %s----", data.Bids)
			default:
			}
		}

		if kobs.ctx.Err() == context.Canceled {
			return
		}
	}

}
