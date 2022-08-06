package commands

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start binance-processor",
		Long:  "Start binance-processor",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			_ = logger.Sugar()

			time.Sleep(1 * time.Minute)
		},
	}

	return cmd
}
