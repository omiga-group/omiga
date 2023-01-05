// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/venue-api/graphql"
	"github.com/omiga-group/omiga/src/venue/venue-api/http"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewHttpServer(logger *zap.SugaredLogger, appConfig configuration.AppConfig, entgoClient entities.EntgoClient) (http.HttpServer, error) {
	server, err := graphql.NewGraphQLServer(entgoClient)
	if err != nil {
		return nil, err
	}
	httpServer, err := http.NewHttpServer(logger, appConfig, server)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}
