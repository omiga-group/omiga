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
		Short: "Return mexc-processor version",
		Long:  "Return mexc-processor version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Mexc Processor\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, Omiga Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
