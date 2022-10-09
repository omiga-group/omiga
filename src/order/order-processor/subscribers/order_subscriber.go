package subscribers

import (
	"context"
	"encoding/json"

	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"go.uber.org/zap"
)

type orderSubscriber struct {
	logger *zap.SugaredLogger
}

func NewOrderSubscriber(logger *zap.SugaredLogger) (orderv1.Subscriber, error) {
	return &orderSubscriber{
		logger: logger,
	}, nil
}

func (os *orderSubscriber) Handle(ctx context.Context, event orderv1.OrderEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	os.logger.Infof("Processing OrderEvent event: %s", string(data))

	return nil
}
