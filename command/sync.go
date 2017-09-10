package command

import (
	"strings"
)

type SyncCommand struct {
	Meta
}

func (c *SyncCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SyncCommand) Synopsis() string {
	return ""
}

func (c *SyncCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
