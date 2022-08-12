package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

// OrderID is the resolver for the orderID field.
func (r *orderWhereInputResolver) OrderID(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDNeq is the resolver for the orderIDNEQ field.
func (r *orderWhereInputResolver) OrderIDNeq(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDIn is the resolver for the orderIDIn field.
func (r *orderWhereInputResolver) OrderIDIn(ctx context.Context, obj *repositories.OrderWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDNotIn is the resolver for the orderIDNotIn field.
func (r *orderWhereInputResolver) OrderIDNotIn(ctx context.Context, obj *repositories.OrderWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDGt is the resolver for the orderIDGT field.
func (r *orderWhereInputResolver) OrderIDGt(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDGte is the resolver for the orderIDGTE field.
func (r *orderWhereInputResolver) OrderIDGte(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDLt is the resolver for the orderIDLT field.
func (r *orderWhereInputResolver) OrderIDLt(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderIDLte is the resolver for the orderIDLTE field.
func (r *orderWhereInputResolver) OrderIDLte(ctx context.Context, obj *repositories.OrderWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// Status is the resolver for the status field.
func (r *outboxWhereInputResolver) Status(ctx context.Context, obj *repositories.OutboxWhereInput, data *shared.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusNeq is the resolver for the statusNEQ field.
func (r *outboxWhereInputResolver) StatusNeq(ctx context.Context, obj *repositories.OutboxWhereInput, data *shared.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusIn is the resolver for the statusIn field.
func (r *outboxWhereInputResolver) StatusIn(ctx context.Context, obj *repositories.OutboxWhereInput, data []shared.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusNotIn is the resolver for the statusNotIn field.
func (r *outboxWhereInputResolver) StatusNotIn(ctx context.Context, obj *repositories.OutboxWhereInput, data []shared.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// OrderWhereInput returns shared.OrderWhereInputResolver implementation.
func (r *Resolver) OrderWhereInput() shared.OrderWhereInputResolver {
	return &orderWhereInputResolver{r}
}

// OutboxWhereInput returns shared.OutboxWhereInputResolver implementation.
func (r *Resolver) OutboxWhereInput() shared.OutboxWhereInputResolver {
	return &outboxWhereInputResolver{r}
}

type orderWhereInputResolver struct{ *Resolver }
type outboxWhereInputResolver struct{ *Resolver }
