package business

import (
	"net/http"

	"github.com/omiga-group/omiga/code/exchange/ftx-processor/integration/ftx"
	"github.com/spf13/viper"
)

func Foo() {
	cfg := ftx.Config{
		APIKey: viper.GetString("integration.api-key"),
	}

	c, err := ftx.NewClient(cfg, *http.DefaultClient)
	if err != nil {
		panic(err)
	}

	c.GetMarkets()
	go c.SyncOrderBooks("ETH-PERP")
}
