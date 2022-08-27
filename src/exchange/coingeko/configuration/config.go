package configuration

import "github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"

type Config struct {
	Coingeko CoingekoSettings        `yaml:"coingeko"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type CoingekoSettings struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_COINGEKO_BASEURL"`
}
