package datetimelib

import (
	"testing"
	"time"
)

func TestIsSameDay(t *testing.T) {

	time1 := time.Now()
	time2 := time.Now()

	isSameDay := IsSameDay(time1, time2)
	t.Log("IsSameDay( time.Now, time.Now ) = ", isSameDay)

	tomorrowDate := time.Now().Add(24 * time.Hour)
	isSameDay = IsSameDay(time1, tomorrowDate)
	t.Log("IsSameDay( time.Now, tomorrowDate ) = ", isSameDay)

	t.Errorf("Error: Testing CalcIsSameDay")
}
