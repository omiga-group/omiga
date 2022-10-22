package client

import (
	"testing"

	"github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
)

func TestGetMarkets(t *testing.T) {
	cfg := configuration.FtxConfig{
		ApiUrl: "https://ftx.com/api",
	}
	c := NewFtxApiClient(cfg)
	mm, err := c.GetMarkets()
	if err != nil {
		t.Error(err)
	}
	if len(mm) < 1 {
		t.Error("empty market maps")
	}
}
