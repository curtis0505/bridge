package datetime

import (
	"sync"
	"time"
)

var loc *time.Location
var mu = new(sync.Mutex)

func init() {
	SetLocation("UTC")
}

// SetLocation SHOULD BE CALLED BEFORE USING ANY OTHER FUNCTIONS IN THIS PACKAGE. sets the location for the time package.
func SetLocation(location string) {
	mu.Lock()
	defer mu.Unlock()
	loc, _ = time.LoadLocation(location)
}

// Now returns the current time in the location set by SetLocation.
func Now() time.Time {
	return time.Now().In(loc)
}

func Yesterday() time.Time {
	return Now().AddDate(0, 0, -1)
}

// NowString returns the current time in the location set by SetLocation in the given layout.
func NowString(layout string) string {
	return Now().Format(layout)
}

// AddHour adds the given number of hours to the given time.
func AddHour(t time.Time, hour int) time.Time {
	return t.In(loc).Add(time.Hour * time.Duration(hour))
}

// AddDay adds the given number of days to the given time.
func AddDay(t time.Time, day int) time.Time {
	return t.In(loc).AddDate(0, 0, day)
}

// AddMonth adds the given number of months to the given time.
func AddMonth(t time.Time, month int) time.Time {
	return t.In(loc).AddDate(0, month, 0)
}

// AddYear adds the given number of years to the given time.
func AddYear(t time.Time, year int) time.Time {
	return t.In(loc).AddDate(year, 0, 0)
}

// After returns true if the given time is after the given duration.
func After(t time.Time, d time.Time) bool {
	return t.In(loc).After(d.In(loc))
}

// Before returns true if the given time is before the given duration.
func Before(t time.Time, d time.Time) bool {
	return t.In(loc).Before(d.In(loc))
}

// Between returns true if the given time is between the given start and end times.
func Between(t, start, end time.Time) bool {
	return After(t.In(loc), start.In(loc)) && Before(t.In(loc), end.In(loc))
}

// Equal returns true if the given times are equal. The truncate parameter can be used to truncate the time to a given duration.
func Equal(t, d time.Time, truncate time.Duration) bool {
	return t.In(loc).Truncate(truncate).Equal(d.In(loc).Truncate(truncate))
}

// StringToTime converts a string to a time.Time object. The string must be in the format "2006-01-02T15:04" or "2006-01-02T15:04:05Z07:00".
func StringToTime(value string) (time.Time, error) {
	t, err := time.ParseInLocation("2006-01-02T15:04", value, loc)
	if err != nil {
		t2, err2 := time.ParseInLocation(time.RFC3339, value, loc)
		if err2 != nil {
			return time.Time{}, err2
		}
		return t2, nil
	}
	return t, nil
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
