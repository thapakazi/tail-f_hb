package main

import (
	"github.com/fatih/color"
	"fmt"
)

func PrintFaults(faults *Faults){
	// Create SprintXxx functions to mix strings with other non-colorized strings:
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for fault_index := range faults.Results {

		last_notice_at := faults.Results[fault_index].LastNoticeAt
		message := faults.Results[fault_index].Message
		notices_count := faults.Results[fault_index].NoticesCount

		fmt.Printf("%s :>> %s %v\n", green(last_notice_at), red(message),  yellow(notices_count))
	}
	
}
