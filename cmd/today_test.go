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

		if d.Year() != tn.Year()+57 {
			t.Errorf("should have received accurate date")
		}
	}
}

func Test_getToday(t *testing.T) {
	t.Log("run test for getToday() should not return panic")
	{
		getToday()
	}
}
