package main

import (
	"fmt"
	"os"
	shift_time "saphieron/shift_time/pkg"
)

func main() {
	fmt.Println("running shift_calc")

	startTime := os.Args[1] //TODO:
	run(startTime)
}

func run(startTime string) {
	conditions := shift_time.NewWorkConditions()
	fmt.Println(startTime)
	result, err := conditions.ShiftEndTime(startTime)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Shift ends at: %v\n", result)
	}
}
