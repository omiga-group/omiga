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
	Huobi    HuobiConfig                          `yaml:"huobi"`
	Postgres postgres.PostgresConfig              `yaml:"postgres"`
}

type HuobiConfig struct {
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_HUOBI_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_HUOBI_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_HUOBI_SECRETKEY"`
}
