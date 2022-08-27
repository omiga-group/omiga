package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar  pulsar.PulsarConfig `yaml:"pulsar"`
	Binance BinanceConfig       `yaml:"binance"`
}

type BinanceConfig struct {
	UseTestnet bool      `yaml:"useTestnet" env:"OMIGA_BINANCE_USETESTNET"`
	OrderBook  OrderBook `yaml:"orderBook" env:"OMIGA_FTX_ORDERBOOK"`
}

type OrderBook struct {
	Symbols []string `yaml:"symbols" env:"OMIGA_FTX_ORDERBOOK_SYMBOLS"`
}
