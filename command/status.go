package command

import (
	"strings"
)

type StatusCommand struct {
	Meta
}

func (c *StatusCommand) Run(args []string) int {
	// Write your code here

	return 0
}

// Synopsis of the model status command
func (c *StatusCommand) Synopsis() string {
	return "Return status of all checked out branches"
}

func (c *StatusCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
