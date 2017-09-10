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

func (c *PullCommand) Synopsis() string {
	return ""
}

func (c *PullCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
