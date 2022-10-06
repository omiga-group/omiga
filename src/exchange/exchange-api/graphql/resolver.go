package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
)

type Resolver struct {
	client *entities.Client
}

func NewGraphQLServer(
	entgoClient entities.EntgoClient) (*handler.Server, error) {
	executableSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client: entgoClient.GetClient(),
		},
	})

	return handler.NewDefaultServer(executableSchema), nil
}
