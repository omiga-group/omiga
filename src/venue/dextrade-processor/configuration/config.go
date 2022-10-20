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
	Id         string `yaml:"id" env:"OMIGA_DEXTRADE_ID"`
	ApiKey     string `yaml:"apiKey" env:"OMIGA_DEXTRADE_APIKEY"`
	Passphrase string `yaml:"passphrase" env:"OMIGA_DEXTRADE_PASSPHRASE"`
	SecretKey  string `yaml:"secretKey" env:"OMIGA_DEXTRADE_SECRETKEY"`
}
