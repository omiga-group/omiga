package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar pulsar.PulsarConfig `yaml:"pulsar"`
	Ftx    FtxConfig           `yaml:"ftx"`
}

type FtxConfig struct {
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_FTX_WEBSOCKETURL"`
	OrderBook    OrderBook `yaml:"orderBook"`
}

type OrderBook struct {
	Markets []MarketConfig `yaml:"markets"`
}

type MarketConfig struct {
	Market string `yaml:"market"`
}
