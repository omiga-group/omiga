package subscribers

import (
	"context"
	"encoding/json"

	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"go.uber.org/zap"
)

type orderBookSubscriber struct {
	logger *zap.SugaredLogger
}

func NewOrderBookSubscriber(logger *zap.SugaredLogger) (orderbookv1.Subscriber, error) {
	return &orderBookSubscriber{
		logger: logger,
	}, nil
}

func (os *orderBookSubscriber) Handle(ctx context.Context, event orderbookv1.OrderBookEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	os.logger.Infof("Processing OrderBookEvent event: %s", string(data))

	return nil
}
