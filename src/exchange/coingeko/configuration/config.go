package configuration

import "github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"

type Config struct {
	Coingeko  CoingekoConfig          `yaml:"coingeko"`
	Postgres  postgres.PostgresConfig `yaml:"postgres"`
	Exchanges map[string]Exchange     `yaml:"exchanges"`
}

type CoingekoConfig struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_COINGEKO_BASEURL"`
}

type Exchange struct {
	MakerFee   float64 `yaml:"makerFee"`
	TakerFee   float64 `yaml:"takerFee"`
	SpreadFee  bool    `yaml:"spreadFee"`
	SupportAPI bool    `yaml:"supportAPI"`
}
