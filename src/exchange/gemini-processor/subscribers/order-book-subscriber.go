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
}

type geminiOrderBookSubscriber struct {
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
		logger:             logger,
		geminiConfig:       geminiConfig,
		orderBookPublisher: orderBookPublisher,
		apiClient:          apiClient,
	}

	go instance.run(ctx)

	return instance, nil
}

func (fobs *geminiOrderBookSubscriber) run(ctx context.Context) {
	for {
		fobs.connectAndSubscribe(ctx)

		if ctx.Err() == context.Canceled {
			return
		}
	}
}

func (fobs *geminiOrderBookSubscriber) connectAndSubscribe(ctx context.Context) {
	connection, _, err := websocket.DefaultDialer.DialContext(ctx, fobs.geminiConfig.WebsocketUrl, nil)
	if err != nil {
		fobs.logger.Errorf("Failed to dial Gemini websocket. Error: %v", err)
		return
	}

	if connection == nil {
		fobs.logger.Error("websocket is not initialized")
		return
	}

	defer func() {
		if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
			fobs.logger.Errorf("Failed to close Gemini websocket connection. Error: %v", connectionCloseErr)
		}
	}()

	mm, err := fobs.apiClient.GetMarkets()
	if err != nil {
		fobs.logger.Errorf("Failed to get Gemini markets list. Error: %v", err)
		return
	}

	channel := "orderbook"
	for name := range mm {
		req := &geminiRequest{Op: OperationTypeSubscribe, Channel: &channel, Market: string(name)}
		if err := connection.WriteJSON(req); err != nil {
			fobs.logger.Errorf("Failed to send request to Gemini websocket. Error: %v", err)
			return
		}
	}

	go fobs.ping(ctx, connection)

	for {
		if ctx.Err() == context.Canceled {
			break
		}

		orderBook := geminiOrderBook{}
		err := connection.ReadJSON(&orderBook)
		if err != nil {
			fobs.logger.Errorf("Failed to read OrderBook JSON. Error: %v", err)
			break
		}

		if orderBook.Data != nil {
			if market, ok := mm[models.MarketName(orderBook.Market)]; ok {
				fobs.publish(ctx, &orderBook, market)
			}
		}

	}

	if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
		fobs.logger.Errorf("Failed to close Gemini websocket connection. Error: %v", connectionCloseErr)
	}
}

func (fobs *geminiOrderBookSubscriber) publish(
	ctx context.Context,
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
		exchangeModels.Currency{
			Name:         market.BaseCurrency,
			Code:         market.BaseCurrency,
			MaxPrecision: 1,    //WHY SEND PRECISION?
			Digital:      true, //WHAT IS THIS?
		},
		exchangeModels.Currency{
			Name:         market.QuoteCurrency,
			Code:         market.QuoteCurrency,
			MaxPrecision: 1,
			Digital:      true,
		},
		asksBids,
	)

	orderBook.ExchangeId = "binance"

	if err := fobs.orderBookPublisher.Publish(ctx, orderBook.ExchangeId, orderBook); err != nil {
		fobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)
	}
}

func (fobs *geminiOrderBookSubscriber) ping(
	ctx context.Context,
	connection *websocket.Conn) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := connection.WriteJSON(&geminiRequest{Op: "ping"}); err != nil {
				fobs.logger.Errorf("Failed to send ping request. Error: %v", err)
				return
			}
		case <-ctx.Done():
		}

		if ctx.Err() == context.Canceled {
			break
		}
	}
}
