package main

import (
	"fmt"
	"os"

	"github.com/dcaba/autoscaling-optimizer/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "fetch",
		Usage:  "",
		Action: command.CmdFetch,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "simulate",
		Usage:  "",
		Action: command.CmdSimulate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "combine",
		Usage:  "",
		Action: command.CmdCombine,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
