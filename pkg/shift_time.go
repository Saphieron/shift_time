package shift_time

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type workConditions struct {
	hoursPerWeek       int16
	numberOfWorkDays   int16
	pauseTimeInMinutes int16
}

const defaultHoursPerWeek int16 = 38
const defaultNumberOfWorkDays int16 = 5
const defaultPauseTimeInMinutes int16 = 45

func NewWorkConditions() workConditions {
	e := workConditions{defaultHoursPerWeek, defaultNumberOfWorkDays, defaultPauseTimeInMinutes}
	return e
}

func SetHoursPerWeek(conditions *workConditions, hours int16) {
	conditions.hoursPerWeek = hours
}

func SetNumberWorkDays(conditions *workConditions, days int16) {
	conditions.numberOfWorkDays = days
}

func SetPauseTimeInMinutes(conditions *workConditions, pauseTime int16) {
	conditions.pauseTimeInMinutes = pauseTime
}

func (conditions workConditions) ShiftEndTime(startTime string) (string, error) {
	if !verifyStartTime(startTime) { //TODO: Proper Error handling
		return "", errors.New("invalid start time argument")
	}

	shiftEndTime := calculateShiftEnd(conditions, startTime)
	shiftEnd := fmt.Sprintf("%02d:%02d", shiftEndTime.Hour(), shiftEndTime.Minute())
	return shiftEnd, nil
}

const referenceTimeLayout string = "15:04"

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
