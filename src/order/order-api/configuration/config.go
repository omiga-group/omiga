package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Outbox   outbox.OutboxConfig     `yaml:"outbox"`
}
