package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/omiga-group/omiga/src/shared/omigactl/configuration"
	"github.com/spf13/cobra"
)

type createDatabaseOptions struct {
	name string
}

func CreateCommand(databaseCommand *cobra.Command) *cobra.Command {
	options := createDatabaseOptions{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Database",
		Long:  "Create Database",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			sugarLogger := logger.CreateLogger()

			configurationHelper, err := appsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			if connectionString, err := databaseCommand.PersistentFlags().GetString("connectionString"); err != nil {
				sugarLogger.Fatal(err)
			} else {
				connectionString = strings.TrimSpace(connectionString)

				if len(connectionString) != 0 {
					config.Postgres.ConnectionString = connectionString
				}
			}

			database, err := appsetup.NewDatabase(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer database.Close()

			var found int

			db := database.GetDB()
			if err := db.QueryRowContext(
				ctx,
				fmt.Sprintf(
					"SELECT 1 FROM pg_database WHERE datname = '%s'",
					options.name)).
				Scan(&found); err == sql.ErrNoRows {
				if _, err = db.ExecContext(
					ctx,
					fmt.Sprintf("CREATE DATABASE \"%s\"", options.name)); err != nil {
					sugarLogger.Fatal(err)
				}

				sugarLogger.Infof("Database %s successfully created.", options.name)
			} else if err != nil {
				sugarLogger.Fatal(err)
			} else {
				sugarLogger.Infof("Database %s already exists. Ignore creating the database.", options.name)
			}
		},
	}

	cmd.Flags().StringVar(&options.name, "name", "", "Specify the database name to create")

	cmd.MarkFlagRequired("name")

	return cmd
}
