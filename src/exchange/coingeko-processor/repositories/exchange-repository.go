package repositories

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	modelsexchange "github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	exchangerepo "github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/ticker"
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
	entgoClient repositories.EntgoClient
}

func NewExchangeRepository(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient) (ExchangeRepository, error) {
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
		func(ticker models.Ticker) *repositories.TickerCreate {
			return client.Ticker.
				Create().
				SetExchangeID(savedExchange.ID).
				SetBase(ticker.Base).
				SetTarget(ticker.Target).
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
				SetCoinID(ticker.CoinId).
				SetTargetCoinID(ticker.TargetCoinId)
		})

	if err = client.Ticker.
		CreateBulk(tickersToCreate...).
		OnConflictColumns(ticker.FieldBase, ticker.FieldTarget, ticker.ExchangeColumn).
		UpdateNewValues().
		Exec(ctx); err != nil {
		er.logger.Errorf("Failed to save tickers for exchange Id: %s. Error: %v", exchangeId, err)

		return err
	}

	tickersToDelete := slices.Filter(
		tickers,
		func(existingTicker *repositories.Ticker) bool {
			if exchange.Tickers == nil {
				return true
			}

			return !slices.Any(exchange.Tickers, func(ticker models.Ticker) bool {
				return ticker.Base == existingTicker.Base && ticker.Target == existingTicker.Target
			})
		})

	tickerIdsToDelete := slices.Map(
		tickersToDelete,
		func(ticker *repositories.Ticker) int {
			return ticker.ID
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
