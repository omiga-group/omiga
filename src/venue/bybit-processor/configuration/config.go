package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Bybit    BybitConfig             `yaml:"bybit"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type BybitConfig struct {
	Id        string `yaml:"id" env:"OMIGA_BYBIT_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_BYBIT_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_BYBIT_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_BYBIT_SECRETKEY"`
}
