package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	KuCoin   KuCoinConfig            `yaml:"kucoin"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type KuCoinConfig struct {
	Id         string `yaml:"id" env:"OMIGA_KUCOIN_ID"`
	ApiKey     string `yaml:"apiKey" env:"OMIGA_KUCOIN_APIKEY"`
	Passphrase string `yaml:"passphrase" env:"OMIGA_KUCOIN_PASSPHRASE"`
	SecretKey  string `yaml:"secretKey" env:"OMIGA_KUCOIN_SECRETKEY"`
}
