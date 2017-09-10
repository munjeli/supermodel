package command

import (
	"strings"
)

type CheckoutCommand struct {
	Meta
}

func (c *CheckoutCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *CheckoutCommand) Synopsis() string {
	return ""
}

func (c *CheckoutCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
