package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("running shift_calc")
	argsWithoutProgrammeName := os.Args[1:]
	startTime := getStartTimeFromArgs(argsWithoutProgrammeName)

	fmt.Printf("startTime %v\n", startTime)

	// referenceTimeLayout := "15:04:05"
	// timeObj, _ := time.Parse(referenceTimeLayout, "08:10:00")
	// startTime := os.Args[1]

	// timePattern := "hh:mm"

	// stringResult := fmt.Sprintf("StartTime: %v", startTime)
	// fmt.Printf("input time %v\n", timeObj)
	// fmt.Println(stringResult)
}

func getStartTimeFromArgs(progArgs []string) string {
	startTime := progArgs[0]
	timePattern := "(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]"
	isMatched, _ := regexp.MatchString(timePattern, startTime)
	if !isMatched {
		fmt.Println("Error: start time format is HH:MM")
		os.Exit(1)
	}
	return startTime
}
