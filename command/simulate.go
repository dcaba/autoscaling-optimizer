package command

import (
	"strings"
)

type SimulateCommand struct {
	Meta
}

func (c *SimulateCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SimulateCommand) Synopsis() string {
	return ""
}

func (c *SimulateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
