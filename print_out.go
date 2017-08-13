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

	// help_from: https://youtu.be/B-r3Wf_I2Lk?t=144
	for _, fault := range faults.Results {

		last_notice_at := fault.LastNoticeAt
		message := fault.Message
		notices_count := fault.NoticesCount

		fmt.Printf("%s :>> %s %v\n", green(last_notice_at), red(message),  yellow(notices_count))
	}
	
}
