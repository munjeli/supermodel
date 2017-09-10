package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSyncCommand_implement(t *testing.T) {
	var _ cli.Command = &SyncCommand{}
}
