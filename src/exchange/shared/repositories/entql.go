// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/coin"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/outbox"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/predicate"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/ticker"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/tradingpairs"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coin.Table,
			Columns: coin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coin.FieldID,
			},
		},
		Type: "Coin",
		Fields: map[string]*sqlgraph.FieldSpec{
			coin.FieldSymbol: {Type: field.TypeString, Column: coin.FieldSymbol},
			coin.FieldName:   {Type: field.TypeString, Column: coin.FieldName},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   exchange.Table,
			Columns: exchange.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: exchange.FieldID,
			},
		},
		Type: "Exchange",
		Fields: map[string]*sqlgraph.FieldSpec{
			exchange.FieldExchangeID:                  {Type: field.TypeString, Column: exchange.FieldExchangeID},
			exchange.FieldName:                        {Type: field.TypeString, Column: exchange.FieldName},
			exchange.FieldYearEstablished:             {Type: field.TypeInt, Column: exchange.FieldYearEstablished},
			exchange.FieldCountry:                     {Type: field.TypeString, Column: exchange.FieldCountry},
			exchange.FieldImage:                       {Type: field.TypeString, Column: exchange.FieldImage},
			exchange.FieldLinks:                       {Type: field.TypeJSON, Column: exchange.FieldLinks},
			exchange.FieldHasTradingIncentive:         {Type: field.TypeBool, Column: exchange.FieldHasTradingIncentive},
			exchange.FieldCentralized:                 {Type: field.TypeBool, Column: exchange.FieldCentralized},
			exchange.FieldPublicNotice:                {Type: field.TypeString, Column: exchange.FieldPublicNotice},
			exchange.FieldAlertNotice:                 {Type: field.TypeString, Column: exchange.FieldAlertNotice},
			exchange.FieldTrustScore:                  {Type: field.TypeInt, Column: exchange.FieldTrustScore},
			exchange.FieldTrustScoreRank:              {Type: field.TypeInt, Column: exchange.FieldTrustScoreRank},
			exchange.FieldTradeVolume24hBtc:           {Type: field.TypeFloat64, Column: exchange.FieldTradeVolume24hBtc},
			exchange.FieldTradeVolume24hBtcNormalized: {Type: field.TypeFloat64, Column: exchange.FieldTradeVolume24hBtcNormalized},
			exchange.FieldMakerFee:                    {Type: field.TypeFloat64, Column: exchange.FieldMakerFee},
			exchange.FieldTakerFee:                    {Type: field.TypeFloat64, Column: exchange.FieldTakerFee},
			exchange.FieldSpreadFee:                   {Type: field.TypeBool, Column: exchange.FieldSpreadFee},
			exchange.FieldSupportAPI:                  {Type: field.TypeBool, Column: exchange.FieldSupportAPI},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   outbox.Table,
			Columns: outbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		},
		Type: "Outbox",
		Fields: map[string]*sqlgraph.FieldSpec{
			outbox.FieldTimestamp:        {Type: field.TypeTime, Column: outbox.FieldTimestamp},
			outbox.FieldTopic:            {Type: field.TypeString, Column: outbox.FieldTopic},
			outbox.FieldKey:              {Type: field.TypeString, Column: outbox.FieldKey},
			outbox.FieldPayload:          {Type: field.TypeBytes, Column: outbox.FieldPayload},
			outbox.FieldHeaders:          {Type: field.TypeJSON, Column: outbox.FieldHeaders},
			outbox.FieldRetryCount:       {Type: field.TypeInt, Column: outbox.FieldRetryCount},
			outbox.FieldStatus:           {Type: field.TypeEnum, Column: outbox.FieldStatus},
			outbox.FieldLastRetry:        {Type: field.TypeTime, Column: outbox.FieldLastRetry},
			outbox.FieldProcessingErrors: {Type: field.TypeJSON, Column: outbox.FieldProcessingErrors},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   ticker.Table,
			Columns: ticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ticker.FieldID,
			},
		},
		Type: "Ticker",
		Fields: map[string]*sqlgraph.FieldSpec{
			ticker.FieldBase:                   {Type: field.TypeString, Column: ticker.FieldBase},
			ticker.FieldBaseCoinID:             {Type: field.TypeString, Column: ticker.FieldBaseCoinID},
			ticker.FieldCounter:                {Type: field.TypeString, Column: ticker.FieldCounter},
			ticker.FieldCounterCoinID:          {Type: field.TypeString, Column: ticker.FieldCounterCoinID},
			ticker.FieldMarket:                 {Type: field.TypeJSON, Column: ticker.FieldMarket},
			ticker.FieldLast:                   {Type: field.TypeFloat64, Column: ticker.FieldLast},
			ticker.FieldVolume:                 {Type: field.TypeFloat64, Column: ticker.FieldVolume},
			ticker.FieldConvertedLast:          {Type: field.TypeJSON, Column: ticker.FieldConvertedLast},
			ticker.FieldConvertedVolume:        {Type: field.TypeJSON, Column: ticker.FieldConvertedVolume},
			ticker.FieldTrustScore:             {Type: field.TypeString, Column: ticker.FieldTrustScore},
			ticker.FieldBidAskSpreadPercentage: {Type: field.TypeFloat64, Column: ticker.FieldBidAskSpreadPercentage},
			ticker.FieldTimestamp:              {Type: field.TypeTime, Column: ticker.FieldTimestamp},
			ticker.FieldLastTradedAt:           {Type: field.TypeTime, Column: ticker.FieldLastTradedAt},
			ticker.FieldLastFetchAt:            {Type: field.TypeTime, Column: ticker.FieldLastFetchAt},
			ticker.FieldIsAnomaly:              {Type: field.TypeBool, Column: ticker.FieldIsAnomaly},
			ticker.FieldIsStale:                {Type: field.TypeBool, Column: ticker.FieldIsStale},
			ticker.FieldTradeURL:               {Type: field.TypeString, Column: ticker.FieldTradeURL},
			ticker.FieldTokenInfoURL:           {Type: field.TypeString, Column: ticker.FieldTokenInfoURL},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tradingpairs.Table,
			Columns: tradingpairs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradingpairs.FieldID,
			},
		},
		Type: "TradingPairs",
		Fields: map[string]*sqlgraph.FieldSpec{
			tradingpairs.FieldSymbol:           {Type: field.TypeString, Column: tradingpairs.FieldSymbol},
			tradingpairs.FieldBase:             {Type: field.TypeString, Column: tradingpairs.FieldBase},
			tradingpairs.FieldBasePrecision:    {Type: field.TypeInt, Column: tradingpairs.FieldBasePrecision},
			tradingpairs.FieldCounter:          {Type: field.TypeString, Column: tradingpairs.FieldCounter},
			tradingpairs.FieldCounterPrecision: {Type: field.TypeInt, Column: tradingpairs.FieldCounterPrecision},
		},
	}
	graph.MustAddE(
		"ticker",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exchange.TickerTable,
			Columns: []string{exchange.TickerColumn},
			Bidi:    false,
		},
		"Exchange",
		"Ticker",
	)
	graph.MustAddE(
		"trading_pairs",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exchange.TradingPairsTable,
			Columns: []string{exchange.TradingPairsColumn},
			Bidi:    false,
		},
		"Exchange",
		"TradingPairs",
	)
	graph.MustAddE(
		"exchange",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticker.ExchangeTable,
			Columns: []string{ticker.ExchangeColumn},
			Bidi:    false,
		},
		"Ticker",
		"Exchange",
	)
	graph.MustAddE(
		"exchange",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpairs.ExchangeTable,
			Columns: []string{tradingpairs.ExchangeColumn},
			Bidi:    false,
		},
		"TradingPairs",
		"Exchange",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CoinQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinQuery builder.
func (cq *CoinQuery) Filter() *CoinFilter {
	return &CoinFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinMutation builder.
func (m *CoinMutation) Filter() *CoinFilter {
	return &CoinFilter{config: m.config, predicateAdder: m}
}

// CoinFilter provides a generic filtering capability at runtime for CoinQuery.
type CoinFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CoinFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(coin.FieldID))
}

// WhereSymbol applies the entql string predicate on the symbol field.
func (f *CoinFilter) WhereSymbol(p entql.StringP) {
	f.Where(p.Field(coin.FieldSymbol))
}

// WhereName applies the entql string predicate on the name field.
func (f *CoinFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(coin.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (eq *ExchangeQuery) addPredicate(pred func(s *sql.Selector)) {
	eq.predicates = append(eq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ExchangeQuery builder.
func (eq *ExchangeQuery) Filter() *ExchangeFilter {
	return &ExchangeFilter{config: eq.config, predicateAdder: eq}
}

// addPredicate implements the predicateAdder interface.
func (m *ExchangeMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ExchangeMutation builder.
func (m *ExchangeMutation) Filter() *ExchangeFilter {
	return &ExchangeFilter{config: m.config, predicateAdder: m}
}

// ExchangeFilter provides a generic filtering capability at runtime for ExchangeQuery.
type ExchangeFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ExchangeFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ExchangeFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(exchange.FieldID))
}

// WhereExchangeID applies the entql string predicate on the exchange_id field.
func (f *ExchangeFilter) WhereExchangeID(p entql.StringP) {
	f.Where(p.Field(exchange.FieldExchangeID))
}

// WhereName applies the entql string predicate on the name field.
func (f *ExchangeFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(exchange.FieldName))
}

// WhereYearEstablished applies the entql int predicate on the year_established field.
func (f *ExchangeFilter) WhereYearEstablished(p entql.IntP) {
	f.Where(p.Field(exchange.FieldYearEstablished))
}

// WhereCountry applies the entql string predicate on the country field.
func (f *ExchangeFilter) WhereCountry(p entql.StringP) {
	f.Where(p.Field(exchange.FieldCountry))
}

// WhereImage applies the entql string predicate on the image field.
func (f *ExchangeFilter) WhereImage(p entql.StringP) {
	f.Where(p.Field(exchange.FieldImage))
}

// WhereLinks applies the entql json.RawMessage predicate on the links field.
func (f *ExchangeFilter) WhereLinks(p entql.BytesP) {
	f.Where(p.Field(exchange.FieldLinks))
}

// WhereHasTradingIncentive applies the entql bool predicate on the has_trading_incentive field.
func (f *ExchangeFilter) WhereHasTradingIncentive(p entql.BoolP) {
	f.Where(p.Field(exchange.FieldHasTradingIncentive))
}

// WhereCentralized applies the entql bool predicate on the centralized field.
func (f *ExchangeFilter) WhereCentralized(p entql.BoolP) {
	f.Where(p.Field(exchange.FieldCentralized))
}

// WherePublicNotice applies the entql string predicate on the public_notice field.
func (f *ExchangeFilter) WherePublicNotice(p entql.StringP) {
	f.Where(p.Field(exchange.FieldPublicNotice))
}

// WhereAlertNotice applies the entql string predicate on the alert_notice field.
func (f *ExchangeFilter) WhereAlertNotice(p entql.StringP) {
	f.Where(p.Field(exchange.FieldAlertNotice))
}

// WhereTrustScore applies the entql int predicate on the trust_score field.
func (f *ExchangeFilter) WhereTrustScore(p entql.IntP) {
	f.Where(p.Field(exchange.FieldTrustScore))
}

// WhereTrustScoreRank applies the entql int predicate on the trust_score_rank field.
func (f *ExchangeFilter) WhereTrustScoreRank(p entql.IntP) {
	f.Where(p.Field(exchange.FieldTrustScoreRank))
}

// WhereTradeVolume24hBtc applies the entql float64 predicate on the trade_volume_24h_btc field.
func (f *ExchangeFilter) WhereTradeVolume24hBtc(p entql.Float64P) {
	f.Where(p.Field(exchange.FieldTradeVolume24hBtc))
}

// WhereTradeVolume24hBtcNormalized applies the entql float64 predicate on the trade_volume_24h_btc_normalized field.
func (f *ExchangeFilter) WhereTradeVolume24hBtcNormalized(p entql.Float64P) {
	f.Where(p.Field(exchange.FieldTradeVolume24hBtcNormalized))
}

// WhereMakerFee applies the entql float64 predicate on the maker_fee field.
func (f *ExchangeFilter) WhereMakerFee(p entql.Float64P) {
	f.Where(p.Field(exchange.FieldMakerFee))
}

// WhereTakerFee applies the entql float64 predicate on the taker_fee field.
func (f *ExchangeFilter) WhereTakerFee(p entql.Float64P) {
	f.Where(p.Field(exchange.FieldTakerFee))
}

// WhereSpreadFee applies the entql bool predicate on the spread_fee field.
func (f *ExchangeFilter) WhereSpreadFee(p entql.BoolP) {
	f.Where(p.Field(exchange.FieldSpreadFee))
}

// WhereSupportAPI applies the entql bool predicate on the support_api field.
func (f *ExchangeFilter) WhereSupportAPI(p entql.BoolP) {
	f.Where(p.Field(exchange.FieldSupportAPI))
}

// WhereHasTicker applies a predicate to check if query has an edge ticker.
func (f *ExchangeFilter) WhereHasTicker() {
	f.Where(entql.HasEdge("ticker"))
}

// WhereHasTickerWith applies a predicate to check if query has an edge ticker with a given conditions (other predicates).
func (f *ExchangeFilter) WhereHasTickerWith(preds ...predicate.Ticker) {
	f.Where(entql.HasEdgeWith("ticker", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTradingPairs applies a predicate to check if query has an edge trading_pairs.
func (f *ExchangeFilter) WhereHasTradingPairs() {
	f.Where(entql.HasEdge("trading_pairs"))
}

// WhereHasTradingPairsWith applies a predicate to check if query has an edge trading_pairs with a given conditions (other predicates).
func (f *ExchangeFilter) WhereHasTradingPairsWith(preds ...predicate.TradingPairs) {
	f.Where(entql.HasEdgeWith("trading_pairs", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (oq *OutboxQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OutboxQuery builder.
func (oq *OutboxQuery) Filter() *OutboxFilter {
	return &OutboxFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OutboxMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OutboxMutation builder.
func (m *OutboxMutation) Filter() *OutboxFilter {
	return &OutboxFilter{config: m.config, predicateAdder: m}
}

// OutboxFilter provides a generic filtering capability at runtime for OutboxQuery.
type OutboxFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OutboxFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *OutboxFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(outbox.FieldID))
}

// WhereTimestamp applies the entql time.Time predicate on the timestamp field.
func (f *OutboxFilter) WhereTimestamp(p entql.TimeP) {
	f.Where(p.Field(outbox.FieldTimestamp))
}

// WhereTopic applies the entql string predicate on the topic field.
func (f *OutboxFilter) WhereTopic(p entql.StringP) {
	f.Where(p.Field(outbox.FieldTopic))
}

// WhereKey applies the entql string predicate on the key field.
func (f *OutboxFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(outbox.FieldKey))
}

// WherePayload applies the entql []byte predicate on the payload field.
func (f *OutboxFilter) WherePayload(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldPayload))
}

// WhereHeaders applies the entql json.RawMessage predicate on the headers field.
func (f *OutboxFilter) WhereHeaders(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldHeaders))
}

// WhereRetryCount applies the entql int predicate on the retry_count field.
func (f *OutboxFilter) WhereRetryCount(p entql.IntP) {
	f.Where(p.Field(outbox.FieldRetryCount))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *OutboxFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(outbox.FieldStatus))
}

// WhereLastRetry applies the entql time.Time predicate on the last_retry field.
func (f *OutboxFilter) WhereLastRetry(p entql.TimeP) {
	f.Where(p.Field(outbox.FieldLastRetry))
}

// WhereProcessingErrors applies the entql json.RawMessage predicate on the processing_errors field.
func (f *OutboxFilter) WhereProcessingErrors(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldProcessingErrors))
}

// addPredicate implements the predicateAdder interface.
func (tq *TickerQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TickerQuery builder.
func (tq *TickerQuery) Filter() *TickerFilter {
	return &TickerFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TickerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TickerMutation builder.
func (m *TickerMutation) Filter() *TickerFilter {
	return &TickerFilter{config: m.config, predicateAdder: m}
}

// TickerFilter provides a generic filtering capability at runtime for TickerQuery.
type TickerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TickerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TickerFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(ticker.FieldID))
}

// WhereBase applies the entql string predicate on the base field.
func (f *TickerFilter) WhereBase(p entql.StringP) {
	f.Where(p.Field(ticker.FieldBase))
}

// WhereBaseCoinID applies the entql string predicate on the base_coin_id field.
func (f *TickerFilter) WhereBaseCoinID(p entql.StringP) {
	f.Where(p.Field(ticker.FieldBaseCoinID))
}

// WhereCounter applies the entql string predicate on the counter field.
func (f *TickerFilter) WhereCounter(p entql.StringP) {
	f.Where(p.Field(ticker.FieldCounter))
}

// WhereCounterCoinID applies the entql string predicate on the counter_coin_id field.
func (f *TickerFilter) WhereCounterCoinID(p entql.StringP) {
	f.Where(p.Field(ticker.FieldCounterCoinID))
}

// WhereMarket applies the entql json.RawMessage predicate on the market field.
func (f *TickerFilter) WhereMarket(p entql.BytesP) {
	f.Where(p.Field(ticker.FieldMarket))
}

// WhereLast applies the entql float64 predicate on the last field.
func (f *TickerFilter) WhereLast(p entql.Float64P) {
	f.Where(p.Field(ticker.FieldLast))
}

// WhereVolume applies the entql float64 predicate on the volume field.
func (f *TickerFilter) WhereVolume(p entql.Float64P) {
	f.Where(p.Field(ticker.FieldVolume))
}

// WhereConvertedLast applies the entql json.RawMessage predicate on the converted_last field.
func (f *TickerFilter) WhereConvertedLast(p entql.BytesP) {
	f.Where(p.Field(ticker.FieldConvertedLast))
}

// WhereConvertedVolume applies the entql json.RawMessage predicate on the converted_volume field.
func (f *TickerFilter) WhereConvertedVolume(p entql.BytesP) {
	f.Where(p.Field(ticker.FieldConvertedVolume))
}

// WhereTrustScore applies the entql string predicate on the trust_score field.
func (f *TickerFilter) WhereTrustScore(p entql.StringP) {
	f.Where(p.Field(ticker.FieldTrustScore))
}

// WhereBidAskSpreadPercentage applies the entql float64 predicate on the bid_ask_spread_percentage field.
func (f *TickerFilter) WhereBidAskSpreadPercentage(p entql.Float64P) {
	f.Where(p.Field(ticker.FieldBidAskSpreadPercentage))
}

// WhereTimestamp applies the entql time.Time predicate on the timestamp field.
func (f *TickerFilter) WhereTimestamp(p entql.TimeP) {
	f.Where(p.Field(ticker.FieldTimestamp))
}

// WhereLastTradedAt applies the entql time.Time predicate on the last_traded_at field.
func (f *TickerFilter) WhereLastTradedAt(p entql.TimeP) {
	f.Where(p.Field(ticker.FieldLastTradedAt))
}

// WhereLastFetchAt applies the entql time.Time predicate on the last_fetch_at field.
func (f *TickerFilter) WhereLastFetchAt(p entql.TimeP) {
	f.Where(p.Field(ticker.FieldLastFetchAt))
}

// WhereIsAnomaly applies the entql bool predicate on the is_anomaly field.
func (f *TickerFilter) WhereIsAnomaly(p entql.BoolP) {
	f.Where(p.Field(ticker.FieldIsAnomaly))
}

// WhereIsStale applies the entql bool predicate on the is_stale field.
func (f *TickerFilter) WhereIsStale(p entql.BoolP) {
	f.Where(p.Field(ticker.FieldIsStale))
}

// WhereTradeURL applies the entql string predicate on the trade_url field.
func (f *TickerFilter) WhereTradeURL(p entql.StringP) {
	f.Where(p.Field(ticker.FieldTradeURL))
}

// WhereTokenInfoURL applies the entql string predicate on the token_info_url field.
func (f *TickerFilter) WhereTokenInfoURL(p entql.StringP) {
	f.Where(p.Field(ticker.FieldTokenInfoURL))
}

// WhereHasExchange applies a predicate to check if query has an edge exchange.
func (f *TickerFilter) WhereHasExchange() {
	f.Where(entql.HasEdge("exchange"))
}

// WhereHasExchangeWith applies a predicate to check if query has an edge exchange with a given conditions (other predicates).
func (f *TickerFilter) WhereHasExchangeWith(preds ...predicate.Exchange) {
	f.Where(entql.HasEdgeWith("exchange", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tpq *TradingPairsQuery) addPredicate(pred func(s *sql.Selector)) {
	tpq.predicates = append(tpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TradingPairsQuery builder.
func (tpq *TradingPairsQuery) Filter() *TradingPairsFilter {
	return &TradingPairsFilter{config: tpq.config, predicateAdder: tpq}
}

// addPredicate implements the predicateAdder interface.
func (m *TradingPairsMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TradingPairsMutation builder.
func (m *TradingPairsMutation) Filter() *TradingPairsFilter {
	return &TradingPairsFilter{config: m.config, predicateAdder: m}
}

// TradingPairsFilter provides a generic filtering capability at runtime for TradingPairsQuery.
type TradingPairsFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TradingPairsFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TradingPairsFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tradingpairs.FieldID))
}

// WhereSymbol applies the entql string predicate on the symbol field.
func (f *TradingPairsFilter) WhereSymbol(p entql.StringP) {
	f.Where(p.Field(tradingpairs.FieldSymbol))
}

// WhereBase applies the entql string predicate on the base field.
func (f *TradingPairsFilter) WhereBase(p entql.StringP) {
	f.Where(p.Field(tradingpairs.FieldBase))
}

// WhereBasePrecision applies the entql int predicate on the base_precision field.
func (f *TradingPairsFilter) WhereBasePrecision(p entql.IntP) {
	f.Where(p.Field(tradingpairs.FieldBasePrecision))
}

// WhereCounter applies the entql string predicate on the counter field.
func (f *TradingPairsFilter) WhereCounter(p entql.StringP) {
	f.Where(p.Field(tradingpairs.FieldCounter))
}

// WhereCounterPrecision applies the entql int predicate on the counter_precision field.
func (f *TradingPairsFilter) WhereCounterPrecision(p entql.IntP) {
	f.Where(p.Field(tradingpairs.FieldCounterPrecision))
}

// WhereHasExchange applies a predicate to check if query has an edge exchange.
func (f *TradingPairsFilter) WhereHasExchange() {
	f.Where(entql.HasEdge("exchange"))
}

// WhereHasExchangeWith applies a predicate to check if query has an edge exchange with a given conditions (other predicates).
func (f *TradingPairsFilter) WhereHasExchangeWith(preds ...predicate.Exchange) {
	f.Where(entql.HasEdgeWith("exchange", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
