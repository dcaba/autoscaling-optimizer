package command

import (
	"github.com/urfave/cli"
	"log"
)

func CmdFetch(c *cli.Context) {
	from := c.String("from")
	to := c.String("to")
	log.Println("Collecting data from", from, "to", to)
}
