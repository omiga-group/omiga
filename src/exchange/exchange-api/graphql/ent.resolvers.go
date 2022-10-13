package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/models"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
)

// Type is the resolver for the type field.
func (r *marketWhereInputResolver) Type(ctx context.Context, obj *entities.MarketWhereInput, data *models.MarketType) error {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// TypeNeq is the resolver for the typeNEQ field.
func (r *marketWhereInputResolver) TypeNeq(ctx context.Context, obj *entities.MarketWhereInput, data *models.MarketType) error {
	panic(fmt.Errorf("not implemented: TypeNeq - typeNEQ"))
}

// TypeIn is the resolver for the typeIn field.
func (r *marketWhereInputResolver) TypeIn(ctx context.Context, obj *entities.MarketWhereInput, data []models.MarketType) error {
	panic(fmt.Errorf("not implemented: TypeIn - typeIn"))
}

// TypeNotIn is the resolver for the typeNotIn field.
func (r *marketWhereInputResolver) TypeNotIn(ctx context.Context, obj *entities.MarketWhereInput, data []models.MarketType) error {
	panic(fmt.Errorf("not implemented: TypeNotIn - typeNotIn"))
}

// Status is the resolver for the status field.
func (r *outboxWhereInputResolver) Status(ctx context.Context, obj *entities.OutboxWhereInput, data *models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// StatusNeq is the resolver for the statusNEQ field.
func (r *outboxWhereInputResolver) StatusNeq(ctx context.Context, obj *entities.OutboxWhereInput, data *models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented: StatusNeq - statusNEQ"))
}

// StatusIn is the resolver for the statusIn field.
func (r *outboxWhereInputResolver) StatusIn(ctx context.Context, obj *entities.OutboxWhereInput, data []models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented: StatusIn - statusIn"))
}

// StatusNotIn is the resolver for the statusNotIn field.
func (r *outboxWhereInputResolver) StatusNotIn(ctx context.Context, obj *entities.OutboxWhereInput, data []models.OutboxStatus) error {
	panic(fmt.Errorf("not implemented: StatusNotIn - statusNotIn"))
}

// MarketWhereInput returns generated.MarketWhereInputResolver implementation.
func (r *Resolver) MarketWhereInput() generated.MarketWhereInputResolver {
	return &marketWhereInputResolver{r}
}

// OutboxWhereInput returns generated.OutboxWhereInputResolver implementation.
func (r *Resolver) OutboxWhereInput() generated.OutboxWhereInputResolver {
	return &outboxWhereInputResolver{r}
}

type marketWhereInputResolver struct{ *Resolver }
type outboxWhereInputResolver struct{ *Resolver }
