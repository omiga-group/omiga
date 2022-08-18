package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Postgres postgres.PostgresConfig `yaml:"postgres"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
}
