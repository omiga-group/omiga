package postgres

type PostgresConfig struct {
	ConnectionString string `yaml:"connectionString",env:"OMIGA_POSTGRES_CONNECTIONSTRING"`
	MaxOpenConns     int    `yaml:"maxOpenConns",env:"OMIGA_POSTGRES_MAXOPENCONNS"`
}
