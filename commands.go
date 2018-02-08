package main

import (
	"fmt"
	"os"

	"github.com/dcaba/autoscaling-optimizer/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "debug, d",
		Usage: "Enable debugging output",
	},
}

var Commands = []cli.Command{
	{
		Name: "fetch",
		UseShortOptionHandling: true,
		Usage:  "Fetch performance metrics from AWS Cloudfront you can use as the scenario for the simulations",
		Action: command.CmdFetch,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "from, f",
				Value: "31",
				Usage: "Days from now from which you want to start collecting data",
			},
			cli.StringFlag{
				Name:  "to, t",
				Value: "1",
				Usage: "Days from now you want to finish collecting data",
			},
		},
	},
	{
		Name:   "simulate",
		UseShortOptionHandling: true,
		Usage:  "",
		Action: command.CmdSimulate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "combine",
		UseShortOptionHandling: true,
		Usage:  "",
		Action: command.CmdCombine,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
