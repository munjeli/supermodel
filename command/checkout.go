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

// Synopsis of the checkout command as with index
func (c *CheckoutCommand) Synopsis() string {
	return "Checkout branches on the model by index"
}

func (c *CheckoutCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
