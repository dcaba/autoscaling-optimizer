package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestCombineCommand_implement(t *testing.T) {
	var _ cli.Command = &CombineCommand{}
}
