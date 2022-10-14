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
	BitMart  BitMartConfig                  `yaml:"bitmart"`
	Postgres postgres.PostgresConfig        `yaml:"postgres"`
}

type BitMartConfig struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_BITMART_BASEURL"`
}
