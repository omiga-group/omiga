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
	OrderBook OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}