package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar pulsar.PulsarConfig `yaml:"pulsar"`
	Ftx    FtxSettings         `yaml:"ftx"`
}

type FtxSettings struct {
	WebsocketUrl string    `yaml:"websocketUrl" env:"OMIGA_FTX_WEBSOCKETURL"`
	OrderBook    OrderBook `yaml:"orderBook" env:"OMIGA_FTX_ORDERBOOK"`
}

type OrderBook struct {
	Markets []string `yaml:"markets" env:"OMIGA_FTX_ORDERBOOK_MARKETS"`
}
