package command

import (
	"strings"
)

type PullCommand struct {
	Meta
}

func (c *PullCommand) Run(args []string) int {
	// Write your code here

	return 0
}

// Synopsis of how to pull a model
func (c *PullCommand) Synopsis() string {
	return "Pulls latest of just the checked out branches"
}

func (c *PullCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
