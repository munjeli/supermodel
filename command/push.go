package command

import (
	"strings"
)

type PushCommand struct {
	Meta
}

func (c *PushCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *PushCommand) Synopsis() string {
	return ""
}

func (c *PushCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
