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

// Synopsis what it means to push all the repos
func (c *PushCommand) Synopsis() string {
	return "Push commits from all repositories to their respective remotes"
}

func (c *PushCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
