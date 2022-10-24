package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Bitmart  BitmartConfig           `yaml:"bitmart"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type BitmartConfig struct {
	Id      string `yaml:"id" env:"OMIGA_BITMART_ID"`
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_BITMART_BASEURL"`
}
