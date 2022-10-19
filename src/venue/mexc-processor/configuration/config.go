package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Mexc     MexcConfig              `yaml:"mexc"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type MexcConfig struct {
	Id        string `yaml:"id" env:"OMIGA_MEXC_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_MEXC_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_MEXC_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_MEXC_SECRETKEY"`
}
