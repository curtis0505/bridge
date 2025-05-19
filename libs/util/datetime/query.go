package datetime

import (
	"fmt"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type TimeRange string

func ParseTimeRange(s string) (TimeRange, error) {
	if !lo.Contains([]TimeRange{TimeRangeWeek, TimeRangeDay, TimeRangeMonth, TimeRangeYear}, TimeRange(s)) {
		return "", fmt.Errorf("invalid time range: %s", s)
	}
	return TimeRange(s), nil
}

const (
	TimeRangeDay   TimeRange = "DAY"
	TimeRangeWeek  TimeRange = "WEEK"
	TimeRangeMonth TimeRange = "MONTH"
	TimeRangeYear  TimeRange = "YEAR"
)

func (r TimeRange) String() string { return string(r) }

// BeforeTime returns the time before the current time based on the time range.
// default is 1 day
func (r TimeRange) BeforeTime(currentTime time.Time) time.Time {
	switch r {
	case TimeRangeWeek:
		return currentTime.AddDate(0, 0, -7)
	case TimeRangeMonth:
		return currentTime.AddDate(0, -1, 0)
	case TimeRangeYear:
		return currentTime.AddDate(-1, 0, 0)
	default:
		return currentTime.AddDate(0, 0, -1)
	}
}

func (r TimeRange) MongoQuery(currentTime time.Time) (matchTime bson.A, beforeTime time.Time) {
	switch r {
	case TimeRangeDay:
		matchTime = bson.A{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
		beforeTime = currentTime.AddDate(0, 0, -1)
	case TimeRangeWeek:
		matchTime = bson.A{19, 1, 7, 13}
		beforeTime = currentTime.AddDate(0, 0, -7)
	case TimeRangeMonth:
		matchTime = bson.A{0}
		beforeTime = currentTime.AddDate(0, -1, 0)
	case TimeRangeYear:
		matchTime = bson.A{0}
		beforeTime = currentTime.AddDate(-1, 0, 0)
	}

	return
}
