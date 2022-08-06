package ftx

import "net/http"

const (
	baseWebSocketURL = "wss://ftx.com/ws/"
	baseAPIURL       = "https://ftx.com/api/"
)

func NewClient(c Config, h http.Client) (Client, error) {
	cfg := Client{
		config: c,
		http:   h,
	}

	return cfg, nil
}

type Client struct {
	config Config
	http   http.Client
}
