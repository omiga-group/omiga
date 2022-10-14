package repositories

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/ticker"
	venuerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/venue"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type VenueRepository interface {
	CreateVenues(
		ctx context.Context,
		venues []models.Venue) (map[string]int, error)

	CreateVenue(
		ctx context.Context,
		venue models.Venue) (int, error)
}

type venueRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewVenueRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (VenueRepository, error) {
	return &venueRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (er *venueRepository) CreateVenues(
	ctx context.Context,
	venues []models.Venue) (map[string]int, error) {
	createdVenues := make(map[string]int)

	for _, venue := range venues {
		if savedVenueId, err := er.CreateVenue(ctx, venue); err != nil {
			return nil, err
		} else {
			createdVenues[venue.VenueId] = savedVenueId
		}
	}

	return createdVenues, nil
}

func (er *venueRepository) CreateVenue(
	ctx context.Context,
	venue models.Venue) (int, error) {
	venueId := venue.VenueId

	client := er.entgoClient.GetClient()
	err := client.Venue.
		Create().
		SetVenueID(venueId).
		SetType(venue.Type).
		SetName(venue.Name).
		SetYearEstablished(venue.YearEstablished).
		SetCountry(venue.Country).
		SetImage(venue.Image).
		SetLinks(venue.Links).
		SetHasTradingIncentive(venue.HasTradingIncentive).
		SetCentralized(venue.Centralized).
		SetPublicNotice(venue.PublicNotice).
		SetAlertNotice(venue.AlertNotice).
		SetTrustScore(venue.TrustScore).
		SetTrustScoreRank(venue.TrustScoreRank).
		SetTradeVolume24hBtc(venue.TradeVolume24hBtc).
		SetTradeVolume24hBtcNormalized(venue.TradeVolume24hBtcNormalized).
		SetNillableMakerFee(venue.MakerFee).
		SetNillableTakerFee(venue.TakerFee).
		SetNillableSpreadFee(venue.SpreadFee).
		SetNillableSupportAPI(venue.SupportAPI).
		OnConflictColumns(venuerepo.FieldVenueID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		er.logger.Errorf("Failed to save venue. Error: %v", err)
		return -1, err
	}

	savedVenue, err := client.Venue.
		Query().
		Where(venuerepo.VenueID(venue.VenueId)).
		First(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch venue with venue Id: %s. Error: %v", venueId, err)

		return -1, err
	}

	if len(venue.Tickers) == 0 {
		return savedVenue.ID, nil
	}

	tickers, err := client.Ticker.
		Query().
		Where(ticker.HasVenueWith(venuerepo.IDEQ(savedVenue.ID))).
		All(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch tickers for venue Id: %s. Error: %v", venueId, err)

		return -1, err
	}

	tickersToCreate := slices.Map(
		venue.Tickers,
		func(ticker models.Ticker) *entities.TickerCreate {
			return client.Ticker.
				Create().
				SetVenueID(savedVenue.ID).
				SetBase(ticker.Base).
				SetCounter(ticker.Counter).
				SetMarket(models.Market{
					HasTradingIncentive: ticker.Market.HasTradingIncentive,
					Identifier:          ticker.Market.Identifier,
					Name:                ticker.Market.Name,
				}).
				SetLast(ticker.Last).
				SetVolume(ticker.Volume).
				SetConvertedLast(models.ConvertedDetails{
					Btc: ticker.ConvertedLast.Btc,
					Eth: ticker.ConvertedLast.Eth,
					Usd: ticker.ConvertedLast.Usd,
				}).
				SetConvertedVolume(models.ConvertedDetails{
					Btc: ticker.ConvertedVolume.Btc,
					Eth: ticker.ConvertedVolume.Eth,
					Usd: ticker.ConvertedVolume.Usd,
				}).
				SetTrustScore(ticker.TrustScore).
				SetBidAskSpreadPercentage(ticker.BidAskSpreadPercentage).
				SetTimestamp(ticker.Timestamp).
				SetLastTradedAt(ticker.LastTradedAt).
				SetLastFetchAt(ticker.LastFetchAt).
				SetIsAnomaly(ticker.IsAnomaly).
				SetIsStale(ticker.IsStale).
				SetTradeURL(ticker.TradeUrl).
				SetNillableTokenInfoURL(ticker.TokenInfoUrl).
				SetBaseCoinID(ticker.BaseCoinId).
				SetCounterCoinID(ticker.CounterCoinId)
		})

	if err = client.Ticker.
		CreateBulk(tickersToCreate...).
		OnConflictColumns(ticker.FieldBase, ticker.FieldCounter, ticker.VenueColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to save tickers for venue Id: %s. Error: %v", venueId, err)

		return -1, err
	}

	tickersToDelete := slices.Filter(
		tickers,
		func(existingTicker *entities.Ticker) bool {
			if venue.Tickers == nil {
				return true
			}

			return !slices.Any(venue.Tickers, func(ticker models.Ticker) bool {
				return ticker.Base == existingTicker.Base && ticker.Counter == existingTicker.Counter
			})
		})

	tickerIdsToDelete := slices.Map(
		tickersToDelete,
		func(ticker *entities.Ticker) int {
			return ticker.ID
		})

	if _, err = client.Ticker.
		Delete().
		Where(ticker.IDIn(tickerIdsToDelete...)).
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to delete old tickers for venue Id: %s. Error: %v", venueId, err)

		return -1, err
	}

	return savedVenue.ID, nil
}
