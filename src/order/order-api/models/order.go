package models

import (
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
)

type Order struct {
	Id      int
	OrderId orderv1.ID
}
