package models

import (
	krakenwebsocket "github.com/aopoltorzhicky/go_kraken/websocket"
)

type KrakenOrderBookEntry struct {
	Symbol string
	Bid    *krakenwebsocket.OrderBookItem
	Ask    *krakenwebsocket.OrderBookItem
}
