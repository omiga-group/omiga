package commands

import (
	"log"

	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func installDependeciesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-dependecies",
		Short: "install dependecies",
		Long:  "install dependecies",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			if err := playwright.Install(); err != nil {
				sugarLogger.Fatalf("Failed to install Playwright. Error: %v", err)
			}
		},
	}

	return cmd
}
