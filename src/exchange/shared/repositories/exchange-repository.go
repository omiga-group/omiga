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
		exchanges []models.Exchange) error

	CreateExchange(
		ctx context.Context,
		exchange models.Exchange) error
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
	exchanges []models.Exchange) error {
	for _, exchange := range exchanges {
		if err := er.CreateExchange(ctx, exchange); err != nil {
			return err
		}
	}

	return nil
}

func (er *exchangeRepository) CreateExchange(
	ctx context.Context,
	exchange models.Exchange) error {
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
		return err
	}

	if len(exchange.Tickers) == 0 {
		return nil
	}

	savedExchange, err := client.Exchange.
		Query().
		Where(exchangerepo.ExchangeID(exchange.ExchangeId)).
		First(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch exchange with exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	tickers, err := client.Ticker.
		Query().
		Where(ticker.HasExchangeWith(exchangerepo.IDEQ(savedExchange.ID))).
		All(ctx)
	if err != nil {
		er.logger.Errorf("Failed to fetch tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	tickersToCreate := slices.Map(
		exchange.Tickers,
		func(item models.Ticker) *entities.TickerCreate {
			return client.Ticker.
				Create().
				SetExchangeID(savedExchange.ID).
				SetBase(item.Base).
				SetCounter(item.Counter).
				SetMarket(modelsexchange.Market{
					HasTradingIncentive: item.Market.HasTradingIncentive,
					Identifier:          item.Market.Identifier,
					Name:                item.Market.Name,
				}).
				SetLast(item.Last).
				SetVolume(item.Volume).
				SetConvertedLast(modelsexchange.ConvertedDetails{
					Btc: item.ConvertedLast.Btc,
					Eth: item.ConvertedLast.Eth,
					Usd: item.ConvertedLast.Usd,
				}).
				SetConvertedVolume(modelsexchange.ConvertedDetails{
					Btc: item.ConvertedVolume.Btc,
					Eth: item.ConvertedVolume.Eth,
					Usd: item.ConvertedVolume.Usd,
				}).
				SetTrustScore(item.TrustScore).
				SetBidAskSpreadPercentage(item.BidAskSpreadPercentage).
				SetTimestamp(item.Timestamp).
				SetLastTradedAt(item.LastTradedAt).
				SetLastFetchAt(item.LastFetchAt).
				SetIsAnomaly(item.IsAnomaly).
				SetIsStale(item.IsStale).
				SetTradeURL(item.TradeUrl).
				SetNillableTokenInfoURL(item.TokenInfoUrl).
				SetBaseCoinID(item.BaseCoinId).
				SetCounterCoinID(item.CounterCoinId)
		})

	if err = client.Ticker.
		CreateBulk(tickersToCreate...).
		OnConflictColumns(ticker.FieldBase, ticker.FieldCounter, ticker.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to save tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return err
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
		func(item *entities.Ticker) int {
			return item.ID
		})

	if _, err = client.Ticker.
		Delete().
		Where(ticker.IDIn(tickerIdsToDelete...)).
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to delete old tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	return nil
}
