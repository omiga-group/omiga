package main

import (
	"log"

	"github.com/micro-business/go-core/pkg/util"
	"github.com/omiga-group/omiga/src/order/shared/repositories/cli/commands"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
)

func main() {
	err := configuration.SetupConfigReader(".")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
