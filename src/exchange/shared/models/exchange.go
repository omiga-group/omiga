package models

type Market struct {
	HasTradingIncentive bool
	Identifier          string
	Name                string
}

type ConvertedDetails struct {
	Btc float64
	Eth float64
	Usd float64
}
