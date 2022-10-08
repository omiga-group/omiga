package repositories

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	coinrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/coin"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type CoinRepository interface {
	CreateCoins(
		ctx context.Context,
		coins []models.Coin) (map[string]int, error)

	CreateCoin(
		ctx context.Context,
		coin models.Coin) (int, error)
}

type coinRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewCoinRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (CoinRepository, error) {
	return &coinRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (cr *coinRepository) CreateCoins(
	ctx context.Context,
	coins []models.Coin) (map[string]int, error) {
	createdCoins := make(map[string]int)

	for _, coin := range coins {
		if savedCoinId, err := cr.CreateCoin(ctx, coin); err != nil {
			return nil, err
		} else {
			createdCoins[coin.Symbol] = savedCoinId
		}
	}

	return createdCoins, nil
}

func (cr *coinRepository) CreateCoin(
	ctx context.Context,
	coin models.Coin) (int, error) {
	client := cr.entgoClient.GetClient()
	err := client.Coin.
		Create().
		SetSymbol(coin.Symbol).
		SetNillableName(coin.Name).
		OnConflictColumns(coinrepo.FieldSymbol).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		cr.logger.Errorf("Failed to save coin. Error: %v", err)
		return -1, err
	}

	savedCoin, err := client.Coin.
		Query().
		Where(coinrepo.SymbolEQ(coin.Symbol)).
		First(ctx)
	if err != nil {
		return -1, err
	}

	return savedCoin.ID, nil
}
