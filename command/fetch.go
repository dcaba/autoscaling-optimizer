package command

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
)

func CmdFetch(commandCliContext *cli.Context) {
	from := commandCliContext.String("from")
	to := commandCliContext.String("to")
	fmt.Fprintln(stdOutput, "Collecting data from", from, "to", to)
	log.Println("Verbose sample output")
	defer fmt.Fprintln(stdOutput, "Command completed successfully")
}
