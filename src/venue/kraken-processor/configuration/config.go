package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Kraken   KrakenConfig            `yaml:"kraken"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type KrakenConfig struct {
	Id        string    `yaml:"id" env:"OMIGA_KRAKEN_ID"`
	ApiKey    string    `yaml:"apiKey" env:"OMIGA_KRAKEN_APIKEY"`
	SecretKey string    `yaml:"secretKey" env:"OMIGA_KRAKEN_SECRETKEY"`
	OrderBook OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}
