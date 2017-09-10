package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestCheckoutCommand_implement(t *testing.T) {
	var _ cli.Command = &CheckoutCommand{}
}
