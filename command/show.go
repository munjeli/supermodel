package command

import (
	"strings"
)

type ShowCommand struct {
	Meta
}

func (c *ShowCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ShowCommand) Synopsis() string {
	return ""
}

func (c *ShowCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
