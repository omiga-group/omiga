package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

type Resolver struct {
	client *repositories.Client
}

type GraphQLServer interface {
}

func NewGraphQLServer(entgoClient repositories.EntgoClient) (*handler.Server, error) {
	client, err := entgoClient.GetClient()
	if err != nil {
		return nil, err
	}

	executableSchema := shared.NewExecutableSchema(shared.Config{
		Resolvers: &Resolver{
			client: client,
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
