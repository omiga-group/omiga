package subscribers

import (
	"encoding/json"
	"math"
	"time"
)

type geminiTime struct {
	Time time.Time
}

func (p *geminiTime) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	sec, nsec := math.Modf(f)
	p.Time = time.Unix(int64(sec), int64(nsec))
	return nil
}

type operationType string

const (
	OperationTypeSubscribe   operationType = "subscribe"
	OperationTypeUnsubscribe operationType = "unsubscribe"
)

type geminiRequest struct {
	Op      operationType `json:"op"`
	Channel *string       `json:"channel,omitempty"`
	Market  string        `json:"market,omitempty"`
}

type responseType string

const (
	ResponseTypeError        responseType = "error"
	ResponseTypeSubscribed   responseType = "subscribed"
	ResponseTypeUnsubscribed responseType = "unsubscribed"
	ResponseTypeInfo         responseType = "info"
	ResponseTypePartial      responseType = "partial"
	ResponseTypeUpdate       responseType = "update"
)

type geminiOrderBook struct {
	Channel string              `json:"channel"`
	Market  string              `json:"market"`
	Type    responseType        `json:"type"`
	Data    *geminirderBookData `json:"data,omitempty"`
}

type geminirderBookData struct {
	Time     geminiTime   `json:"time"`
	Checksum int          `json:"checksum"`
	Bids     [][2]float64 `json:"bids"`
	Asks     [][2]float64 `json:"asks"`
	Action   string       `json:"action"`
}
