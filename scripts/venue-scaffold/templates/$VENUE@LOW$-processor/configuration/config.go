package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	$VENUE@PAS$  $VENUE@PAS$Config           `yaml:"$VENUE@LOW$"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type $VENUE@PAS$Config struct {
	Id         string    `yaml:"id" env:"OMIGA_$VENUE@UP$_ID"`
	ApiKey     string    `yaml:"apiKey" env:"OMIGA_$VENUE@UP$_APIKEY"`
	SecretKey  string    `yaml:"secretKey" env:"OMIGA_$VENUE@UP$_SECRETKEY"`
	UseTestnet bool      `yaml:"useTestnet" env:"OMIGA_$VENUE@UP$_USETESTNET"`
	OrderBook  OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}
