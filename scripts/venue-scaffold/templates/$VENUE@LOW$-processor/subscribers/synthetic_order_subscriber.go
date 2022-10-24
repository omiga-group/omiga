package subscribers

import (
	"context"
	"encoding/json"

	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"go.uber.org/zap"
)

type syntheticOrderSubscriber struct {
	logger *zap.SugaredLogger
}

func NewSyntheticOrderSubscriber(logger *zap.SugaredLogger) (syntheticorderv1.Subscriber, error) {
	return &syntheticOrderSubscriber{
		logger: logger,
	}, nil
}

func (os *syntheticOrderSubscriber) Handle(ctx context.Context, event syntheticorderv1.SyntheticOrderEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	os.logger.Infof("Processing SyntheticOrderEvent event: %s", string(data))

	return nil
}
