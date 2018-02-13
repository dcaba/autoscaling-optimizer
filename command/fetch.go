package command

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
)

func CmdFetch(commandCliContext *cli.Context) {
	from := commandCliContext.Int("from")
	to := commandCliContext.Int("to")
	log.Println("Data request from", from, "days ago until", to, "days ago")

	outputFile := commandCliContext.String("output")
	outFileDescriptor, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintln(stdOutput, "Unable to create output file", outputFile)
		currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		log.Fatalln("Unable to create output file", outputFile, "when running in dir", currentPath,
			". Complete error:", err)
	}
	log.Println("Output file", outputFile, "was successufully initialized", outFileDescriptor)
	//
	//dateFrom := calculateDateFromNow(from)
	//dateTo := calculateDateFromNow(to)
	//fmt.Fprintln(stdOutput, "Getting metrics from", dateFrom, "until", dateTo)
	//collectedAwsMetrics := awsMetricsCollector.FetchAwsMetrics(dateFrom, dateTo)
	//collectedAwsMetrics.SaveTo(outFileDescriptor)

	defer fmt.Fprintln(stdOutput, "Command completed successfully")
}
