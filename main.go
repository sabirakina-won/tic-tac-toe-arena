package main

import (
	"os"

	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/cli"
	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/game"
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

	game.Start(config)
}
