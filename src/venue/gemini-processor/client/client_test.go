package client

import (
	"testing"

	"github.com/omiga-group/omiga/src/venue/gemini-processor/configuration"
)

func TestGetMarkets(t *testing.T) {
	cfg := configuration.GeminiConfig{
		ApiUrl: "https://api.sandbox.gemini.com/v1",
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
