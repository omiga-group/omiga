package subscribers

import (
	"context"
	"encoding/json"
	"math"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/models"
	exchangeModels "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
)

type FtxOrderBookSubscriber interface {
}

type FtxTime struct {
	Time time.Time
}

func (p *FtxTime) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	sec, nsec := math.Modf(f)
	p.Time = time.Unix(int64(sec), int64(nsec))
	return nil
}

type operationType string

const (
	OperationTypeSubscribe   operationType = "subscribe"
	OperationTypeUnsubscribe operationType = "unsubscribe"
)

type ftxRequest struct {
	Op      operationType `json:"op"`
	Channel *string       `json:"channel,omitempty"`
	Market  *string       `json:"market,omitempty"`
}

type responseType string

const (
	ResponseTypeError        responseType = "error"
	ResponseTypeSubscribed   responseType = "subscribed"
	ResponseTypeUnsubscribed responseType = "unsubscribed"
	ResponseTypeInfo         responseType = "info"
	ResponseTypePartial      responseType = "partial"
	ResponseTypeUpdate       responseType = "update"
)

type ftxOrderBook struct {
	Channel string            `json:"channel"`
	Market  string            `json:"market"`
	Type    responseType      `json:"type"`
	Data    *ftxOrderBookData `json:"data,omitempty"`
}

type ftxOrderBookData struct {
	Time     FtxTime      `json:"time"`
	Checksum int          `json:"checksum"`
	Bids     [][2]float64 `json:"bids"`
	Asks     [][2]float64 `json:"asks"`
	Action   string       `json:"action"`
}

type ftxOrderBookSubscriber struct {
	logger             *zap.SugaredLogger
	market             string
	ftxConfig          configuration.FtxConfig
	orderBookPublisher publishers.OrderBookPublisher
}

func NewFtxOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	ftxConfig configuration.FtxConfig,
	marketConfig configuration.MarketConfig) (FtxOrderBookSubscriber, error) {

	instance := &ftxOrderBookSubscriber{
		logger:    logger,
		market:    marketConfig.Market,
		ftxConfig: ftxConfig,
	}

	go instance.run(ctx)

	return instance, nil
}

func (fobs *ftxOrderBookSubscriber) run(ctx context.Context) {
	for {
		fobs.connectAndSubscribe(ctx)

		if ctx.Err() == context.Canceled {
			return
		}
	}
}

func (fobs *ftxOrderBookSubscriber) connectAndSubscribe(ctx context.Context) {
	connection, _, err := websocket.DefaultDialer.DialContext(ctx, fobs.ftxConfig.WebsocketUrl, nil)
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

	channel := "orderbook"
	req := &ftxRequest{Op: OperationTypeSubscribe, Channel: &channel, Market: &fobs.market}
	if err := connection.WriteJSON(req); err != nil {
		fobs.logger.Errorf("Failed to send request to FTX websocket. Error: %v", err)
		return
	}

	go fobs.ping(ctx, connection)

	for {
		if ctx.Err() == context.Canceled {
			break
		}

		orderBook := ftxOrderBook{}
		err := connection.ReadJSON(&orderBook)
		if err != nil {
			fobs.logger.Errorf("Failed to read OrderBook JSON. Error: %v", err)
			break
		}

		fobs.publish(ctx, &orderBook)
	}

	if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
		fobs.logger.Errorf("Failed to close FTX websocket connection. Error: %v", connectionCloseErr)
	}
}

func (fobs *ftxOrderBookSubscriber) publish(ctx context.Context, ob *ftxOrderBook) {
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

	binanceOrderBook := slices.Concat(asks, bids)

	orderBook := mappers.ToModelOrderBook(
		exchangeModels.Currency{
			Name:         "TODO",
			Code:         "TODO",
			MaxPrecision: 1,
			Digital:      true,
		},
		exchangeModels.Currency{
			Name:         "TODO",
			Code:         "TODO",
			MaxPrecision: 1,
			Digital:      true,
		},
		binanceOrderBook,
	)

	orderBook.ExchangeId = "binance"

	if err := fobs.orderBookPublisher.Publish(ctx, orderBook.ExchangeId, orderBook); err != nil {
		fobs.logger.Errorf("Failed to publish order book for Binance exchange. Error: %v", err)
	}
}

func (fobs *ftxOrderBookSubscriber) ping(
	ctx context.Context,
	connection *websocket.Conn) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := connection.WriteJSON(&ftxRequest{
				Op: "ping",
			}); err != nil {
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
