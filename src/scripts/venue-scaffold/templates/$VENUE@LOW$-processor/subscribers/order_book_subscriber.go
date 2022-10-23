package subscribers

import (
	"context"
	"strings"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/services"
	"go.uber.org/zap"
)

type $VENUE@PAS$OrderBookSubscriber interface {
	Close()
}

type $VENUE@LOW$OrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	pair               string
	orderBookPublisher publishers.OrderBookPublisher
	coinHelper         services.CurrencyHelper
	symbol1            string
	symbol2            string
}

func New$VENUE@PAS$OrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pairConfig configuration.PairConfig,
	orderBookPublisher publishers.OrderBookPublisher,
	coinHelper services.CurrencyHelper) ($VENUE@PAS$OrderBookSubscriber, error) {

	pairs := strings.Split(pairConfig.Pair, "/")
	symbol1 := strings.ToLower(pairs[0])
	symbol2 := strings.ToLower(pairs[1])

	instance := &$VENUE@LOW$OrderBookSubscriber{
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

func (bobs *$VENUE@LOW$OrderBookSubscriber) Close() {
	bobs.orderBookPublisher.Close()
}

func (bobs *$VENUE@LOW$OrderBookSubscriber) run() {
	for {
		bobs.connectAndSubscribe()

		if bobs.ctx.Err() == context.Canceled {
			return
		}
	}
}

func (bobs *$VENUE@LOW$OrderBookSubscriber) connectAndSubscribe() {
	_, stopChannel, err := $VENUE@LOW$.WsDepthServe100Ms(
		bobs.pair,
		bobs.wsDepthHandler,
		bobs.wsErrorHandler)
	if err != nil {
		bobs.logger.Errorf("Failed to connect to $VENUE@PAS$ websocket. Error: %v", err)

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

func (bobs *$VENUE@LOW$OrderBookSubscriber) wsDepthHandler(event *$VENUE@LOW$.WsDepthEvent) {
	if event == nil {
		bobs.logger.Warnf(
			"$VENUE@PAS$ websocket returned nil event for pair %s",
			bobs.pair)

		return
	}

	entryTime := time.UnixMilli(event.Time)

	asks := slices.Map(event.Asks, func(ask $VENUE@LOW$.Ask) models.$VENUE@PAS$OrderBookEntry {
		return models.$VENUE@PAS$OrderBookEntry{
			Symbol: bobs.pair,
			Time:   entryTime,
			Ask:    &ask,
			Bid:    nil,
		}
	})

	bids := slices.Map(event.Bids, func(bid $VENUE@LOW$.Bid) models.$VENUE@PAS$OrderBookEntry {
		return models.$VENUE@PAS$OrderBookEntry{
			Symbol: bobs.pair,
			Time:   entryTime,
			Ask:    nil,
			Bid:    &bid,
		}
	})

	$VENUE@LOW$OrderBook := slices.Concat(asks, bids)

	coins, err := bobs.coinHelper.GetCoinsNames(bobs.ctx, []string{bobs.symbol1, bobs.symbol2})
	if err != nil {
		bobs.logger.Errorf("Failed to fetch coin names. Error: %v", err)

		return
	}

	baseCoinName := coins[bobs.symbol1]
	counterCoinName := coins[bobs.symbol2]

	orderBook := mappers.$VENUE@PAS$OrderBookToOrderBook(
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
		$VENUE@LOW$OrderBook,
	)

	orderBook.ExchangeId = "$VENUE@LOW$"

	if err := bobs.orderBookPublisher.Publish(
		bobs.ctx,
		orderBook.ExchangeId,
		orderBook); err != nil {
		bobs.logger.Errorf("Failed to publish order book for $VENUE@PAS$ exchange. Error: %v", err)

		return
	}
}

func (bobs *$VENUE@LOW$OrderBookSubscriber) wsErrorHandler(err error) {
	bobs.logger.Errorf(
		"$VENUE@PAS$ websocket returned error for pair %s. Error: %v",
		bobs.pair,
		err)
}
