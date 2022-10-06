package repositories

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	exchangerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpairs"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type TradingPairsRepository interface {
	CreateTradingPairs(
		ctx context.Context,
		exchangeId string,
		tradingPairs []models.TradingPairs) error
}

type tradingPairsRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewTradingPairsRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (TradingPairsRepository, error) {
	return &tradingPairsRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (tpr *tradingPairsRepository) CreateTradingPairs(
	ctx context.Context,
	exchangeId string,
	tradingPairs []models.TradingPairs) error {
	client := tpr.entgoClient.GetClient()

	err := client.Exchange.
		Create().
		SetExchangeID(exchangeId).
		OnConflictColumns(exchangerepo.FieldExchangeID).
		Ignore().
		Exec(ctx)
	if err != nil {
		tpr.logger.Errorf("Failed to create exchange with exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	savedExchange, err := client.Exchange.
		Query().
		Where(exchange.ExchangeIDEQ(exchangeId)).
		First(ctx)
	if err != nil {
		return err
	}

	existingTradingPairs, err := client.TradingPairs.
		Query().
		Where(tradingpairs.HasExchangeWith(exchangerepo.ExchangeID(exchangeId))).
		All(ctx)
	if err != nil {
		tpr.logger.Errorf("Failed to fetch existing trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	tickersToCreate := slices.Map(
		tradingPairs,
		func(item models.TradingPairs) *entities.TradingPairsCreate {
			return client.TradingPairs.
				Create().
				SetExchange(savedExchange).
				SetSymbol(item.Symbol).
				SetBase(item.Base).
				SetBasePrecision(item.BasePrecision).
				SetCounter(item.Counter).
				SetCounterPrecision(item.CounterPrecision)
		})

	if err = client.TradingPairs.
		CreateBulk(tickersToCreate...).
		OnConflictColumns(tradingpairs.FieldSymbol, tradingpairs.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to save trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	tradingPairsToDelete := slices.Filter(existingTradingPairs, func(newTradingPairs *entities.TradingPairs) bool {
		return !slices.Any(tradingPairs, func(item models.TradingPairs) bool {
			return item.Symbol == newTradingPairs.Symbol
		})
	})

	tradingPairsIdsToDelete := slices.Map(
		tradingPairsToDelete,
		func(item *entities.TradingPairs) int {
			return item.ID
		})

	if _, err = client.TradingPairs.
		Delete().
		Where(tradingpairs.IDIn(tradingPairsIdsToDelete...)).
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to delete old trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	return nil
}
