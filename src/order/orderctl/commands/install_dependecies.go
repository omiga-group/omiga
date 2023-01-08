package commands

import (
	"github.com/spf13/cobra"
)

func installDependeciesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-dependecies",
		Short: "install dependecies",
		Long:  "install dependecies",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
