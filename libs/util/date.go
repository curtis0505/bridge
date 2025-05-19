package util

import (
	"errors"
	"time"
)

func BeforeNDayToYesterday24H(n int, location *time.Location) (time.Time, time.Time, error) {

	if n <= 0 || location == nil {
		return time.Time{}, time.Time{}, errors.New("invalid n or location is nil")
	}

	calcStartDate := time.Now().In(location).AddDate(0, 0, n*-1)
	calcEndDate := time.Now().In(location).AddDate(0, 0, -1)

	startDate := time.Date(calcStartDate.Year(), calcStartDate.Month(), calcStartDate.Day(), 0, 0, 0, 0, location)
	endDate := time.Date(calcEndDate.Year(), calcEndDate.Month(), calcEndDate.Day(), 23, 59, 59, 0, location)

	return startDate, endDate, nil
}

func NewDailyDate(location *time.Location, truncate *time.Duration) time.Time {
	t := time.Now().In(location)
	if truncate != nil {
		return t.Truncate(*truncate)
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
	return t
}
