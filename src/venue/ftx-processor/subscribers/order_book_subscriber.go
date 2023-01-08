package subscribers

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	ftxv1 "github.com/omiga-group/omiga/src/venue/ftx-processor/ftxclient/v1"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
)

type FtxOrderBookSubscriber interface {
	Close()
}

type ftxOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	venueConfig        configuration.FtxConfig
	orderBookPublisher publishers.OrderBookPublisher
}

func NewFtxOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.FtxConfig,
	orderBookPublisher publishers.OrderBookPublisher) (FtxOrderBookSubscriber, error) {

	instance := &ftxOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		venueConfig:        venueConfig,
		orderBookPublisher: orderBookPublisher,
	}

	go instance.run()

	return instance, nil
}

func (fobs *ftxOrderBookSubscriber) Close() {
	fobs.orderBookPublisher.Close()
}

func (fobs *ftxOrderBookSubscriber) run() {
	for {
		fobs.connectAndSubscribe()

		if fobs.ctx.Err() == context.Canceled {
			return
		}
	}
}

func (fobs *ftxOrderBookSubscriber) connectAndSubscribe() {
	connection, _, err := websocket.DefaultDialer.DialContext(fobs.ctx, fobs.venueConfig.WebsocketUrl, nil)
	if err != nil {
		fobs.logger.Errorf("Failed to dial FTX websocket. Error: %v", err)
		return
	}

	if connection == nil {
		fobs.logger.Error("websocket is not initialized")
		return
	}

	defer func() {
		if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
			fobs.logger.Errorf("Failed to close FTX websocket connection. Error: %v", connectionCloseErr)
		}
	}()

	client, err := ftxv1.NewClientWithResponses(fobs.venueConfig.ApiUrl)
	if err != nil {
		fobs.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetMarketsWithResponse(fobs.ctx)
	if err != nil {
		fobs.logger.Errorf("Failed to call getMarkets endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		fobs.logger.Errorf("Failed to call getMarkets endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		fobs.logger.Errorf("Returned JSON object is nil")

		return
	}

	rs := *response.JSON200.Result
	mm := models.MarketsMap{}
	for _, m := range rs {
		if !m.Enabled || m.Type != string(models.MarketTypeSpot) {
			continue
		}

		mm[m.Name] = m
	}

	channel := "orderbook"
	for name := range mm {
		req := &ftxRequest{Op: OperationTypeSubscribe, Channel: &channel, Market: name}
		if err := connection.WriteJSON(req); err != nil {
			fobs.logger.Errorf("Failed to send request to FTX websocket. Error: %v", err)
			return
		}
	}

	go fobs.ping(connection)

	for {
		if fobs.ctx.Err() == context.Canceled {
			break
		}

		orderBook := ftxOrderBook{}
		err := connection.ReadJSON(&orderBook)
		if err != nil {
			fobs.logger.Errorf("Failed to read OrderBook JSON. Error: %v", err)
			break
		}

		if orderBook.Data != nil {
			if market, ok := mm[orderBook.Market]; ok {
				fobs.publish(&orderBook, market)
			}
		}

	}

	if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
		fobs.logger.Errorf("Failed to close FTX websocket connection. Error: %v", connectionCloseErr)
	}
}

func (fobs *ftxOrderBookSubscriber) publish(
	ob *ftxOrderBook,
	market ftxv1.Market) {
	asks := slices.Map(ob.Data.Asks, func(ask [2]float64) models.OrderBookEntry {
		return models.OrderBookEntry{
			Symbol: ob.Market,
			Time:   ob.Data.Time.Time,
			Ask:    &models.PriceLevel{Price: ask[0], Quantity: ask[1]},
			Bid:    nil,
		}
	})

	bids := slices.Map(ob.Data.Bids, func(bid [2]float64) models.OrderBookEntry {
		return models.OrderBookEntry{
			Symbol: ob.Market,
			Time:   ob.Data.Time.Time,
			Ask:    nil,
			Bid:    &models.PriceLevel{Price: bid[0], Quantity: bid[1]},
		}
	})

	asksBids := slices.Concat(asks, bids)

	orderBook := mappers.ToModelOrderBook(
		exchangeModels.OrderCurrency{
			Name:         market.BaseCurrency,
			Code:         market.BaseCurrency,
			MaxPrecision: 1,    //WHY SEND PRECISION?
			Digital:      true, //WHAT IS THIS?
		},
		exchangeModels.OrderCurrency{
			Name:         market.QuoteCurrency,
			Code:         market.QuoteCurrency,
			MaxPrecision: 1,
			Digital:      true,
		},
		asksBids,
	)

	orderBook.ExchangeId = "binance"

	if err := fobs.orderBookPublisher.Publish(fobs.ctx, orderBook.ExchangeId, orderBook); err != nil {
		fobs.logger.Errorf("Failed to publish order book for Binance venue. Error: %v", err)
	}
}

func (fobs *ftxOrderBookSubscriber) ping(
	connection *websocket.Conn) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := connection.WriteJSON(&ftxRequest{Op: "ping"}); err != nil {
				fobs.logger.Errorf("Failed to send ping request. Error: %v", err)
				return
			}
		case <-fobs.ctx.Done():
		}

		if fobs.ctx.Err() == context.Canceled {
			break
		}
	}
}
