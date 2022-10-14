package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	exchangeConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
)

type Config struct {
	App      configuration.AppConfig              `yaml:"app"`
	Exchange exchangeConfiguration.ExchangeConfig `yaml:"exchange"`
	Pulsar   pulsar.PulsarConfig                  `yaml:"pulsar"`
	Bittrex  BittrexConfig                        `yaml:"bittrex"`
	Postgres postgres.PostgresConfig              `yaml:"postgres"`
}

type BittrexConfig struct {
	ApiKey    string `yaml:"apiKey" env:"OMIGA_BITTREX_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_BITTREX_SECRETKEY"`
}