package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Xt       XtConfig                `yaml:"xt"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type XtConfig struct {
	Id        string `yaml:"id" env:"OMIGA_XT_ID"`
	BaseUrl   string `yaml:"baseUrl" env:"OMIGA_XT_BASEURL"`
	ApiKey    string `yaml:"apiKey" env:"OMIGA_XT_APIKEY"`
	SecretKey string `yaml:"secretKey" env:"OMIGA_XT_SECRETKEY"`
}
