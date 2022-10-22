package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	BitMart  BitMartConfig           `yaml:"bitmart"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type BitMartConfig struct {
	Id      string `yaml:"id" env:"OMIGA_BITMART_ID"`
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_BITMART_BASEURL"`
}
