package subscribers

import (
	"context"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/ticker"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type CoingekoSubscriber interface {
}

type coingekoSubscriber struct {
	ctx              context.Context
	logger           *zap.SugaredLogger
	coingekoSettings configuration.CoingekoSettings
	entgoClient      repositories.EntgoClient
}

func NewCoingekoSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingekoSettings configuration.CoingekoSettings,
	entgoClient repositories.EntgoClient) (CoingekoSubscriber, error) {
	instance := &coingekoSubscriber{
		ctx:              ctx,
		logger:           logger,
		coingekoSettings: coingekoSettings,
		entgoClient:      entgoClient,
	}

	if _, err := cronService.GetCron().AddJob("0/1 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (cs *coingekoSubscriber) Run() {
	coingekoClient, err := coingekov3.NewClientWithResponses(cs.coingekoSettings.BaseUrl)
	if err != nil {
		cs.logger.Errorf(
			"Failed to create coingeko client. Error: %v",
			err)
	}

	exchangesListResponse, err := coingekoClient.GetExchangesListWithResponse(cs.ctx)
	if err != nil {
		cs.logger.Errorf(
			"Failed to get exchange list. Error: %v",
			err)

		return
	}

	if exchangesListResponse.HTTPResponse.StatusCode != 200 {
		cs.logger.Errorf(
			"Failed to get exchange list. Return status code is %d",
			exchangesListResponse.HTTPResponse.StatusCode)

		return
	}

	for _, exchangeIdName := range *exchangesListResponse.JSON200 {
		// This is to avoid coingeko rate limiter blocking us from querying exchanges details
		select {
		case <-cs.ctx.Done():
		case <-time.After(2 * time.Second):
		}

		if cs.ctx.Err() == context.Canceled {
			break
		}

		exchangeIdResponse, err := coingekoClient.GetExchangesIdWithResponse(
			cs.ctx,
			exchangeIdName.Id)
		if err != nil {
			cs.logger.Errorf("Failed to get exchange details. Error: %v", err)

			continue
		}

		if exchangeIdResponse.HTTPResponse.StatusCode != 200 {
			cs.logger.Errorf(
				"Failed to get exchange details. Return status code is %d",
				exchangeIdResponse.HTTPResponse.StatusCode)

			continue
		}

		exchangeDetails := *exchangeIdResponse.JSON200

		links := make(map[string]string)
		links["website"] = exchangeDetails.Url
		links["facebook"] = exchangeDetails.FacebookUrl
		links["reddit"] = exchangeDetails.RedditUrl
		links["twitter"] = exchangeDetails.TwitterHandle
		links["slack"] = exchangeDetails.SlackUrl
		links["telegram"] = exchangeDetails.TelegramUrl
		links["other1"] = exchangeDetails.OtherUrl1
		links["other2"] = exchangeDetails.OtherUrl2

		client := cs.entgoClient.GetClient()

		if err = client.Exchange.
			Create().
			SetExchangeID(exchangeIdName.Id).
			SetName(exchangeDetails.Name).
			SetYearEstablished(exchangeDetails.YearEstablished).
			SetCountry(exchangeDetails.Country).
			SetImage(exchangeDetails.Image).
			SetLinks(links).
			SetHasTradingIncentive(exchangeDetails.HasTradingIncentive).
			SetCentralized(exchangeDetails.Centralized).
			SetPublicNotice(exchangeDetails.PublicNotice).
			SetAlertNotice(exchangeDetails.AlertNotice).
			SetTrustScore(exchangeDetails.TrustScore).
			SetTrustScoreRank(exchangeDetails.TrustScoreRank).
			SetTradeVolume24hBtc(exchangeDetails.TradeVolume24hBtc).
			SetTradeVolume24hBtcNormalized(exchangeDetails.TradeVolume24hBtcNormalized).
			OnConflictColumns(exchange.FieldExchangeID).
			UpdateNewValues().
			Exec(cs.ctx); err != nil {
			cs.logger.Errorf("Failed to save exchange details. Error: %v", err)

			continue
		}

		tickers, err := client.Ticker.
			Query().
			Where(ticker.HasExchangeWith(exchange.ExchangeIDEQ(exchangeIdName.Id))).
			All(cs.ctx)
		if err != nil {
			cs.logger.Errorf("Failed to fetch tickers for exchange Id: %s. Error: %v", exchangeIdName.Id, err)

			continue
		}

		if exchangeDetails.Tickers != nil {
			tickersToCreate := slices.Map(
				*exchangeDetails.Tickers,
				func(ticker coingekov3.Ticker) *repositories.TickerCreate {
					return client.Ticker.
						Create().
						SetBase(ticker.Base).
						SetTarget(ticker.Target).
						SetMarket(models.Market{
							HasTradingIncentive: *ticker.Market.HasTradingIncentive,
							Identifier:          *ticker.Market.Identifier,
							Name:                *ticker.Market.Name,
						}).
						SetLast(ticker.Last).
						SetVolume(ticker.Volume).
						SetConvertedLast(models.ConvertedDetails{
							Btc: *ticker.ConvertedLast.Btc,
							Eth: *ticker.ConvertedLast.Eth,
							Usd: *ticker.ConvertedLast.Usd,
						}).
						SetConvertedVolume(models.ConvertedDetails{
							Btc: *ticker.ConvertedVolume.Btc,
							Eth: *ticker.ConvertedVolume.Eth,
							Usd: *ticker.ConvertedVolume.Usd,
						}).
						SetTrustScore(ticker.TrustScore).
						SetBidAskSpreadPercentage(ticker.BidAskSpreadPercentage).
						SetTimestamp(ticker.Timestamp).
						SetLastTradedAt(ticker.LastTradedAt).
						SetLastFetchAt(ticker.LastFetchAt).
						SetIsAnomaly(ticker.IsAnomaly).
						SetIsStale(ticker.IsStale).
						SetTradeURL(ticker.TradeUrl).
						SetTokenInfoURL(*ticker.TokenInfoUrl).
						SetCoinID(ticker.CoinId).
						SetTargetCoinID(ticker.TargetCoinId)
				})

			if err = client.Ticker.
				CreateBulk(tickersToCreate...).
				OnConflictColumns(ticker.FieldBase, ticker.FieldTarget).
				UpdateNewValues().
				Exec(cs.ctx); err != nil {
				cs.logger.Errorf("Failed to save tickers for exchange Id: %s. Error: %v", exchangeIdName.Id, err)

				continue
			}
		}

		tickersToDelete := slices.Filter(
			tickers,
			func(existingTicker *repositories.Ticker) bool {
				if exchangeDetails.Tickers == nil {
					return true
				}

				return !slices.Any(*exchangeDetails.Tickers, func(ticker coingekov3.Ticker) bool {
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
			Exec(cs.ctx); err != nil {
			cs.logger.Errorf("Failed to delete old tickers for exchange Id: %s. Error: %v", exchangeIdName.Id, err)

			continue
		}
	}
}
