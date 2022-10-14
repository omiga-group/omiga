package subscribers

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/client"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
)

type FtxOrderBookSubscriber interface {
	Close()
}

type ftxOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	ftxConfig          configuration.FtxConfig
	orderBookPublisher publishers.OrderBookPublisher
	apiClient          client.ApiClient
}

func NewFtxOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	apiClient client.ApiClient,
	ftxConfig configuration.FtxConfig,
	orderBookPublisher publishers.OrderBookPublisher) (FtxOrderBookSubscriber, error) {

	instance := &ftxOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		ftxConfig:          ftxConfig,
		orderBookPublisher: orderBookPublisher,
		apiClient:          apiClient,
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
	connection, _, err := websocket.DefaultDialer.DialContext(fobs.ctx, fobs.ftxConfig.WebsocketUrl, nil)
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

	mm, err := fobs.apiClient.GetMarkets()
	if err != nil {
		fobs.logger.Errorf("Failed to get FTX markets list. Error: %v", err)
		return
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
	market models.Market) {
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
		fobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)
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
