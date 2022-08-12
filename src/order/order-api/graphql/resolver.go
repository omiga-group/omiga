package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

type Resolver struct {
	client       *repositories.Client
	orderService services.OrderService
}

type GraphQLServer interface {
}

func NewGraphQLServer(
	entgoClient repositories.EntgoClient,
	orderService services.OrderService) (*handler.Server, error) {
	executableSchema := shared.NewExecutableSchema(shared.Config{
		Resolvers: &Resolver{
			client:       entgoClient.GetClient(),
			orderService: orderService,
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
