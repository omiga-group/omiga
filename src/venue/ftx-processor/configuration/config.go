package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	exchangeConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
)

type Config struct {
	App      configuration.AppConfig              `yaml:"app"`
	Exchange exchangeConfiguration.ExchangeConfig `yaml:"exchange"`
	Pulsar   pulsar.PulsarConfig                  `yaml:"pulsar"`
	Ftx      FtxConfig                            `yaml:"ftx"`
}

type FtxConfig struct {
	ApiUrl       string    `yaml:"apiUrl" env:"OMIGA_FTX_APIURL"`
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_FTX_WEBSOCKETURL"`
	Timeout      int       `yaml:"timeout"`
	OrderBook    OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Markets []MarketConfig `yaml:"markets"`
}

type MarketConfig struct {
	Market string `yaml:"market"`
}
