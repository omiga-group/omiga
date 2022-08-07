// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"github.com/omiga-group/omiga/src/order/order-api/graphql"
	"github.com/omiga-group/omiga/src/order/order-api/http"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewHttpServer(logger *zap.SugaredLogger, appSettings configuration.AppSettings, postgresSettings postgres.PostgresSettings) (http.HttpServer, error) {
	database, err := postgres.NewPostgres(postgresSettings)
	if err != nil {
		return nil, err
	}
	entgoClient, err := repositories.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	orderService, err := services.NewOrderService(entgoClient)
	if err != nil {
		return nil, err
	}
	server, err := graphql.NewGraphQLServer(entgoClient, orderService)
	if err != nil {
		return nil, err
	}
	httpServer, err := http.NewHttpServer(appSettings, server)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}