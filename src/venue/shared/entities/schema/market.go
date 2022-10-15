package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Market holds the schema definition for the Market entity.
type Market struct {
	ent.Schema
}

// Fields of the Market.
func (Market) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entgql.OrderField("name")),

		field.Enum("type").
			NamedValues(
				"SPOT_TRADING", "SPOT_TRADING",
				"MARGIN_TRADING", "MARGIN_TRADING",
				"DERIVATIVES", "DERIVATIVES",
				"EARN", "EARN",
				"PERPETUAL", "PERPETUAL",
				"FUTURES", "FUTURES",
				"WARRANT", "WARRANT",
				"OTC", "OTC",
				"YIELD", "YIELD",
				"P2P", "P2P",
				"STRATEGY_TRADING", "STRATEGY_TRADING",
				"SWAP_FARMING", "SWAP_FARMING",
				"FAN_TOKEN", "FAN_TOKEN",
				"ETF", "ETF",
				"NFT", "NFT",
				"Swap", "SWAP",
				"LIQUIDITY", "LIQUIDITY").
			Annotations(entgql.OrderField("type")),
	}
}

// Edges of the Market.
func (Market) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).
			Ref("market").
			Unique().
			Required(),
		edge.To("trading_pair", TradingPair.Type),
	}
}

func (Market) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("type"),
	}
}
