package subscribers

import (
	"context"

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
	os.logger.Infof("Processing OrderEvent event: %v", event)

	return nil
}
