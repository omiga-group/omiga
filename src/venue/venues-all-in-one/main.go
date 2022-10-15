package main

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/omiga-group/omiga/src/venue/venues-all-in-one/commands"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
