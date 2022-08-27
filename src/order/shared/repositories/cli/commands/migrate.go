package commands

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/omiga-group/omiga/src/order/shared/repositories/cli/configuration"
	"github.com/omiga-group/omiga/src/order/shared/repositories/migrate"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type migrateDBOptions struct {
	postMigrationScriptPath string
}

func migrateCommand() *cobra.Command {
	opt := migrateDBOptions{}
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database to the latest version",
		Long:  "Migrate database to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			index := strings.LastIndex(config.Postgres.ConnectionString, "/")
			connectionStringWithoutDatabase := config.Postgres.ConnectionString[:index]
			databaseName := config.Postgres.ConnectionString[index+1:]

			database, err := NewDatabase(
				sugarLogger,
				postgres.PostgresConfig{
					ConnectionString: connectionStringWithoutDatabase,
					MaxOpenConns:     config.Postgres.MaxOpenConns,
				})
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
					databaseName)).
				Scan(&found); err == sql.ErrNoRows {
				if _, err = db.ExecContext(
					ctx,
					fmt.Sprintf("CREATE DATABASE \"%s\"", databaseName)); err != nil {
					sugarLogger.Fatal(err)
				}
			} else if err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := NewEntgoClient(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer entgoClient.Close()

			osHelper, err := NewOsHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			client := entgoClient.GetClient()

			if err = client.Schema.WriteTo(
				ctx,
				os.Stdout,
				migrate.WithGlobalUniqueID(true)); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = client.Schema.Create(
				ctx,
				migrate.WithGlobalUniqueID(true)); err != nil {
				sugarLogger.Fatal(err)
			}

			if len(opt.postMigrationScriptPath) > 0 {
				if !osHelper.FileExist(opt.postMigrationScriptPath) {
					err := fmt.Errorf("post migration script file does not exist. Path: %s", opt.postMigrationScriptPath)

					sugarLogger.Fatal(err)
				}

				postMigrationScript, err := osHelper.GetFileAsString(opt.postMigrationScriptPath)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				sugarLogger.Info("Applying post migration script...")

				database, err = NewDatabase(
					sugarLogger,
					config.Postgres)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				_, err = database.GetDB().ExecContext(ctx, postMigrationScript)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				sugarLogger.Info("Applying post migration script done.")
			}

			sugarLogger.Info("Successfully migrated database to the latest schema")
		},
	}

	cmd.Flags().StringVar(&opt.postMigrationScriptPath, "post-migration-script-path", "", "Specify the path to the SQL script to run after entgo migration is done")

	return cmd
}
