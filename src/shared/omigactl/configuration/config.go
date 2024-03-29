package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
