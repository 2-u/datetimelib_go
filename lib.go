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

func GetCurrentWeeklyOptionsExpirationDate() time.Time {

	// This gets the current week's weekly options expiration date
	// using today's date as the reference date.
	// It is usually on Friday, unless the current Friday is a holiday

	dayOfWeek := time.Now().Weekday().String()

	oneDay := 24 * time.Hour

	currentExpirationFridayMap := map[string]time.Time{
		"Monday":    time.Now().Add(4 * oneDay),
		"Tuesday":   time.Now().Add(3 * oneDay),
		"Wednesday": time.Now().Add(2 * oneDay),
		"Thursday":  time.Now().Add(oneDay),
		"Friday":    time.Now(),
	}

	return currentExpirationFridayMap[dayOfWeek]
}

func GetNextWeeklyOptionsExpirationDate() time.Time {

	// This gets the next week's weekly options expiration date
	// using today's date as the reference date.
	// We do this by getting this week's options expiration date and
	// adding 7 days to it.

	currentWeekOptionsExpirationDate := GetCurrentWeeklyOptionsExpirationDate()
	return currentWeekOptionsExpirationDate.Add(7 * 24 * time.Hour)
}
