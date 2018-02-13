package command

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"testing"
)

func TestCmdFetch(t *testing.T) {
	fakeFlags := flag.NewFlagSet("dummy", flag.ErrorHandling(0))
	fakeFlags.Int("from", 2, "")
	fakeFlags.Int("to", 1, "")
	testFile := "data_test/test_output_file.json"
	fakeFlags.String("output", testFile, "")
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

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		t.Error("Command CmdFetch did not generate the expected output file", testFile, "in exec dir", currentPath,
			". Actual error:", err)
	}
}
