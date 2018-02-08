package main

import (
	"os"

	"github.com/urfave/cli"
	"io/ioutil"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Author = "dcaba"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Before = func(c *cli.Context) error {
		debug := c.Bool("debug")
		if !debug {
			log.SetOutput(ioutil.Discard)
		}
		log.Println("Program starting")
		return nil
	}
	app.After = func(c *cli.Context) error {
		defer log.Println("Program finished")
		return nil
	}
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
