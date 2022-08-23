package time

import "time"

var (
	froze      = false
	frozenDate time.Time
)

// Check function to check if a time is valid or not.
func Check(date *time.Time) (time.Time, bool) {
	var defaultTime = time.Time{}
	if date != nil && *date != defaultTime {
		return *date, true
	}
	return time.Time{}, false
}

// Unfreeze function to unfreeze date to specific date for mocking purpose.
func Unfreeze() {
	froze = false
	frozenDate = time.Time{}
}

// Freeze function to freeze date to specific date for mocking purpose.
func Freeze(date time.Time) {
	froze = true
	frozenDate = date
}

// Now function to get current time.
func Now() time.Time {
	if froze {
		return frozenDate
	}
	return time.Now()
}
