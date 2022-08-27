package configuration

import "github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"

type Config struct {
	Coingeko CoingekoConfig          `yaml:"coingeko"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type CoingekoConfig struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_COINGEKO_BASEURL"`
}
