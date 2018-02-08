package command

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/urfave/cli"
	"os"
	"testing"
)

func TestCmdFetch(t *testing.T) {
	fakeFlags := flag.NewFlagSet("dummy", flag.ErrorHandling(0))
	fakeFlags.String("from", "2", "")
	fakeFlags.String("to", "1", "")
	fakeCmdFetchCtx := cli.NewContext(nil, fakeFlags, nil)

	var b bytes.Buffer
	stdOutput = &b
	defer func() { stdOutput = os.Stdout }()

	CmdFetch(fakeCmdFetchCtx)

	scanner := bufio.NewScanner(&b)
	var lastCollectedOutput string
	for scanner.Scan() {
		lastCollectedOutput = scanner.Text()
	}
	if lastCollectedOutput != successfulExecutionOutput {
		t.Error("Command CmdFetch did not generate a succesful output:", lastCollectedOutput)
	}
}
