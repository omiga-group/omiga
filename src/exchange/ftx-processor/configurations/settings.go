package configurations

const ConfigKey = "ftx"

type FtxSettings struct {
	WebsocketUrl string
	OrderBook    OrderBook
}

type OrderBook struct {
	Markets []string
}
