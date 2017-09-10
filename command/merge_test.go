package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestMergeCommand_implement(t *testing.T) {
	var _ cli.Command = &MergeCommand{}
}
