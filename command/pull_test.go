package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestPullCommand_implement(t *testing.T) {
	var _ cli.Command = &PullCommand{}
}
