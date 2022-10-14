package repositories

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	currencyrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/currency"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type CurrencyRepository interface {
	CreateCurrencies(
		ctx context.Context,
		currencies []models.Currency) (map[string]int, error)

	CreateCurrency(
		ctx context.Context,
		currency models.Currency) (int, error)
}

type currencyRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewCurrencyRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (CurrencyRepository, error) {
	return &currencyRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (cr *currencyRepository) CreateCurrencies(
	ctx context.Context,
	currencies []models.Currency) (map[string]int, error) {
	createdCurrencies := make(map[string]int)

	for _, currency := range currencies {
		if savedCurrencyId, err := cr.CreateCurrency(ctx, currency); err != nil {
			return nil, err
		} else {
			createdCurrencies[currency.Symbol] = savedCurrencyId
		}
	}

	return createdCurrencies, nil
}

func (cr *currencyRepository) CreateCurrency(
	ctx context.Context,
	currency models.Currency) (int, error) {
	client := cr.entgoClient.GetClient()
	err := client.Currency.
		Create().
		SetSymbol(currency.Symbol).
		SetNillableName(currency.Name).
		SetType(currency.Type).
		OnConflictColumns(currencyrepo.FieldSymbol).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		cr.logger.Errorf("Failed to save currency. Error: %v", err)
		return -1, err
	}

	savedCurrency, err := client.Currency.
		Query().
		Where(currencyrepo.SymbolEQ(currency.Symbol)).
		First(ctx)
	if err != nil {
		return -1, err
	}

	return savedCurrency.ID, nil
}
