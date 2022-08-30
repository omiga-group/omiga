package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingeko/models"
	exchangemodels "github.com/omiga-group/omiga/src/exchange/shared/models"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
)

func FromCoingekoExchangeToExchange(exchange coingekov3.Exchange) models.Exchange {
	links := make(map[string]string)
	links["website"] = exchange.Url

	return models.Exchange{
		ExchangeId:          exchange.Id,
		Name:                exchange.Name,
		YearEstablished:     exchange.YearEstablished,
		Country:             exchange.Country,
		Image:               exchange.Image,
		Links:               links,
		HasTradingIncentive: exchange.HasTradingIncentive,
		TrustScore:          exchange.TrustScore,
		TrustScoreRank:      exchange.TrustScoreRank,
	}
}

func FromCoingekoExchangeDetailsToExchange(exchangeId string, exchangeDetails coingekov3.ExchangeDetails) models.Exchange {
	links := make(map[string]string)
	links["website"] = exchangeDetails.Url
	links["facebook"] = exchangeDetails.FacebookUrl
	links["reddit"] = exchangeDetails.RedditUrl
	links["twitter"] = exchangeDetails.TwitterHandle
	links["slack"] = exchangeDetails.SlackUrl
	links["telegram"] = exchangeDetails.TelegramUrl
	links["other1"] = exchangeDetails.OtherUrl1
	links["other2"] = exchangeDetails.OtherUrl2

	mappedExchange := models.Exchange{
		ExchangeId:                  exchangeId,
		Name:                        exchangeDetails.Name,
		YearEstablished:             exchangeDetails.YearEstablished,
		Country:                     exchangeDetails.Country,
		Image:                       exchangeDetails.Image,
		Links:                       links,
		HasTradingIncentive:         exchangeDetails.HasTradingIncentive,
		Centralized:                 exchangeDetails.Centralized,
		PublicNotice:                exchangeDetails.PublicNotice,
		AlertNotice:                 exchangeDetails.AlertNotice,
		TrustScore:                  exchangeDetails.TrustScore,
		TrustScoreRank:              exchangeDetails.TrustScoreRank,
		TradeVolume24hBtc:           exchangeDetails.TradeVolume24hBtc,
		TradeVolume24hBtcNormalized: exchangeDetails.TradeVolume24hBtcNormalized,
	}

	if exchangeDetails.Tickers != nil {
		mappedExchange.Tickers = slices.Map(*exchangeDetails.Tickers, func(ticker coingekov3.Ticker) models.Ticker {
			return fromCoingekoTickerToTicker(ticker)
		})
	}

	return mappedExchange
}

func fromCoingekoTickerToTicker(ticker coingekov3.Ticker) models.Ticker {
	return models.Ticker{
		Base:                   ticker.Base,
		BidAskSpreadPercentage: ticker.BidAskSpreadPercentage,
		CoinId:                 ticker.CoinId,
		ConvertedLast: exchangemodels.ConvertedDetails{
			Btc: *ticker.ConvertedLast.Btc,
			Eth: *ticker.ConvertedLast.Eth,
			Usd: *ticker.ConvertedLast.Usd,
		},
		ConvertedVolume: exchangemodels.ConvertedDetails{
			Btc: *ticker.ConvertedVolume.Btc,
			Eth: *ticker.ConvertedVolume.Eth,
			Usd: *ticker.ConvertedVolume.Usd,
		},
		IsAnomaly:    ticker.IsAnomaly,
		IsStale:      ticker.IsStale,
		Last:         ticker.Last,
		LastFetchAt:  ticker.LastFetchAt,
		LastTradedAt: ticker.LastTradedAt,
		Market: exchangemodels.Market{
			HasTradingIncentive: *ticker.Market.HasTradingIncentive,
			Identifier:          *ticker.Market.Identifier,
			Name:                *ticker.Market.Name,
		},
		Target:       ticker.Target,
		TargetCoinId: ticker.TargetCoinId,
		Timestamp:    ticker.Timestamp,
		TokenInfoUrl: ticker.TokenInfoUrl,
		TradeUrl:     ticker.TradeUrl,
		TrustScore:   ticker.TrustScore,
		Volume:       ticker.Volume,
	}
}
