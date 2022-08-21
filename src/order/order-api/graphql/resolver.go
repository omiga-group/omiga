package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/omiga-group/omiga/src/order/order-api/graphql/generated"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

type Resolver struct {
	client                       *repositories.Client
	orderService                 services.OrderService
	orderOutboxBackgroundService outbox.OutboxBackgroundService
}

type GraphQLServer interface {
}

func NewGraphQLServer(
	entgoClient repositories.EntgoClient,
	orderService services.OrderService,
	orderOutboxBackgroundService outbox.OutboxBackgroundService) (*handler.Server, error) {
	executableSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:                       entgoClient.GetClient(),
			orderService:                 orderService,
			orderOutboxBackgroundService: orderOutboxBackgroundService,
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
