package subscribers

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/client"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
)

type GeminiOrderBookSubscriber interface {
	Close()
}

type geminiOrderBookSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	geminiConfig       configuration.GeminiConfig
	orderBookPublisher publishers.OrderBookPublisher
	apiClient          client.ApiClient
}

func NewGeminiOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	apiClient client.ApiClient,
	geminiConfig configuration.GeminiConfig,
	orderBookPublisher publishers.OrderBookPublisher) (GeminiOrderBookSubscriber, error) {

	instance := &geminiOrderBookSubscriber{
		ctx:                ctx,
		logger:             logger,
		geminiConfig:       geminiConfig,
		orderBookPublisher: orderBookPublisher,
		apiClient:          apiClient,
	}

	go instance.run()

	return instance, nil
}

func (gobs *geminiOrderBookSubscriber) Close() {
	gobs.orderBookPublisher.Close()
}

func (gobs *geminiOrderBookSubscriber) run() {
	for {
		gobs.connectAndSubscribe()

		if gobs.ctx.Err() == context.Canceled {
			return
		}
	}
}

func (gobs *geminiOrderBookSubscriber) connectAndSubscribe() {
	connection, _, err := websocket.DefaultDialer.DialContext(gobs.ctx, gobs.geminiConfig.WebsocketUrl, nil)
	if err != nil {
		gobs.logger.Errorf("Failed to dial Gemini websocket. Error: %v", err)
		return
	}

	if connection == nil {
		gobs.logger.Error("websocket is not initialized")
		return
	}

	defer func() {
		if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
			gobs.logger.Errorf("Failed to close Gemini websocket connection. Error: %v", connectionCloseErr)
		}
	}()

	mm, err := gobs.apiClient.GetMarkets()
	if err != nil {
		gobs.logger.Errorf("Failed to get Gemini markets list. Error: %v", err)
		return
	}

	channel := "orderbook"
	for name := range mm {
		req := &geminiRequest{Op: OperationTypeSubscribe, Channel: &channel, Market: string(name)}
		if err := connection.WriteJSON(req); err != nil {
			gobs.logger.Errorf("Failed to send request to Gemini websocket. Error: %v", err)
			return
		}
	}

	go gobs.ping(connection)

	for {
		if gobs.ctx.Err() == context.Canceled {
			break
		}

		orderBook := geminiOrderBook{}
		err := connection.ReadJSON(&orderBook)
		if err != nil {
			gobs.logger.Errorf("Failed to read OrderBook JSON. Error: %v", err)
			break
		}

		if orderBook.Data != nil {
			if market, ok := mm[models.MarketName(orderBook.Market)]; ok {
				gobs.publish(&orderBook, market)
			}
		}

	}

	if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
		gobs.logger.Errorf("Failed to close Gemini websocket connection. Error: %v", connectionCloseErr)
	}
}

func (gobs *geminiOrderBookSubscriber) publish(
	ob *geminiOrderBook,
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

	if err := gobs.orderBookPublisher.Publish(gobs.ctx, orderBook.ExchangeId, orderBook); err != nil {
		gobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)
	}
}

func (gobs *geminiOrderBookSubscriber) ping(
	connection *websocket.Conn) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := connection.WriteJSON(&geminiRequest{Op: "ping"}); err != nil {
				gobs.logger.Errorf("Failed to send ping request. Error: %v", err)
				return
			}
		case <-gobs.ctx.Done():
		}

		if gobs.ctx.Err() == context.Canceled {
			break
		}
	}
}
