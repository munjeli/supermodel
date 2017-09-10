package command

import (
	"strings"
)

type UpdateCommand struct {
	Meta
}

func (c *UpdateCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *UpdateCommand) Synopsis() string {
	return ""
}

func (c *UpdateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
