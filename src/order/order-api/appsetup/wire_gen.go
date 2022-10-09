// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/omiga-group/omiga/src/order/order-api/graphql"
	"github.com/omiga-group/omiga/src/order/order-api/http"
	"github.com/omiga-group/omiga/src/order/order-api/publishers"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/entities"
	outbox2 "github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewCronService(logger *zap.SugaredLogger) (cron.CronService, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	cronService, err := cron.NewCronService(logger, timeHelper)
	if err != nil {
		return nil, err
	}
	return cronService, nil
}

func NewEntgoClient(logger *zap.SugaredLogger, postgresConfig postgres.PostgresConfig) (entities.EntgoClient, error) {
	database, err := postgres.NewPostgres(logger, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := entities.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	return entgoClient, nil
}

func NewOutboxBackgroundService(ctx context.Context, logger *zap.SugaredLogger, pulsarConfig pulsar.PulsarConfig, outboxConfig outbox.OutboxConfig, entgoClient entities.EntgoClient, cronService cron.CronService) (outbox2.OutboxBackgroundService, error) {
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	pulsarClient, err := pulsar.NewPulsarClient(logger, pulsarConfig, osHelper)
	if err != nil {
		return nil, err
	}
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarClient)
	if err != nil {
		return nil, err
	}
	outboxBackgroundService, err := outbox2.NewOutboxBackgroundService(ctx, logger, outboxConfig, messageProducer, entgoClient, cronService)
	if err != nil {
		return nil, err
	}
	return outboxBackgroundService, nil
}

func NewHttpServer(logger *zap.SugaredLogger, appConfig configuration.AppConfig, entgoClient entities.EntgoClient, orderOutboxBackgroundService outbox2.OutboxBackgroundService) (http.HttpServer, error) {
	orderRepository, err := repositories.NewOrderRepository()
	if err != nil {
		return nil, err
	}
	outboxPublisher, err := outbox2.NewOutboxPublisher(logger, entgoClient)
	if err != nil {
		return nil, err
	}
	orderPublisher, err := publishers.NewOrderPublisher(logger, appConfig, outboxPublisher)
	if err != nil {
		return nil, err
	}
	orderService, err := services.NewOrderService(logger, entgoClient, orderRepository, orderPublisher)
	if err != nil {
		return nil, err
	}
	server, err := graphql.NewGraphQLServer(entgoClient, orderService, orderOutboxBackgroundService)
	if err != nil {
		return nil, err
	}
	httpServer, err := http.NewHttpServer(logger, appConfig, server)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}
