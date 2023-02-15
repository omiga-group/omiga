package main

import (
	"os"

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

    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        config.App.ListeningInterface = ":" + val
    }else{
		config.App.ListeningInterface = ":8080"
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

	err = httpServer.ListenAndServe()
	if err != nil {
		sugarLogger.Fatal(err)
	}
}
