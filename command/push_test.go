package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestPushCommand_implement(t *testing.T) {
	var _ cli.Command = &PushCommand{}
}
