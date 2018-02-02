package command

import (
	"strings"
)

type FetchCommand struct {
	Meta
}

func (c *FetchCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *FetchCommand) Synopsis() string {
	return ""
}

func (c *FetchCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
