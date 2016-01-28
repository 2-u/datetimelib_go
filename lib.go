package datetimelib

import (
	"time"
)

func IsSameDay(time1 time.Time, time2 time.Time) bool {

	time1Year, time1Month, time1Day := time1.Date()
	time2Year, time2Month, time2Day := time2.Date()

	if time1Year == time2Year &&
		time1Month == time2Month &&
		time1Day == time2Day {

		return true
	}

	return false
}

func Tomorrow() time.Time {

	return time.Now().Add(24 * time.Hour)
}

func Yesterday() time.Time {

	return time.Now().Add(-(24 * time.Hour))
}
