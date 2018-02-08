package command

import (
	"github.com/urfave/cli"
	"log"
	"fmt"
)

func CmdFetch(c *cli.Context) {
	from := c.String("from")
	to := c.String("to")
	fmt.Println("Collecting data from", from, "to", to)
	defer log.Println("Data collected successfully")
}
