package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/omiga-group/omiga/code/shared/messaging/pulsar"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start order-processor",
		Long:  "Start order-processor",
		Run: func(cmd *cobra.Command, args []string) {
			var pulsarSettings pulsar.PulsarSettings
			mapstructure.Decode(viper.Get(pulsar.ConfigKey), &pulsarSettings)

			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			ctx, cancelFunc := context.WithCancel(context.Background())

			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			go func() {
				<-sigc
				cancelFunc()
			}()

			orderConsumer, err := NewOrderConsumer(sugarLogger, pulsarSettings)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = orderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			for {
				if ctx.Err() == context.Canceled {
					break
				}

				select {
				case <-ctx.Done():
				case <-time.After(time.Second):
				}
			}
		},
	}

	return cmd
}
