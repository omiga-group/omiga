package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App     configuration.AppConfig `yaml:"app"`
	Pulsar  pulsar.PulsarConfig     `yaml:"pulsar"`
	Binance BinanceConfig           `yaml:"binance"`
}

type BinanceConfig struct {
	UseTestnet bool      `yaml:"useTestnet" env:"OMIGA_BINANCE_USETESTNET"`
	OrderBook  OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Symbols []SymbolConfig `yaml:"symbols"`
}

type SymbolConfig struct {
	Symbol    string `yaml:"symbol"`
	PurgeTime string `yaml:"purgeTime"`
}
