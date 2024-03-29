package subscribers

import (
	"context"
	"strings"

	krakenwebsocket "github.com/aopoltorzhicky/go_kraken/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/kraken-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/kraken-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/kraken-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/services"
	"go.uber.org/zap"
)

type KrakenOrderBookSubscriber interface {
	Close()
}

type krakenOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	pairs              []configuration.PairConfig
	orderBookPublisher publishers.OrderBookPublisher
	coinHelper         services.CurrencyHelper
}

func NewKrakenOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.KrakenConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	coinHelper services.CurrencyHelper) (KrakenOrderBookSubscriber, error) {

	instance := &krakenOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		pairs:              venueConfig.OrderBook.Pairs,
		orderBookPublisher: orderBookPublisher,
		coinHelper:         coinHelper,
	}

	go instance.run()

	return instance, nil
}

func (kobs *krakenOrderBookSubscriber) Close() {
	kobs.orderBookPublisher.Close()
}

func (kobs *krakenOrderBookSubscriber) run() {
	for {
		kobs.connectAndSubscribe()

		if kobs.ctx.Err() == context.Canceled {
			return
		}
	}
}

func (kobs *krakenOrderBookSubscriber) connectAndSubscribe() {
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
				asks := slices.Map(data.Asks, func(ask krakenwebsocket.OrderBookItem) models.KrakenOrderBookEntry {
					return models.KrakenOrderBookEntry{
						Symbol: update.Pair,
						Ask:    &ask,
						Bid:    nil,
					}
				})

				bids := slices.Map(data.Bids, func(bid krakenwebsocket.OrderBookItem) models.KrakenOrderBookEntry {
					return models.KrakenOrderBookEntry{
						Symbol: update.Pair,
						Ask:    nil,
						Bid:    &bid,
					}
				})

				krakenOrderBook := slices.Concat(asks, bids)
				pairs := strings.Split(update.Pair, "/")
				symbol1 := strings.ToLower(pairs[0])
				symbol2 := strings.ToLower(pairs[1])

				coins, err := kobs.coinHelper.GetCoinsNames(kobs.ctx, []string{symbol1, symbol2})
				if err != nil {
					kobs.logger.Errorf("Failed to fetch coin names. Error: %v", err)

					return
				}

				baseCoinName := coins[symbol1]
				counterCoinName := coins[symbol2]

				orderBook := mappers.KrakenOrderBookToOrderBook(
					exchangeModels.OrderCurrency{
						Name:         baseCoinName,
						Code:         pairs[0],
						MaxPrecision: 1,
						Digital:      true,
					},
					exchangeModels.OrderCurrency{
						Name:         counterCoinName,
						Code:         pairs[1],
						MaxPrecision: 1,
						Digital:      true,
					},
					krakenOrderBook,
				)

				orderBook.ExchangeId = "kraken"

				if err := kobs.orderBookPublisher.Publish(
					kobs.ctx,
					orderBook.ExchangeId,
					orderBook); err != nil {
					kobs.logger.Errorf("Failed to publish order book for Kraken venue. Error: %v", err)

					return
				}

			default:
			}
		}

		if kobs.ctx.Err() == context.Canceled {
			return
		}
	}

}
