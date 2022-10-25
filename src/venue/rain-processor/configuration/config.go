package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Rain     RainConfig              `yaml:"rain"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type RainConfig struct {
	Id        string `yaml:"id" env:"OMIGA_RAIN_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_RAIN_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_RAIN_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_RAIN_SECRETKEY"`
}
