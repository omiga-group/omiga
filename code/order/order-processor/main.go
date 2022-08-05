package main

import (
	"log"

	"github.com/micro-business/go-core/pkg/util"
	"github.com/omiga-group/omiga/code/order/order-processor/commands"
	"github.com/omiga-group/omiga/code/shared/enterprise/configuration"
)

func main() {
	err := configuration.SetupConfigReader(".")
	if err != nil {
		log.Fatal(err)
	}
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
