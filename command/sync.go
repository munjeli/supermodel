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

// Synopsis of what it means to sync for a model
func (c *SyncCommand) Synopsis() string {
	return "Pull latest on all branches in all repositories"
}

func (c *SyncCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
