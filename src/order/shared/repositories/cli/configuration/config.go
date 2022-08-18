package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
)

type Config struct {
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
