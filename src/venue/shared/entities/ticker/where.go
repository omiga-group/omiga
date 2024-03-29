// Code generated by ent, DO NOT EDIT.

package ticker

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldID, id))
}

// Base applies equality check predicate on the "base" field. It's identical to BaseEQ.
func Base(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBase, v))
}

// BaseCoinID applies equality check predicate on the "base_coin_id" field. It's identical to BaseCoinIDEQ.
func BaseCoinID(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBaseCoinID, v))
}

// Counter applies equality check predicate on the "counter" field. It's identical to CounterEQ.
func Counter(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldCounter, v))
}

// CounterCoinID applies equality check predicate on the "counter_coin_id" field. It's identical to CounterCoinIDEQ.
func CounterCoinID(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldCounterCoinID, v))
}

// Last applies equality check predicate on the "last" field. It's identical to LastEQ.
func Last(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLast, v))
}

// Volume applies equality check predicate on the "volume" field. It's identical to VolumeEQ.
func Volume(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldVolume, v))
}

// TrustScore applies equality check predicate on the "trust_score" field. It's identical to TrustScoreEQ.
func TrustScore(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTrustScore, v))
}

// BidAskSpreadPercentage applies equality check predicate on the "bid_ask_spread_percentage" field. It's identical to BidAskSpreadPercentageEQ.
func BidAskSpreadPercentage(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBidAskSpreadPercentage, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTimestamp, v))
}

// LastTradedAt applies equality check predicate on the "last_traded_at" field. It's identical to LastTradedAtEQ.
func LastTradedAt(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLastTradedAt, v))
}

// LastFetchAt applies equality check predicate on the "last_fetch_at" field. It's identical to LastFetchAtEQ.
func LastFetchAt(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLastFetchAt, v))
}

// IsAnomaly applies equality check predicate on the "is_anomaly" field. It's identical to IsAnomalyEQ.
func IsAnomaly(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldIsAnomaly, v))
}

// IsStale applies equality check predicate on the "is_stale" field. It's identical to IsStaleEQ.
func IsStale(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldIsStale, v))
}

// TradeURL applies equality check predicate on the "trade_url" field. It's identical to TradeURLEQ.
func TradeURL(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTradeURL, v))
}

// TokenInfoURL applies equality check predicate on the "token_info_url" field. It's identical to TokenInfoURLEQ.
func TokenInfoURL(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTokenInfoURL, v))
}

// BaseEQ applies the EQ predicate on the "base" field.
func BaseEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBase, v))
}

// BaseNEQ applies the NEQ predicate on the "base" field.
func BaseNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldBase, v))
}

// BaseIn applies the In predicate on the "base" field.
func BaseIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldBase, vs...))
}

// BaseNotIn applies the NotIn predicate on the "base" field.
func BaseNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldBase, vs...))
}

// BaseGT applies the GT predicate on the "base" field.
func BaseGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldBase, v))
}

// BaseGTE applies the GTE predicate on the "base" field.
func BaseGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldBase, v))
}

// BaseLT applies the LT predicate on the "base" field.
func BaseLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldBase, v))
}

// BaseLTE applies the LTE predicate on the "base" field.
func BaseLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldBase, v))
}

// BaseContains applies the Contains predicate on the "base" field.
func BaseContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldBase, v))
}

// BaseHasPrefix applies the HasPrefix predicate on the "base" field.
func BaseHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldBase, v))
}

// BaseHasSuffix applies the HasSuffix predicate on the "base" field.
func BaseHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldBase, v))
}

// BaseEqualFold applies the EqualFold predicate on the "base" field.
func BaseEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldBase, v))
}

// BaseContainsFold applies the ContainsFold predicate on the "base" field.
func BaseContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldBase, v))
}

// BaseCoinIDEQ applies the EQ predicate on the "base_coin_id" field.
func BaseCoinIDEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBaseCoinID, v))
}

// BaseCoinIDNEQ applies the NEQ predicate on the "base_coin_id" field.
func BaseCoinIDNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldBaseCoinID, v))
}

// BaseCoinIDIn applies the In predicate on the "base_coin_id" field.
func BaseCoinIDIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldBaseCoinID, vs...))
}

// BaseCoinIDNotIn applies the NotIn predicate on the "base_coin_id" field.
func BaseCoinIDNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldBaseCoinID, vs...))
}

// BaseCoinIDGT applies the GT predicate on the "base_coin_id" field.
func BaseCoinIDGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldBaseCoinID, v))
}

// BaseCoinIDGTE applies the GTE predicate on the "base_coin_id" field.
func BaseCoinIDGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldBaseCoinID, v))
}

// BaseCoinIDLT applies the LT predicate on the "base_coin_id" field.
func BaseCoinIDLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldBaseCoinID, v))
}

// BaseCoinIDLTE applies the LTE predicate on the "base_coin_id" field.
func BaseCoinIDLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldBaseCoinID, v))
}

// BaseCoinIDContains applies the Contains predicate on the "base_coin_id" field.
func BaseCoinIDContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldBaseCoinID, v))
}

// BaseCoinIDHasPrefix applies the HasPrefix predicate on the "base_coin_id" field.
func BaseCoinIDHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldBaseCoinID, v))
}

// BaseCoinIDHasSuffix applies the HasSuffix predicate on the "base_coin_id" field.
func BaseCoinIDHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldBaseCoinID, v))
}

// BaseCoinIDIsNil applies the IsNil predicate on the "base_coin_id" field.
func BaseCoinIDIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldBaseCoinID))
}

// BaseCoinIDNotNil applies the NotNil predicate on the "base_coin_id" field.
func BaseCoinIDNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldBaseCoinID))
}

// BaseCoinIDEqualFold applies the EqualFold predicate on the "base_coin_id" field.
func BaseCoinIDEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldBaseCoinID, v))
}

// BaseCoinIDContainsFold applies the ContainsFold predicate on the "base_coin_id" field.
func BaseCoinIDContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldBaseCoinID, v))
}

// CounterEQ applies the EQ predicate on the "counter" field.
func CounterEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldCounter, v))
}

// CounterNEQ applies the NEQ predicate on the "counter" field.
func CounterNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldCounter, v))
}

// CounterIn applies the In predicate on the "counter" field.
func CounterIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldCounter, vs...))
}

// CounterNotIn applies the NotIn predicate on the "counter" field.
func CounterNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldCounter, vs...))
}

// CounterGT applies the GT predicate on the "counter" field.
func CounterGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldCounter, v))
}

// CounterGTE applies the GTE predicate on the "counter" field.
func CounterGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldCounter, v))
}

// CounterLT applies the LT predicate on the "counter" field.
func CounterLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldCounter, v))
}

// CounterLTE applies the LTE predicate on the "counter" field.
func CounterLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldCounter, v))
}

// CounterContains applies the Contains predicate on the "counter" field.
func CounterContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldCounter, v))
}

// CounterHasPrefix applies the HasPrefix predicate on the "counter" field.
func CounterHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldCounter, v))
}

// CounterHasSuffix applies the HasSuffix predicate on the "counter" field.
func CounterHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldCounter, v))
}

// CounterEqualFold applies the EqualFold predicate on the "counter" field.
func CounterEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldCounter, v))
}

// CounterContainsFold applies the ContainsFold predicate on the "counter" field.
func CounterContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldCounter, v))
}

// CounterCoinIDEQ applies the EQ predicate on the "counter_coin_id" field.
func CounterCoinIDEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldCounterCoinID, v))
}

// CounterCoinIDNEQ applies the NEQ predicate on the "counter_coin_id" field.
func CounterCoinIDNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldCounterCoinID, v))
}

// CounterCoinIDIn applies the In predicate on the "counter_coin_id" field.
func CounterCoinIDIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldCounterCoinID, vs...))
}

// CounterCoinIDNotIn applies the NotIn predicate on the "counter_coin_id" field.
func CounterCoinIDNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldCounterCoinID, vs...))
}

// CounterCoinIDGT applies the GT predicate on the "counter_coin_id" field.
func CounterCoinIDGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldCounterCoinID, v))
}

// CounterCoinIDGTE applies the GTE predicate on the "counter_coin_id" field.
func CounterCoinIDGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldCounterCoinID, v))
}

// CounterCoinIDLT applies the LT predicate on the "counter_coin_id" field.
func CounterCoinIDLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldCounterCoinID, v))
}

// CounterCoinIDLTE applies the LTE predicate on the "counter_coin_id" field.
func CounterCoinIDLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldCounterCoinID, v))
}

// CounterCoinIDContains applies the Contains predicate on the "counter_coin_id" field.
func CounterCoinIDContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldCounterCoinID, v))
}

// CounterCoinIDHasPrefix applies the HasPrefix predicate on the "counter_coin_id" field.
func CounterCoinIDHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldCounterCoinID, v))
}

// CounterCoinIDHasSuffix applies the HasSuffix predicate on the "counter_coin_id" field.
func CounterCoinIDHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldCounterCoinID, v))
}

// CounterCoinIDIsNil applies the IsNil predicate on the "counter_coin_id" field.
func CounterCoinIDIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldCounterCoinID))
}

// CounterCoinIDNotNil applies the NotNil predicate on the "counter_coin_id" field.
func CounterCoinIDNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldCounterCoinID))
}

// CounterCoinIDEqualFold applies the EqualFold predicate on the "counter_coin_id" field.
func CounterCoinIDEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldCounterCoinID, v))
}

// CounterCoinIDContainsFold applies the ContainsFold predicate on the "counter_coin_id" field.
func CounterCoinIDContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldCounterCoinID, v))
}

// MarketIsNil applies the IsNil predicate on the "market" field.
func MarketIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldMarket))
}

// MarketNotNil applies the NotNil predicate on the "market" field.
func MarketNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldMarket))
}

// LastEQ applies the EQ predicate on the "last" field.
func LastEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLast, v))
}

// LastNEQ applies the NEQ predicate on the "last" field.
func LastNEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldLast, v))
}

// LastIn applies the In predicate on the "last" field.
func LastIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldLast, vs...))
}

// LastNotIn applies the NotIn predicate on the "last" field.
func LastNotIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldLast, vs...))
}

// LastGT applies the GT predicate on the "last" field.
func LastGT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldLast, v))
}

// LastGTE applies the GTE predicate on the "last" field.
func LastGTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldLast, v))
}

// LastLT applies the LT predicate on the "last" field.
func LastLT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldLast, v))
}

// LastLTE applies the LTE predicate on the "last" field.
func LastLTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldLast, v))
}

// LastIsNil applies the IsNil predicate on the "last" field.
func LastIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldLast))
}

// LastNotNil applies the NotNil predicate on the "last" field.
func LastNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldLast))
}

// VolumeEQ applies the EQ predicate on the "volume" field.
func VolumeEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldVolume, v))
}

// VolumeNEQ applies the NEQ predicate on the "volume" field.
func VolumeNEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldVolume, v))
}

// VolumeIn applies the In predicate on the "volume" field.
func VolumeIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldVolume, vs...))
}

// VolumeNotIn applies the NotIn predicate on the "volume" field.
func VolumeNotIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldVolume, vs...))
}

// VolumeGT applies the GT predicate on the "volume" field.
func VolumeGT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldVolume, v))
}

// VolumeGTE applies the GTE predicate on the "volume" field.
func VolumeGTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldVolume, v))
}

// VolumeLT applies the LT predicate on the "volume" field.
func VolumeLT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldVolume, v))
}

// VolumeLTE applies the LTE predicate on the "volume" field.
func VolumeLTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldVolume, v))
}

// VolumeIsNil applies the IsNil predicate on the "volume" field.
func VolumeIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldVolume))
}

// VolumeNotNil applies the NotNil predicate on the "volume" field.
func VolumeNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldVolume))
}

// ConvertedLastIsNil applies the IsNil predicate on the "converted_last" field.
func ConvertedLastIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldConvertedLast))
}

// ConvertedLastNotNil applies the NotNil predicate on the "converted_last" field.
func ConvertedLastNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldConvertedLast))
}

// ConvertedVolumeIsNil applies the IsNil predicate on the "converted_volume" field.
func ConvertedVolumeIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldConvertedVolume))
}

// ConvertedVolumeNotNil applies the NotNil predicate on the "converted_volume" field.
func ConvertedVolumeNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldConvertedVolume))
}

// TrustScoreEQ applies the EQ predicate on the "trust_score" field.
func TrustScoreEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTrustScore, v))
}

// TrustScoreNEQ applies the NEQ predicate on the "trust_score" field.
func TrustScoreNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldTrustScore, v))
}

// TrustScoreIn applies the In predicate on the "trust_score" field.
func TrustScoreIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldTrustScore, vs...))
}

// TrustScoreNotIn applies the NotIn predicate on the "trust_score" field.
func TrustScoreNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldTrustScore, vs...))
}

// TrustScoreGT applies the GT predicate on the "trust_score" field.
func TrustScoreGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldTrustScore, v))
}

// TrustScoreGTE applies the GTE predicate on the "trust_score" field.
func TrustScoreGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldTrustScore, v))
}

// TrustScoreLT applies the LT predicate on the "trust_score" field.
func TrustScoreLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldTrustScore, v))
}

// TrustScoreLTE applies the LTE predicate on the "trust_score" field.
func TrustScoreLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldTrustScore, v))
}

// TrustScoreContains applies the Contains predicate on the "trust_score" field.
func TrustScoreContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldTrustScore, v))
}

// TrustScoreHasPrefix applies the HasPrefix predicate on the "trust_score" field.
func TrustScoreHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldTrustScore, v))
}

// TrustScoreHasSuffix applies the HasSuffix predicate on the "trust_score" field.
func TrustScoreHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldTrustScore, v))
}

// TrustScoreIsNil applies the IsNil predicate on the "trust_score" field.
func TrustScoreIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldTrustScore))
}

// TrustScoreNotNil applies the NotNil predicate on the "trust_score" field.
func TrustScoreNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldTrustScore))
}

// TrustScoreEqualFold applies the EqualFold predicate on the "trust_score" field.
func TrustScoreEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldTrustScore, v))
}

// TrustScoreContainsFold applies the ContainsFold predicate on the "trust_score" field.
func TrustScoreContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldTrustScore, v))
}

// BidAskSpreadPercentageEQ applies the EQ predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageNEQ applies the NEQ predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageNEQ(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageIn applies the In predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldBidAskSpreadPercentage, vs...))
}

// BidAskSpreadPercentageNotIn applies the NotIn predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageNotIn(vs ...float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldBidAskSpreadPercentage, vs...))
}

// BidAskSpreadPercentageGT applies the GT predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageGT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageGTE applies the GTE predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageGTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageLT applies the LT predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageLT(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageLTE applies the LTE predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageLTE(v float64) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldBidAskSpreadPercentage, v))
}

// BidAskSpreadPercentageIsNil applies the IsNil predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldBidAskSpreadPercentage))
}

// BidAskSpreadPercentageNotNil applies the NotNil predicate on the "bid_ask_spread_percentage" field.
func BidAskSpreadPercentageNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldBidAskSpreadPercentage))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldTimestamp, v))
}

// TimestampIsNil applies the IsNil predicate on the "timestamp" field.
func TimestampIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldTimestamp))
}

// TimestampNotNil applies the NotNil predicate on the "timestamp" field.
func TimestampNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldTimestamp))
}

// LastTradedAtEQ applies the EQ predicate on the "last_traded_at" field.
func LastTradedAtEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLastTradedAt, v))
}

// LastTradedAtNEQ applies the NEQ predicate on the "last_traded_at" field.
func LastTradedAtNEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldLastTradedAt, v))
}

// LastTradedAtIn applies the In predicate on the "last_traded_at" field.
func LastTradedAtIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldLastTradedAt, vs...))
}

// LastTradedAtNotIn applies the NotIn predicate on the "last_traded_at" field.
func LastTradedAtNotIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldLastTradedAt, vs...))
}

// LastTradedAtGT applies the GT predicate on the "last_traded_at" field.
func LastTradedAtGT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldLastTradedAt, v))
}

// LastTradedAtGTE applies the GTE predicate on the "last_traded_at" field.
func LastTradedAtGTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldLastTradedAt, v))
}

// LastTradedAtLT applies the LT predicate on the "last_traded_at" field.
func LastTradedAtLT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldLastTradedAt, v))
}

// LastTradedAtLTE applies the LTE predicate on the "last_traded_at" field.
func LastTradedAtLTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldLastTradedAt, v))
}

// LastTradedAtIsNil applies the IsNil predicate on the "last_traded_at" field.
func LastTradedAtIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldLastTradedAt))
}

// LastTradedAtNotNil applies the NotNil predicate on the "last_traded_at" field.
func LastTradedAtNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldLastTradedAt))
}

// LastFetchAtEQ applies the EQ predicate on the "last_fetch_at" field.
func LastFetchAtEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldLastFetchAt, v))
}

// LastFetchAtNEQ applies the NEQ predicate on the "last_fetch_at" field.
func LastFetchAtNEQ(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldLastFetchAt, v))
}

// LastFetchAtIn applies the In predicate on the "last_fetch_at" field.
func LastFetchAtIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldLastFetchAt, vs...))
}

// LastFetchAtNotIn applies the NotIn predicate on the "last_fetch_at" field.
func LastFetchAtNotIn(vs ...time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldLastFetchAt, vs...))
}

// LastFetchAtGT applies the GT predicate on the "last_fetch_at" field.
func LastFetchAtGT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldLastFetchAt, v))
}

// LastFetchAtGTE applies the GTE predicate on the "last_fetch_at" field.
func LastFetchAtGTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldLastFetchAt, v))
}

// LastFetchAtLT applies the LT predicate on the "last_fetch_at" field.
func LastFetchAtLT(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldLastFetchAt, v))
}

// LastFetchAtLTE applies the LTE predicate on the "last_fetch_at" field.
func LastFetchAtLTE(v time.Time) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldLastFetchAt, v))
}

// LastFetchAtIsNil applies the IsNil predicate on the "last_fetch_at" field.
func LastFetchAtIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldLastFetchAt))
}

// LastFetchAtNotNil applies the NotNil predicate on the "last_fetch_at" field.
func LastFetchAtNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldLastFetchAt))
}

// IsAnomalyEQ applies the EQ predicate on the "is_anomaly" field.
func IsAnomalyEQ(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldIsAnomaly, v))
}

// IsAnomalyNEQ applies the NEQ predicate on the "is_anomaly" field.
func IsAnomalyNEQ(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldIsAnomaly, v))
}

// IsAnomalyIsNil applies the IsNil predicate on the "is_anomaly" field.
func IsAnomalyIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldIsAnomaly))
}

// IsAnomalyNotNil applies the NotNil predicate on the "is_anomaly" field.
func IsAnomalyNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldIsAnomaly))
}

// IsStaleEQ applies the EQ predicate on the "is_stale" field.
func IsStaleEQ(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldIsStale, v))
}

// IsStaleNEQ applies the NEQ predicate on the "is_stale" field.
func IsStaleNEQ(v bool) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldIsStale, v))
}

// IsStaleIsNil applies the IsNil predicate on the "is_stale" field.
func IsStaleIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldIsStale))
}

// IsStaleNotNil applies the NotNil predicate on the "is_stale" field.
func IsStaleNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldIsStale))
}

// TradeURLEQ applies the EQ predicate on the "trade_url" field.
func TradeURLEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTradeURL, v))
}

// TradeURLNEQ applies the NEQ predicate on the "trade_url" field.
func TradeURLNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldTradeURL, v))
}

// TradeURLIn applies the In predicate on the "trade_url" field.
func TradeURLIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldTradeURL, vs...))
}

// TradeURLNotIn applies the NotIn predicate on the "trade_url" field.
func TradeURLNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldTradeURL, vs...))
}

// TradeURLGT applies the GT predicate on the "trade_url" field.
func TradeURLGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldTradeURL, v))
}

// TradeURLGTE applies the GTE predicate on the "trade_url" field.
func TradeURLGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldTradeURL, v))
}

// TradeURLLT applies the LT predicate on the "trade_url" field.
func TradeURLLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldTradeURL, v))
}

// TradeURLLTE applies the LTE predicate on the "trade_url" field.
func TradeURLLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldTradeURL, v))
}

// TradeURLContains applies the Contains predicate on the "trade_url" field.
func TradeURLContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldTradeURL, v))
}

// TradeURLHasPrefix applies the HasPrefix predicate on the "trade_url" field.
func TradeURLHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldTradeURL, v))
}

// TradeURLHasSuffix applies the HasSuffix predicate on the "trade_url" field.
func TradeURLHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldTradeURL, v))
}

// TradeURLIsNil applies the IsNil predicate on the "trade_url" field.
func TradeURLIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldTradeURL))
}

// TradeURLNotNil applies the NotNil predicate on the "trade_url" field.
func TradeURLNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldTradeURL))
}

// TradeURLEqualFold applies the EqualFold predicate on the "trade_url" field.
func TradeURLEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldTradeURL, v))
}

// TradeURLContainsFold applies the ContainsFold predicate on the "trade_url" field.
func TradeURLContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldTradeURL, v))
}

// TokenInfoURLEQ applies the EQ predicate on the "token_info_url" field.
func TokenInfoURLEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEQ(FieldTokenInfoURL, v))
}

// TokenInfoURLNEQ applies the NEQ predicate on the "token_info_url" field.
func TokenInfoURLNEQ(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNEQ(FieldTokenInfoURL, v))
}

// TokenInfoURLIn applies the In predicate on the "token_info_url" field.
func TokenInfoURLIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldIn(FieldTokenInfoURL, vs...))
}

// TokenInfoURLNotIn applies the NotIn predicate on the "token_info_url" field.
func TokenInfoURLNotIn(vs ...string) predicate.Ticker {
	return predicate.Ticker(sql.FieldNotIn(FieldTokenInfoURL, vs...))
}

// TokenInfoURLGT applies the GT predicate on the "token_info_url" field.
func TokenInfoURLGT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGT(FieldTokenInfoURL, v))
}

// TokenInfoURLGTE applies the GTE predicate on the "token_info_url" field.
func TokenInfoURLGTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldGTE(FieldTokenInfoURL, v))
}

// TokenInfoURLLT applies the LT predicate on the "token_info_url" field.
func TokenInfoURLLT(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLT(FieldTokenInfoURL, v))
}

// TokenInfoURLLTE applies the LTE predicate on the "token_info_url" field.
func TokenInfoURLLTE(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldLTE(FieldTokenInfoURL, v))
}

// TokenInfoURLContains applies the Contains predicate on the "token_info_url" field.
func TokenInfoURLContains(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContains(FieldTokenInfoURL, v))
}

// TokenInfoURLHasPrefix applies the HasPrefix predicate on the "token_info_url" field.
func TokenInfoURLHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasPrefix(FieldTokenInfoURL, v))
}

// TokenInfoURLHasSuffix applies the HasSuffix predicate on the "token_info_url" field.
func TokenInfoURLHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldHasSuffix(FieldTokenInfoURL, v))
}

// TokenInfoURLIsNil applies the IsNil predicate on the "token_info_url" field.
func TokenInfoURLIsNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldIsNull(FieldTokenInfoURL))
}

// TokenInfoURLNotNil applies the NotNil predicate on the "token_info_url" field.
func TokenInfoURLNotNil() predicate.Ticker {
	return predicate.Ticker(sql.FieldNotNull(FieldTokenInfoURL))
}

// TokenInfoURLEqualFold applies the EqualFold predicate on the "token_info_url" field.
func TokenInfoURLEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldEqualFold(FieldTokenInfoURL, v))
}

// TokenInfoURLContainsFold applies the ContainsFold predicate on the "token_info_url" field.
func TokenInfoURLContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(sql.FieldContainsFold(FieldTokenInfoURL, v))
}

// HasVenue applies the HasEdge predicate on the "venue" edge.
func HasVenue() predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Ticker
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenueWith applies the HasEdge predicate on the "venue" edge with a given conditions (other predicates).
func HasVenueWith(preds ...predicate.Venue) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VenueInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Ticker
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		p(s.Not())
	})
}
