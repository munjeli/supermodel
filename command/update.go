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

// Synopsis of how to update a model
func (c *UpdateCommand) Synopsis() string {
	return "Reread the model configuration and rebuild the working model"
}

func (c *UpdateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
