package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
