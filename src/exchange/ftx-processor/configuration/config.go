package configuration

import (
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig              `yaml:"app"`
	Exchange exchangeConfiguration.ExchangeConfig `yaml:"exchange"`
	Pulsar   pulsar.PulsarConfig                  `yaml:"pulsar"`
	Ftx      FtxConfig                            `yaml:"ftx"`
	Postgres postgres.PostgresConfig              `yaml:"postgres"`
}

type FtxConfig struct {
	ApiUrl       string    `yaml:"apiUrl" env:"OMIGA_FTX_APIURL"`
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_FTX_WEBSOCKETURL"`
	Timeout      int       `yaml:"timeout"`
	OrderBook    OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}
