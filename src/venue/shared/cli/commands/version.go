package commands

import (
	"fmt"
	"time"

	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Return venuectl version",
		Long:  "Return venuectl version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Venue Cli\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, Omiga Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
