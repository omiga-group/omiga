package models

import (
	currencyrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/currency"
)

type Currency struct {
	Symbol string
	Name   *string
	Type   currencyrepo.Type
}
