package main

import (
	"fmt"
	"os"

	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/cli"
)

func main() {
	config, showHelp, err := cli.Parse()

	if showHelp {
		cli.PrintUsage()
		return
	}

	if err != nil {
		cli.PrintErrorAndUsage(err)
		os.Exit(1)
	}

	fmt.Println("Configuration accepted:")
	fmt.Printf("%+v\n", config)
}
