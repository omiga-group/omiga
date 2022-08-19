package subscribers

import (
	"context"
	"encoding/json"
	"math"
	"time"

	"github.com/gorilla/websocket"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/configuration"
	"go.uber.org/zap"
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
	Time     FtxTime     `json:"time"`
	Checksum int         `json:"checksum"`
	Bids     [][]float64 `json:"bids"`
	Asks     [][]float64 `json:"asks"`
	Action   string      `json:"action"`
}

type ftxOrderBookSubscriber struct {
	logger      *zap.SugaredLogger
	market      string
	ftxSettings configuration.FtxSettings
}

func NewFtxOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	ftxSettings configuration.FtxSettings,
	market string) (FtxOrderBookSubscriber, error) {
	instance := &ftxOrderBookSubscriber{
		logger:      logger,
		market:      market,
		ftxSettings: ftxSettings,
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
	connection, _, err := websocket.DefaultDialer.DialContext(
		ctx,
		fobs.ftxSettings.WebsocketUrl,
		nil)
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
	if err := connection.WriteJSON(&ftxRequest{
		Op:      OperationTypeSubscribe,
		Channel: &channel,
		Market:  &fobs.market,
	}); err != nil {
		fobs.logger.Errorf("Failed to send request to FTX websocket. Error: %v", err)

		return
	}

	go fobs.ping(ctx, connection)

	for {
		if ctx.Err() == context.Canceled {
			break
		}

		var orderBook ftxOrderBook

		err := connection.ReadJSON(&orderBook)
		if err != nil {
			fobs.logger.Errorf("Failed to read OrderBook JSON. Error: %v", err)

			break
		}
	}

	if connectionCloseErr := connection.Close(); connectionCloseErr != nil {
		fobs.logger.Errorf("Failed to close FTX websocket connection. Error: %v", connectionCloseErr)
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
