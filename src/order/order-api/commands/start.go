package commands

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start order-api",
		Long:  "Start order-api",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var appSettings configuration.AppSettings
			if err := mapstructure.Decode(viper.Get(configuration.AppSettingsConfigKey), &appSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			var postgresSettings postgres.PostgresSettings
			if err := mapstructure.Decode(viper.Get(postgres.ConfigKey), &postgresSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			httpServer, err := NewHttpServer(sugarLogger, appSettings, postgresSettings)
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
