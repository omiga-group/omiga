package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	DexTrade DexTradeConfig          `yaml:"dextrade"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type DexTradeConfig struct {
	Id        string `yaml:"id" env:"OMIGA_DEXTRADE_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_DEXTRADE_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_DEXTRADE_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_DEXTRADE_SECRETKEY"`
}
