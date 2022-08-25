package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar  pulsar.PulsarConfig `yaml:"pulsar"`
	Binance BinanceSettings     `yaml:"binance"`
}

type BinanceSettings struct {
	UseTestnet bool      `json:"useTestnet" env:"OMIGA_BINANCE_USETESTNET"`
	OrderBook  OrderBook `json:"orderBook" env:"OMIGA_FTX_ORDERBOOK"`
}

type OrderBook struct {
	Symbols []string `json:"symbols" env:"OMIGA_FTX_ORDERBOOK_SYMBOLS"`
}
