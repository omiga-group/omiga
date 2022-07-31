package ftx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/omiga-group/omiga/code/exchange/ftx-processor/integration/entities"
)

func (c Client) SyncOrderBooks(market string) error {
	ws, _, err := websocket.DefaultDialer.Dial(baseWebSocketURL, nil)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New("websocket is not initialized")
	}

	query := []byte(fmt.Sprintf(`{"op": "subscribe", "channel": "trades", "market": "%s"}`, market))
	ws.WriteMessage(websocket.TextMessage, query)
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err) // TODO:log
		}

		var orderBook entities.OrderBook
		if err = json.Unmarshal(msg, &orderBook); err != nil {
			fmt.Println(err) // TODO:log
		}

		fmt.Println("%+v\n", orderBook)
	}
}
