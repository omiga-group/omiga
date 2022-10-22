// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/omiga-group/omiga/src/venue/shared/entities/schema","Package":"github.com/omiga-group/omiga/src/venue/shared/entities","Schemas":[{"name":"Currency","config":{"Table":""},"edges":[{"name":"currency_base","type":"TradingPair"},{"name":"currency_counter","type":"TradingPair"}],"fields":[{"name":"symbol","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"symbol"}}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name"}}},{"name":"type","type":{"Type":6,"Ident":"currency.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"DIGITAL","V":"DIGITAL"},{"N":"FIAT","V":"FIAT"}],"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"type"}}}],"indexes":[{"fields":["symbol"]},{"fields":["name"]},{"fields":["type"]}]},{"name":"Market","config":{"Table":""},"edges":[{"name":"venue","type":"Venue","ref_name":"market","unique":true,"inverse":true,"required":true},{"name":"trading_pair","type":"TradingPair"}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name"}}},{"name":"type","type":{"Type":6,"Ident":"market.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"SPOT_TRADING","V":"SPOT_TRADING"},{"N":"MARGIN_TRADING","V":"MARGIN_TRADING"},{"N":"DERIVATIVES","V":"DERIVATIVES"},{"N":"EARN","V":"EARN"},{"N":"PERPETUAL","V":"PERPETUAL"},{"N":"FUTURES","V":"FUTURES"},{"N":"WARRANT","V":"WARRANT"},{"N":"OTC","V":"OTC"},{"N":"YIELD","V":"YIELD"},{"N":"P2P","V":"P2P"},{"N":"STRATEGY_TRADING","V":"STRATEGY_TRADING"},{"N":"SWAP_FARMING","V":"SWAP_FARMING"},{"N":"FAN_TOKEN","V":"FAN_TOKEN"},{"N":"ETF","V":"ETF"},{"N":"NFT","V":"NFT"},{"N":"SWAP","V":"SWAP"},{"N":"CFD","V":"CFD"},{"N":"LIQUIDITY","V":"LIQUIDITY"},{"N":"FARM","V":"FARM"}],"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"type"}}}],"indexes":[{"fields":["name"]},{"fields":["type"]}]},{"name":"Outbox","config":{"Table":""},"fields":[{"name":"timestamp","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"topic","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"key","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"payload","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"headers","type":{"Type":3,"Ident":"map[string]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"map[string]string","Kind":21,"PkgPath":"","Methods":{}}},"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"retry_count","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"status","type":{"Type":6,"Ident":"outbox.Status","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"Pending","V":"PENDING"},{"N":"Succeeded","V":"SUCCEEDED"},{"N":"Failed","V":"FAILED"}],"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"last_retry","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"processing_errors","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["last_retry"]},{"fields":["status"]}]},{"name":"Ticker","config":{"Table":""},"edges":[{"name":"venue","type":"Venue","ref_name":"ticker","unique":true,"inverse":true,"required":true}],"fields":[{"name":"base","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"base"}}},{"name":"base_coin_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"baseCoinId"}}},{"name":"counter","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counter"}}},{"name":"counter_coin_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counterCoinId"}}},{"name":"market","type":{"Type":3,"Ident":"models.Market","PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","PkgName":"models","Nillable":false,"RType":{"Name":"Market","Ident":"models.Market","Kind":25,"PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","Methods":{}}},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"last","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"last"}}},{"name":"volume","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"volume"}}},{"name":"converted_last","type":{"Type":3,"Ident":"models.ConvertedDetails","PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","PkgName":"models","Nillable":false,"RType":{"Name":"ConvertedDetails","Ident":"models.ConvertedDetails","Kind":25,"PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","Methods":{}}},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"converted_volume","type":{"Type":3,"Ident":"models.ConvertedDetails","PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","PkgName":"models","Nillable":false,"RType":{"Name":"ConvertedDetails","Ident":"models.ConvertedDetails","Kind":25,"PkgPath":"github.com/omiga-group/omiga/src/venue/shared/models","Methods":{}}},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"trust_score","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"trustScore"}}},{"name":"bid_ask_spread_percentage","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"bidAskSpreadPercentage"}}},{"name":"timestamp","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"timestamp"}}},{"name":"last_traded_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"lastTradedAt"}}},{"name":"last_fetch_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"lastFetchAt"}}},{"name":"is_anomaly","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":14,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"isAnomaly"}}},{"name":"is_stale","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":15,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"isStale"}}},{"name":"trade_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":16,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"tradeUrl"}}},{"name":"token_info_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":17,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"tokenInfoUrl"}}}],"indexes":[{"fields":["base"]},{"fields":["base_coin_id"]},{"fields":["counter"]},{"fields":["counter_coin_id"]},{"fields":["last"]},{"fields":["volume"]},{"fields":["trust_score"]},{"fields":["bid_ask_spread_percentage"]},{"fields":["timestamp"]},{"fields":["last_traded_at"]},{"fields":["last_fetch_at"]},{"fields":["is_anomaly"]},{"fields":["is_stale"]},{"fields":["trade_url"]},{"fields":["token_info_url"]}]},{"name":"TradingPair","config":{"Table":""},"edges":[{"name":"venue","type":"Venue","ref_name":"trading_pair","unique":true,"inverse":true,"required":true},{"name":"base","type":"Currency","ref_name":"currency_base","unique":true,"inverse":true,"required":true},{"name":"counter","type":"Currency","ref_name":"currency_counter","unique":true,"inverse":true,"required":true},{"name":"market","type":"Market","ref_name":"trading_pair","inverse":true}],"fields":[{"name":"symbol","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"symbol"}}},{"name":"base_price_min_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"basePriceMinPrecision"}}},{"name":"base_price_max_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"basePriceMaxPrecision"}}},{"name":"base_quantity_min_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"baseQuantityMinPrecision"}}},{"name":"base_quantity_max_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"baseQuantityMaxPrecision"}}},{"name":"counter_price_min_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counterPriceMinPrecision"}}},{"name":"counter_price_max_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counterPriceMaxPrecision"}}},{"name":"counter_quantity_min_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counterQuantityMinPrecision"}}},{"name":"counter_quantity_max_precision","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"counterQuantityMaxPrecision"}}}],"indexes":[{"fields":["symbol"]},{"fields":["base_price_min_precision"]},{"fields":["base_price_max_precision"]},{"fields":["base_quantity_min_precision"]},{"fields":["base_quantity_max_precision"]},{"fields":["counter_price_min_precision"]},{"fields":["counter_price_max_precision"]},{"fields":["counter_quantity_min_precision"]},{"fields":["counter_quantity_max_precision"]}]},{"name":"Venue","config":{"Table":""},"edges":[{"name":"ticker","type":"Ticker","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"trading_pair","type":"TradingPair","annotations":{"EntSQL":{"on_delete":"CASCADE"}}},{"name":"market","type":"Market","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"venue_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"venueId"}}},{"name":"type","type":{"Type":6,"Ident":"venue.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"EXCHANGE","V":"EXCHANGE"}],"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"type"}}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"name"}}},{"name":"year_established","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"yearEstablished"}}},{"name":"country","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"country"}}},{"name":"image","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"image"}}},{"name":"links","type":{"Type":3,"Ident":"map[string]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"map[string]string","Kind":21,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"has_trading_incentive","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"hasTradingIncentive"}}},{"name":"centralized","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"centralized"}}},{"name":"public_notice","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"publicNotice"}}},{"name":"alert_notice","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"alertNotice"}}},{"name":"trust_score","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"trustScore"}}},{"name":"trust_score_rank","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"trustScoreRank"}}},{"name":"trade_volume_24h_btc","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"tradeVolume24hBtc"}}},{"name":"trade_volume_24h_btc_normalized","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":14,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"tradeVolume24hBtcNormalized"}}},{"name":"maker_fee","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":15,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"makerFee"}}},{"name":"taker_fee","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":16,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"takerFee"}}},{"name":"spread_fee","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":17,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"spreadFee"}}},{"name":"support_api","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":18,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"supportAPI"}}}],"indexes":[{"fields":["venue_id"]},{"fields":["type"]},{"fields":["name"]},{"fields":["year_established"]},{"fields":["country"]},{"fields":["image"]},{"fields":["has_trading_incentive"]},{"fields":["centralized"]},{"fields":["public_notice"]},{"fields":["alert_notice"]},{"fields":["trust_score"]},{"fields":["trust_score_rank"]},{"fields":["trade_volume_24h_btc"]},{"fields":["trade_volume_24h_btc_normalized"]},{"fields":["maker_fee"]},{"fields":["taker_fee"]},{"fields":["spread_fee"]},{"fields":["support_api"]}]}],"Features":["privacy","entql","namedges","schema/snapshot","sql/schemaconfig","sql/lock","sql/modifier","sql/execquery","sql/upsert","sql/versioned-migration","namedges"]}`
