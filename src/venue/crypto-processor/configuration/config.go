package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Crypto   CryptoConfig            `yaml:"crypto"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type CryptoConfig struct {
	Id      string `yaml:"id" env:"OMIGA_CRYPTO_ID"`
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_CRYPTO_BASEURL"`
}
