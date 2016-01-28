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

func TestTomorrow(t *testing.T) {

	tomorrow := Tomorrow()
	t.Log("tomorrow =", tomorrow)

	t.Errorf("Error: Testing Tomorow")
}

func TestYesterday(t *testing.T) {

	yesterday := Yesterday()
	t.Log("yesterday =", yesterday)

	t.Errorf("Error: Testing Yesterday")
}
