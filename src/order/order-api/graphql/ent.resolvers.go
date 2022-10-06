package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/order/order-api/graphql/generated"
	"github.com/omiga-group/omiga/src/order/order-api/graphql/models"
	"github.com/omiga-group/omiga/src/order/shared/entities"
)

// Status is the resolver for the status field.
func (r *outboxWhereInputResolver) Status(ctx context.Context, obj *entities.OutboxWhereInput, data *models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusNeq is the resolver for the statusNEQ field.
func (r *outboxWhereInputResolver) StatusNeq(ctx context.Context, obj *entities.OutboxWhereInput, data *models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusIn is the resolver for the statusIn field.
func (r *outboxWhereInputResolver) StatusIn(ctx context.Context, obj *entities.OutboxWhereInput, data []models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// StatusNotIn is the resolver for the statusNotIn field.
func (r *outboxWhereInputResolver) StatusNotIn(ctx context.Context, obj *entities.OutboxWhereInput, data []models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented"))
}

// OutboxWhereInput returns generated.OutboxWhereInputResolver implementation.
func (r *Resolver) OutboxWhereInput() generated.OutboxWhereInputResolver {
	return &outboxWhereInputResolver{r}
}

type outboxWhereInputResolver struct{ *Resolver }
