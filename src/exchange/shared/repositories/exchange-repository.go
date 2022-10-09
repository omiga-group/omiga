package repositories

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	exchangerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/ticker"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	modelsexchange "github.com/omiga-group/omiga/src/exchange/shared/models"
	"go.uber.org/zap"
)

type ExchangeRepository interface {
	CreateExchanges(
		ctx context.Context,
		exchanges []models.Exchange) (map[string]int, error)

	CreateExchange(
		ctx context.Context,
		exchange models.Exchange) (int, error)
}

type exchangeRepository struct {
	logger      *zap.SugaredLogger
	entgoClient entities.EntgoClient
}

func NewExchangeRepository(
	logger *zap.SugaredLogger,
	entgoClient entities.EntgoClient) (ExchangeRepository, error) {
	return &exchangeRepository{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (er *exchangeRepository) CreateExchanges(
	ctx context.Context,
	exchanges []models.Exchange) (map[string]int, error) {
	createdExchanges := make(map[string]int)

	for _, exchange := range exchanges {
		if savedExchangeId, err := er.CreateExchange(ctx, exchange); err != nil {
			return nil, err
		} else {
			createdExchanges[exchange.ExchangeId] = savedExchangeId
		}
	}

	return createdExchanges, nil
}

func (er *exchangeRepository) CreateExchange(
	ctx context.Context,
	exchange models.Exchange) (int, error) {
	exchangeId := exchange.ExchangeId

	client := er.entgoClient.GetClient()
	err := client.Exchange.
		Create().
		SetExchangeID(exchangeId).
		SetName(exchange.Name).
		SetYearEstablished(exchange.YearEstablished).
		SetCountry(exchange.Country).
		SetImage(exchange.Image).
		SetLinks(exchange.Links).
		SetHasTradingIncentive(exchange.HasTradingIncentive).
		SetCentralized(exchange.Centralized).
		SetPublicNotice(exchange.PublicNotice).
		SetAlertNotice(exchange.AlertNotice).
		SetTrustScore(exchange.TrustScore).
		SetTrustScoreRank(exchange.TrustScoreRank).
		SetTradeVolume24hBtc(exchange.TradeVolume24hBtc).
		SetTradeVolume24hBtcNormalized(exchange.TradeVolume24hBtcNormalized).
		SetNillableMakerFee(exchange.MakerFee).
		SetNillableTakerFee(exchange.TakerFee).
		SetNillableSpreadFee(exchange.SpreadFee).
		SetNillableSupportAPI(exchange.SupportAPI).
		OnConflictColumns(exchangerepo.FieldExchangeID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		er.logger.Errorf("Failed to save exchange. Error: %v", err)
		return -1, err
	}

	savedExchange, err := client.Exchange.
		Query().
		Where(exchangerepo.ExchangeID(exchange.ExchangeId)).
		First(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch exchange with exchange Id: %s. Error: %v", exchangeId, err)

		return -1, err
	}

	if len(exchange.Tickers) == 0 {
		return savedExchange.ID, nil
	}

	tickers, err := client.Ticker.
		Query().
		Where(ticker.HasExchangeWith(exchangerepo.IDEQ(savedExchange.ID))).
		All(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return -1, err
	}

	tickersToCreate := slices.Map(
		exchange.Tickers,
		func(ticker models.Ticker) *entities.TickerCreate {
			return client.Ticker.
				Create().
				SetExchangeID(savedExchange.ID).
				SetBase(ticker.Base).
				SetCounter(ticker.Counter).
				SetMarket(modelsexchange.Market{
					HasTradingIncentive: ticker.Market.HasTradingIncentive,
					Identifier:          ticker.Market.Identifier,
					Name:                ticker.Market.Name,
				}).
				SetLast(ticker.Last).
				SetVolume(ticker.Volume).
				SetConvertedLast(modelsexchange.ConvertedDetails{
					Btc: ticker.ConvertedLast.Btc,
					Eth: ticker.ConvertedLast.Eth,
					Usd: ticker.ConvertedLast.Usd,
				}).
				SetConvertedVolume(modelsexchange.ConvertedDetails{
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
		OnConflictColumns(ticker.FieldBase, ticker.FieldCounter, ticker.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to save tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return -1, err
	}

	tickersToDelete := slices.Filter(
		tickers,
		func(existingTicker *entities.Ticker) bool {
			if exchange.Tickers == nil {
				return true
			}

			return !slices.Any(exchange.Tickers, func(ticker models.Ticker) bool {
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
		er.logger.Errorf("Failed to delete old tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return -1, err
	}

	return savedExchange.ID, nil
}
