package repositories

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	exchangerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpair"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type TradingPairRepository interface {
	CreateTradingPairs(
		ctx context.Context,
		exchangeId string,
		tradingPairs []models.TradingPair) error
}

type tradingPairRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewTradingPairRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (TradingPairRepository, error) {
	return &tradingPairRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (tpr *tradingPairRepository) CreateTradingPairs(
	ctx context.Context,
	exchangeId string,
	tradingPairs []models.TradingPair) error {
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

	existingTradingPairs, err := client.TradingPair.
		Query().
		Where(tradingpair.HasExchangeWith(exchangerepo.ExchangeID(exchangeId))).
		All(ctx)
	if err != nil {
		tpr.logger.Errorf("Failed to fetch existing trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	tickersToCreate := slices.Map(
		tradingPairs,
		func(tradingPair models.TradingPair) *entities.TradingPairCreate {
			return client.TradingPair.
				Create().
				SetExchange(savedExchange).
				SetSymbol(tradingPair.Symbol).
				SetBase(tradingPair.Base).
				SetBasePrecision(tradingPair.BasePrecision).
				SetCounter(tradingPair.Counter).
				SetCounterPrecision(tradingPair.CounterPrecision)
		})

	if err = client.TradingPair.
		CreateBulk(tickersToCreate...).
		OnConflictColumns(tradingpair.FieldSymbol, tradingpair.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to save trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	tradingPairsToDelete := slices.Filter(existingTradingPairs, func(existingTradingPair *entities.TradingPair) bool {
		return !slices.Any(tradingPairs, func(tradingPair models.TradingPair) bool {
			return tradingPair.Symbol == existingTradingPair.Symbol
		})
	})

	tradingPairIdsToDelete := slices.Map(
		tradingPairsToDelete,
		func(tradingPair *entities.TradingPair) int {
			return tradingPair.ID
		})

	if _, err = client.TradingPair.
		Delete().
		Where(tradingpair.IDIn(tradingPairIdsToDelete...)).
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to delete old trading pairs for exchange Id: %s. Error: %v", savedExchange.ExchangeID, err)

		return err
	}

	return nil
}
