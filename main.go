package main

import (
	"fmt"
	"os"
	shift_time "saphieron/shift_time/shift_time_lib"
)

func main() {
	fmt.Println("running shift_calc")

	startTime := os.Args[1]
	// if len(os.Args) > 2 {
	//TODO: Create some parsing logic
	// }
	run(startTime)
}

func run(startTime string) {
	conditions := shift_time.NewWorkConditions()
	fmt.Println(startTime)
	fmt.Printf("Shift ends at: %v\n", conditions.ShiftEndTime(startTime))
}
