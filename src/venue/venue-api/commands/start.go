package commands

import (
	"log"

	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/venue/venue-api/appsetup"
	"github.com/omiga-group/omiga/src/venue/venue-api/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start venue-api",
		Long:  "Start venue-api",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := appsetup.NewEntgoClient(
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
		},
	}

	return cmd
}
