package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Huobi    HuobiConfig             `yaml:"huobi"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type HuobiConfig struct {
	Id        string `yaml:"id" env:"OMIGA_HUOBI_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_HUOBI_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_HUOBI_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_HUOBI_SECRETKEY"`
}
