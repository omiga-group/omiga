package main

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/omiga-group/omiga/src/shared/omigactl/commands"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
