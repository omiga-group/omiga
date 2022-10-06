package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/omiga-group/omiga/src/order/order-api/graphql/generated"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/entities"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
)

type Resolver struct {
	client                  *entities.Client
	orderService            services.OrderService
	outboxBackgroundService outbox.OutboxBackgroundService
}

type GraphQLServer interface {
}

func NewGraphQLServer(
	entgoClient entities.EntgoClient,
	orderService services.OrderService,
	outboxBackgroundService outbox.OutboxBackgroundService) (*handler.Server, error) {
	executableSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:                  entgoClient.GetClient(),
			orderService:            orderService,
			outboxBackgroundService: outboxBackgroundService,
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
