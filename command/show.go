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

// Synopsis what it means to show a model on the command line
func (c *ShowCommand) Synopsis() string {
	return "Show the model on the command line with current branch hinting"
}

func (c *ShowCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
