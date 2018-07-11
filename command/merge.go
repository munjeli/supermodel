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

// Synopsis what it means to merge a model
func (c *MergeCommand) Synopsis() string {
	return "Merge two branches on each repository by index"
}

func (c *MergeCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
