// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package shared

import (
	"fmt"
	"io"
	"strconv"

	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

type CancelOrderInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               int     `json:"id"`
}

type OrderPayload struct {
	ClientMutationID *string             `json:"clientMutationId"`
	Order            *repositories.Order `json:"order"`
}

type SubmitOrderInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               int     `json:"id"`
}

type OutboxStatus string

const (
	OutboxStatusPending   OutboxStatus = "PENDING"
	OutboxStatusSucceeded OutboxStatus = "SUCCEEDED"
	OutboxStatusFailed    OutboxStatus = "FAILED"
)

var AllOutboxStatus = []OutboxStatus{
	OutboxStatusPending,
	OutboxStatusSucceeded,
	OutboxStatusFailed,
}

func (e OutboxStatus) IsValid() bool {
	switch e {
	case OutboxStatusPending, OutboxStatusSucceeded, OutboxStatusFailed:
		return true
	}
	return false
}

func (e OutboxStatus) String() string {
	return string(e)
}

func (e *OutboxStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OutboxStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OutboxStatus", str)
	}
	return nil
}

func (e OutboxStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
