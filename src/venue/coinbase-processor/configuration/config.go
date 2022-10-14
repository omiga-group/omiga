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
	Coinbase CoinbaseConfig                 `yaml:"coinbase"`
	Postgres postgres.PostgresConfig        `yaml:"postgres"`
}

type CoinbaseConfig struct {
	UseSandbox bool   `yaml:"useSandbox" env:"OMIGA_COINBASE_USESANDBOX"`
	BaseUrl    string `yaml:"baseUrl" env:"OMIGA_COINBASE_BASEURL"`
	ApiKey     string `yaml:"apiKey" env:"OMIGA_COINBASE_APIKEY"`
	Passphrase string `yaml:"passphrase" env:"OMIGA_COINBASE_PASSPHRASE"`
	SecretKey  string `yaml:"secretKey" env:"OMIGA_COINBASE_SECRETKEY"`
}
