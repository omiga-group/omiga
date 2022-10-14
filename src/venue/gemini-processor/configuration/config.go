package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	exchangeConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
)

type Config struct {
	App      configuration.AppConfig              `yaml:"app"`
	Exchange exchangeConfiguration.ExchangeConfig `yaml:"exchange"`
	Pulsar   pulsar.PulsarConfig                  `yaml:"pulsar"`
	Gemini   GeminiConfig                         `yaml:"gemini"`
	Postgres postgres.PostgresConfig              `yaml:"postgres"`
}

type GeminiConfig struct {
	ApiUrl       string    `yaml:"apiUrl" env:"OMIGA_GEMINI_APIURL"`
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_GEMINI_WEBSOCKETURL"`
	Timeout      int       `yaml:"timeout" env:"OMIGA_GEMINI_TIMEOUT"`
	OrderBook    OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Markets []MarketConfig `yaml:"markets"`
}

type MarketConfig struct {
	Market string `yaml:"market"`
}
