package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	venueConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
)

type Config struct {
	App      configuration.AppConfig        `yaml:"app"`
	Venue    venueConfiguration.VenueConfig `yaml:"venue"`
	Pulsar   pulsar.PulsarConfig            `yaml:"pulsar"`
	KuCoin   KuCoinConfig                   `yaml:"kucoin"`
	Postgres postgres.PostgresConfig        `yaml:"postgres"`
}

type KuCoinConfig struct {
	ApiKey     string `yaml:"apiKey" env:"OMIGA_KUCOIN_APIKEY"`
	Passphrase string `yaml:"passphrase" env:"OMIGA_KUCOIN_PASSPHRASE"`
	SecretKey  string `yaml:"secretKey" env:"OMIGA_KUCOIN_SECRETKEY"`
}
