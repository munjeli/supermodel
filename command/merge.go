package command

import (
	"strings"
)

type MergeCommand struct {
	Meta
}

func (c *MergeCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *MergeCommand) Synopsis() string {
	return ""
}

func (c *MergeCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
