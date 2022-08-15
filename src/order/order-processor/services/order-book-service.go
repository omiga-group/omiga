package services

import (
	"context"
	"time"

	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/order/shared/repositories/orderbook"
	"go.uber.org/zap"
)

type OrderBookService interface {
	Handle(
		ctx context.Context,
		exchangeId string,
		orderbookTimestamp time.Time,
		orderBook models.OrderBook) error
}

type orderBookService struct {
	logger      *zap.SugaredLogger
	entgoClient repositories.EntgoClient
}

func NewOrderBookService(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient) (OrderBookService, error) {
	return &orderBookService{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (obs *orderBookService) Handle(
	ctx context.Context,
	exchangeId string,
	orderbookTimestamp time.Time,
	orderBook models.OrderBook) error {
	client := obs.entgoClient.GetClient()
	orderBookEntry, err := client.OrderBook.
		Query().
		Where(orderbook.ExchangeIDEQ(exchangeId)).
		First(ctx)

	if _, ok := err.(*repositories.NotFoundError); ok {
		if _, err = client.OrderBook.
			Create().
			SetExchangeID(exchangeId).
			SetLastUpdated(orderbookTimestamp).
			SetOrderBook(orderBook).
			Save(ctx); err != nil {
			obs.logger.Errorf("Failed to save order book for exchange with ID: %s. Error: %v", exchangeId, err)

			return err
		}

		return nil
	} else if err != nil {
		obs.logger.Errorf("Failed to fetch current order book for exchange with ID: %s. Error: %v", exchangeId, err)

		return err
	}

	if orderBookEntry.LastUpdated.After(orderbookTimestamp) {
		obs.logger.Infof("Discarding order book for exchange with ID: %s.", exchangeId)

		return nil
	}

	if _, err = client.OrderBook.
		UpdateOne(orderBookEntry).
		SetLastUpdated(orderbookTimestamp).
		SetOrderBook(orderBook).
		Save(ctx); err != nil {
		obs.logger.Errorf("Failed to save order book for exchange with ID: %s. Error: %v", exchangeId, err)

		return err
	}

	return nil
}
