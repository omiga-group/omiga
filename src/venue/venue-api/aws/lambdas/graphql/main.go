package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	enterpriseappsetup "github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	venueappsetup "github.com/omiga-group/omiga/src/venue/shared/appsetup"
	"github.com/omiga-group/omiga/src/venue/venue-api/appsetup"
	"github.com/omiga-group/omiga/src/venue/venue-api/configuration"
)

func main() {
	sugarLogger := logger.CreateLogger()

	configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	var config configuration.Config
	if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
		sugarLogger.Fatal(err)
	}

	entgoClient, err := venueappsetup.NewEntgoClient(
		sugarLogger,
		config.Postgres)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App,
		entgoClient)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	lambda.Start(httpadapter.New(httpServer.GetGraphQLHandler()).ProxyWithContext)
}
