package repositories

import (
	"context"
	"strings"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	currencyrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/currency"
	tradingpairrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpair"
	venuerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/venue"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type TradingPairRepository interface {
	CreateTradingPairs(
		ctx context.Context,
		venueId string,
		tradingPairs []models.TradingPair) error
}

type tradingPairRepository struct {
	logger             *zap.SugaredLogger
	entgoClient        entities.EntgoClient
	currencyRepository CurrencyRepository
	venueRepository    VenueRepository
}

func NewTradingPairRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient,
	currencyRepository CurrencyRepository,
	venueRepository VenueRepository) (TradingPairRepository, error) {
	return &tradingPairRepository{
		logger:             logger,
		entgoClient:        entgoClient,
		currencyRepository: currencyRepository,
		venueRepository:    venueRepository,
	}, nil
}

func (tpr *tradingPairRepository) CreateTradingPairs(
	ctx context.Context,
	venueId string,
	tradingPairs []models.TradingPair) error {
	savedVenueId, err := tpr.venueRepository.CreateVenue(
		ctx,
		models.Venue{
			VenueId: venueId,
			Type:    venuerepo.TypeEXCHANGE,
		})
	if err != nil {
		tpr.logger.Errorf("Failed to create venue for venue Id: %s. Error: %v", venueId, err)

		return err
	}

	allCurrencies := maps.Keys(slices.Reduce(
		tradingPairs,
		make(map[models.Currency]struct{}),
		func(tradingPair models.TradingPair, reduction map[models.Currency]struct{}) map[models.Currency]struct{} {
			reduction[models.Currency{
				Symbol: strings.ToUpper(tradingPair.Base),
				Type:   currencyrepo.TypeDIGITAL,
			}] = struct{}{}

			reduction[models.Currency{
				Symbol: strings.ToUpper(tradingPair.Counter),
				Type:   currencyrepo.TypeDIGITAL,
			}] = struct{}{}

			return reduction
		}))

	savedCurrenciesIds, err := tpr.currencyRepository.CreateCurrencies(ctx, allCurrencies)
	if err != nil {
		tpr.logger.Errorf("Failed to create currencies. Error: %v", err)

		return err
	}

	client := tpr.entgoClient.GetClient()

	tradingpairsToCreate := slices.Map(
		tradingPairs,
		func(tradingPair models.TradingPair) *entities.TradingPairCreate {
			return client.TradingPair.
				Create().
				SetVenueID(savedVenueId).
				SetSymbol(tradingPair.Symbol).
				SetBaseID(savedCurrenciesIds[strings.ToUpper(tradingPair.Base)]).
				SetNillableBasePriceMinPrecision(tradingPair.BasePriceMinPrecision).
				SetNillableBasePriceMaxPrecision(tradingPair.BasePriceMaxPrecision).
				SetNillableBaseQuantityMinPrecision(tradingPair.BaseQuantityMinPrecision).
				SetNillableBaseQuantityMaxPrecision(tradingPair.BaseQuantityMaxPrecision).
				SetCounterID(savedCurrenciesIds[strings.ToUpper(tradingPair.Counter)]).
				SetNillableCounterPriceMinPrecision(tradingPair.CounterPriceMinPrecision).
				SetNillableCounterPriceMaxPrecision(tradingPair.CounterPriceMaxPrecision).
				SetNillableCounterQuantityMinPrecision(tradingPair.CounterQuantityMinPrecision).
				SetNillableCounterQuantityMaxPrecision(tradingPair.CounterQuantityMaxPrecision)
		})

	if err = client.TradingPair.
		CreateBulk(tradingpairsToCreate...).
		OnConflictColumns(tradingpairrepo.FieldSymbol, tradingpairrepo.VenueColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to save trading pairs for venue Id: %s. Error: %v", venueId, err)

		return err
	}

	existingTradingPairs, err := client.TradingPair.
		Query().
		Where(tradingpairrepo.HasVenueWith(venuerepo.VenueID(venueId))).
		All(ctx)
	if err != nil {
		tpr.logger.Errorf("Failed to fetch existing trading pairs for venue Id: %s. Error: %v", venueId, err)

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
		Where(tradingpairrepo.IDIn(tradingPairIdsToDelete...)).
		Exec(ctx); err != nil {
		tpr.logger.Errorf("Failed to delete old trading pairs for venue Id: %s. Error: %v", venueId, err)

		return err
	}

	return nil
}
