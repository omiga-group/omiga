
// Code generated by go-omiga-template, DO NOT EDIT.

package orderv1

import (
	"context"
)

type Subscriber interface {
  Handle(ctx context.Context, event DomainEvent) error
}
