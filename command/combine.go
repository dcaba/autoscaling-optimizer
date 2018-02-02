package command

import (
	"strings"
)

type CombineCommand struct {
	Meta
}

func (c *CombineCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *CombineCommand) Synopsis() string {
	return ""
}

func (c *CombineCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
