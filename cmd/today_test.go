package cmd

import (
	"testing"
	"time"
)

func Test_getNepToday(t *testing.T) {

	t.Log("run test for getNepToday()")
	{
		d := getNepToday()
		tn := time.Now()

		// Nepali year is 56-57 years ahead of English year
		// Before Nepali new year (mid-April) it's +56, after it's +57
		diff := d.Year() - tn.Year()
		if diff != 56 && diff != 57 {
			t.Errorf("should have received accurate date, got diff: %d", diff)
		}
	}
}

func Test_getToday(t *testing.T) {
	t.Log("run test for getToday() should not return panic")
	{
		today := getToday()
		if today.Year != time.Now().Year() {
			t.Errorf("invalid today date")
		}
	}
}

func Test_properResponse(t *testing.T) {
	t.Log("run test for json output")
	{
		e := getToday()
		n := getNepToday()
		resp := properResponse(n, e)
		if resp.English.Year != time.Now().Year() {
			t.Errorf("invalid today date")
		}
	}
}
