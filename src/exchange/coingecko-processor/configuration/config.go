package configuration

import "github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"

type Config struct {
	Coingecko CoingeckoConfig         `yaml:"coingecko"`
	Postgres  postgres.PostgresConfig `yaml:"postgres"`
	Exchanges map[string]Exchange     `yaml:"exchanges"`
}

type CoingeckoConfig struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_COINGECKO_BASEURL"`
}

type Exchange struct {
	MakerFee   float64 `yaml:"makerFee"`
	TakerFee   float64 `yaml:"takerFee"`
	SpreadFee  bool    `yaml:"spreadFee"`
	SupportAPI bool    `yaml:"supportAPI"`
}
