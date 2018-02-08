package command

import (
	"io"
	"os"
)

var successfulExecutionOutput = "Command completed successfully"
var stdOutput = io.Writer(os.Stdout)
