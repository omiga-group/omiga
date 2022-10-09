package repositories

import (
	"context"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
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
	logger             *zap.SugaredLogger
	entgoClient        entities.EntgoClient
	coinRepository     CoinRepository
	exchangeRepository ExchangeRepository
}

func NewTradingPairRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient,
	coinRepository CoinRepository,
	exchangeRepository ExchangeRepository) (TradingPairRepository, error) {
	return &tradingPairRepository{
		logger:             logger,
		entgoClient:        entgoClient,
		coinRepository:     coinRepository,
		exchangeRepository: exchangeRepository,
	}, nil
}

func (tpr *tradingPairRepository) CreateTradingPairs(
	ctx context.Context,
	exchangeId string,
	tradingPairs []models.TradingPair) error {
	savedExchangeId, err := tpr.exchangeRepository.CreateExchange(
		ctx,
		models.Exchange{
			ExchangeId: exchangeId,
		})
	if err != nil {
		tpr.logger.Errorf("Failed to create exchange for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	allCoins := maps.Keys(slices.Reduce(
		tradingPairs,
		make(map[models.Coin]struct{}),
		func(tradingPair models.TradingPair, reduction map[models.Coin]struct{}) map[models.Coin]struct{} {
			reduction[models.Coin{
				Symbol: tradingPair.Base,
			}] = struct{}{}

			reduction[models.Coin{
				Symbol: tradingPair.Counter,
			}] = struct{}{}

			return reduction
		}))

	savedCoinsIds, err := tpr.coinRepository.CreateCoins(ctx, allCoins)
	if err != nil {
		tpr.logger.Errorf("Failed to create coins. Error: %v", err)

		return err
	}

	client := tpr.entgoClient.GetClient()

	tradingpairsToCreate := slices.Map(
		tradingPairs,
		func(tradingPair models.TradingPair) *entities.TradingPairCreate {
			return client.TradingPair.
				Create().
				SetExchangeID(savedExchangeId).
				SetSymbol(tradingPair.Symbol).
				SetBaseID(savedCoinsIds[tradingPair.Base]).
				SetBasePrecision(tradingPair.BasePrecision).
				SetCounterID(savedCoinsIds[tradingPair.Counter]).
				SetCounterPrecision(tradingPair.CounterPrecision)
		})

	if err = client.TradingPair.
		CreateBulk(tradingpairsToCreate...).
		OnConflictColumns(tradingpair.FieldSymbol, tradingpair.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to save trading pairs for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	existingTradingPairs, err := client.TradingPair.
		Query().
		Where(tradingpair.HasExchangeWith(exchangerepo.ExchangeID(exchangeId))).
		All(ctx)
	if err != nil {
		tpr.logger.Errorf("Failed to fetch existing trading pairs for exchange Id: %s. Error: %v", exchangeId, err)

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
		tpr.logger.Errorf("Failed to delete old trading pairs for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	return nil
}
