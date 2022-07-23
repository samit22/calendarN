package data

import "time"

var (
	EngDayLong  = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	EngDayShort = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
)

func GetDayLong(t time.Time) string {
	return EngDayLong[t.Weekday()]
}
func GetDayShort(t time.Time) string {
	return EngDayShort[t.Weekday()]
}
