package postgres

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/viper"
)

const ConfigKey = "postgres"

type PostgresSettings struct {
	ConnectionString string `json:"connectionString"`
	MaxOpenConns     int    `json:"maxOpenConns"`
}

func GetPostgresSettings(viper *viper.Viper) PostgresSettings {
	key := ConfigKey + configuration.KeyDelimiter

	return PostgresSettings{
		ConnectionString: viper.GetString(key + "connectionString"),
		MaxOpenConns:     viper.GetInt(key + "maxOpenConns"),
	}
}
