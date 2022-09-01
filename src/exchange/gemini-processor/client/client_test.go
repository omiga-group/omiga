package client

import (
	"testing"

	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
)

func TestGetMarkets(t *testing.T) {
	cfg := configuration.GeminiConfig{
		ApiUrl: "https://gemini.com/api",
	}
	c := NewGeminiApiClient(cfg)
	mm, err := c.GetMarkets()
	if err != nil {
		t.Error(err)
	}
	if len(mm) < 1 {
		t.Error("empty market maps")
	}
}
