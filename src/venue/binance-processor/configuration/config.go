package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Binance  BinanceConfig           `yaml:"binance"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type BinanceConfig struct {
	Id         string    `yaml:"id" env:"OMIGA_BINANCE_ID"`
	ApiKey     string    `yaml:"apiKey" env:"OMIGA_BINANCE_APIKEY"`
	SecretKey  string    `yaml:"secretKey" env:"OMIGA_BINANCE_SECRETKEY"`
	UseTestnet bool      `yaml:"useTestnet" env:"OMIGA_BINANCE_USETESTNET"`
	OrderBook  OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}
