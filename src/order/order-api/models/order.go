package models

import (
	"github.com/google/uuid"
)

type Order struct {
	Id      int
	OrderID uuid.UUID
}
