package publishers

import (
	"go.uber.org/zap"
)

type OrderPublisher interface {
}

type orderPublisher struct {
	logger *zap.SugaredLogger
}

func NewOrderPublisher(
	logger *zap.SugaredLogger) (OrderPublisher, error) {
	return &orderPublisher{
		logger: logger,
	}, nil
}
