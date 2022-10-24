package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Ftx      FtxConfig               `yaml:"ftx"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
}

type FtxConfig struct {
	Id           string    `yaml:"id" env:"OMIGA_FTX_ID"`
	ApiUrl       string    `yaml:"apiUrl" env:"OMIGA_FTX_APIURL"`
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_FTX_WEBSOCKETURL"`
	ApiKey       string    `yaml:"apiKey" env:"OMIGA_FTX_APIKEY"`
	Passphrase   string    `yaml:"passphrase" env:"OMIGA_FTX_PASSPHRASE"`
	SecretKey    string    `yaml:"secretKey" env:"OMIGA_FTX_SECRETKEY"`
	Timeout      int       `yaml:"timeout"`
	OrderBook    OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Pairs []PairConfig `yaml:"pairs"`
}

type PairConfig struct {
	Pair string `yaml:"pair"`
}
