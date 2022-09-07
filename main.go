package main

import (
	"github.com/TeemoKill/GGBot/actions"
	"log"
	"os"

	"github.com/TeemoKill/GGBot/consts"

	"github.com/urfave/cli"
)

func main() {
	botCli := cli.NewApp()
	botCli.Name = consts.GGAppName

	botCli.Flags = []cli.Flag{
		cli.StringFlag{
			Name: consts.CliFlagConfig,
			Usage: "specify config file path, overrides default config path if provided",
			Required: false,
		},
	}

	botCli.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "start running ggbot",
			Action: actions.Run,
		},
	}

	// run cli
	err := botCli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

