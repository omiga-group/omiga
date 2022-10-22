package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Bittrex  BittrexConfig           `yaml:"bittrex"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type BittrexConfig struct {
	Id        string `yaml:"id" env:"OMIGA_BITTREX_ID"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_BITTREX_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_BITTREX_SECRETKEY"`
}
