package shift_time

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

type workConditions struct {
	hoursPerWeek       int8
	numberOfWorkDays   int8
	pauseTimeInMinutes int8
}

const defaultHoursPerWeek int8 = 38
const defaultNumberOfWorkDays int8 = 5
const defaultPauseTimeInMinutes int8 = 45

func NewWorkConditions() workConditions {
	e := workConditions{defaultHoursPerWeek, defaultNumberOfWorkDays, defaultPauseTimeInMinutes}
	return e
}

func SetHoursPerWeek(conditions *workConditions, hours int8) {
	conditions.hoursPerWeek = hours
}

func SetNumberWorkDays(conditions *workConditions, days int8) {
	conditions.numberOfWorkDays = days
}

func SetPauseTimeInMinutes(conditions *workConditions, pauseTime int8) {
	conditions.pauseTimeInMinutes = pauseTime
}

func (conditions workConditions) ShiftEndTime(startTime string) string {
	if !verifyStartTime(startTime) { //TODO: Proper Error handling
		fmt.Println("Error: start time format is HH:MM")
		os.Exit(1)
	}

	shiftEndTime := calculateShiftEnd(conditions, startTime)
	shiftEnd := fmt.Sprintf("%v:%v", shiftEndTime.Hour(), shiftEndTime.Minute())
	return shiftEnd
}

const referenceTimeLayout string = "15:04:05"

func calculateShiftEnd(condition workConditions, startTime string) time.Time {
	timeObj, _ := time.Parse(referenceTimeLayout, startTime)
	workPerDay := (condition.hoursPerWeek * 60) / condition.numberOfWorkDays
	totalTimeAtWork := workPerDay + condition.pauseTimeInMinutes
	shiftEnd := timeObj.Add(time.Minute * time.Duration(totalTimeAtWork))
	return shiftEnd
}

func verifyStartTime(startTime string) bool {
	// startTime := progArgs[0]
	timePattern := "(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]"
	isMatched, _ := regexp.MatchString(timePattern, startTime)

	return isMatched
}
