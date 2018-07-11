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

// Synopsis of how to delete the contents of a model
func (c *DeleteCommand) Synopsis() string {
	return "Delete all of the repositories and namespaces leaving just  the root and .smdl"
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
