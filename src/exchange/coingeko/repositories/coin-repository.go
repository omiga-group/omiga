package repositories

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coinrepo "github.com/omiga-group/omiga/src/exchange/shared/repositories/coin"
	"go.uber.org/zap"
)

type CoinRepository interface {
	CreateCoins(
		ctx context.Context,
		coins []models.Coin) error

	CreateCoin(
		ctx context.Context,
		coin models.Coin) error
}

type coinRepository struct {
	logger      *zap.SugaredLogger
	entgoClient repositories.EntgoClient
}

func NewCoinRepository(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient) (CoinRepository, error) {
	return &coinRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (cr *coinRepository) CreateCoins(
	ctx context.Context,
	coins []models.Coin) error {
	for _, coin := range coins {
		if err := cr.CreateCoin(ctx, coin); err != nil {
			return err
		}
	}

	return nil
}

func (cr *coinRepository) CreateCoin(
	ctx context.Context,
	coin models.Coin) error {
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
		return err
	}

	return nil
}
