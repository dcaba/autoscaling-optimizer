package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSimulateCommand_implement(t *testing.T) {
	var _ cli.Command = &SimulateCommand{}
}
