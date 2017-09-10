package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestUpdateCommand_implement(t *testing.T) {
	var _ cli.Command = &UpdateCommand{}
}
