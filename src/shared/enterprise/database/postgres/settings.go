package postgres

const ConfigKey = "postgres"

type PostgresSettings struct {
	ConnectionString string
	MaxOpenConns     int
}
