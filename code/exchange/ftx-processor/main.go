package main

import (
	"fmt"
	"os"

	"github.com/omiga-group/omiga/code/exchange/ftx-processor/business"
	"github.com/omiga-group/omiga/code/exchange/ftx-processor/config"
)

const (
	configPath = "config.yaml"
)

var (
	environment = os.Getenv("OMIGA-ENV")
)

func init() {
	if environment == "" {
		environment = "prod"
	}
}

func main() {
	config.Intialize(configPath, environment)

	business.Foo()

	wait := make(chan bool)
	select {
	case done := <-wait:
		fmt.Println(done)
	}
	// rootCmd := commands.Root()
	// util.PrintIfError(rootCmd.Execute())
}
