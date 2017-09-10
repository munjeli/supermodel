package command

import (
	"strings"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return ""
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
