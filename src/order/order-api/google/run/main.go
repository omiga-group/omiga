package main

import (
	"context"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/order/order-api/appsetup"
	"github.com/omiga-group/omiga/src/order/order-api/configuration"
	orderappsetup "github.com/omiga-group/omiga/src/order/shared/appsetup"
	enterpriseappsetup "github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
)

func main() {
	ctx := context.Background()
	sugarLogger := logger.CreateLogger()

	configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	var config configuration.Config
	if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
		sugarLogger.Fatal(err)
	}

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		config.App.ListeningInterface = ":8080"
	} else {
		config.App.ListeningInterface = ":" + port
	}

	entgoClient, err := orderappsetup.NewEntgoClient(
		sugarLogger,
		config.Postgres)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	jobScheduler := gocron.NewScheduler(time.UTC)
	jobScheduler.StartAsync()
	defer jobScheduler.Stop()

	pulsarClient, err := enterpriseappsetup.NewPulsarClient(
		sugarLogger,
		config.Pulsar)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	defer pulsarClient.Close()

	outboxBackgroundService, err := orderappsetup.NewOutboxBackgroundService(
		ctx,
		sugarLogger,
		pulsarClient,
		config.Pulsar,
		config.Outbox,
		entgoClient,
		jobScheduler)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	defer outboxBackgroundService.Close()

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App,
		entgoClient,
		outboxBackgroundService)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		sugarLogger.Fatal(err)
	}
}
