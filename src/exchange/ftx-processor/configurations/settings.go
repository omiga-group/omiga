package configurations

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/viper"
)

const ConfigKey = "ftx"

type FtxSettings struct {
	WebsocketUrl string    `json:"websocketUrl"`
	OrderBook    OrderBook `json:"orderBook"`
}

type OrderBook struct {
	Markets []string `json:"markets"`
}

func GetFtxSettings(viper *viper.Viper) FtxSettings {
	key := ConfigKey + configuration.KeyDelimiter

	return FtxSettings{
		WebsocketUrl: viper.GetString(key + "websocketUrl"),
		OrderBook: OrderBook{
			Markets: viper.GetStringSlice(key + "orderBook" + configuration.KeyDelimiter + "markets"),
		},
	}
}
