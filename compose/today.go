package main

import (
	"fmt"
	"time"

	"github.com/samit22/calendarN/dateconv"
)

const (
	total    float64 = 100
	barWidth int     = 30
)

func getToday() string {
	t := time.Now().UTC().Add(time.Hour*5 + time.Minute*45)
	dc := dateconv.Converter{}
	nDate, _ := dc.EtoN(t.Format("2006-01-02"))
	text := fmt.Sprintf("आज:  %s साल, %s महिनाको %s गते\n", nDate.DevanagariYear(), nDate.DevanagariMonth(), nDate.DevanagariDay())
	text += printProgressBar(float64(nDate.YearDay()) / float64(nDate.CurrentYearDays()))
	text = fmt.Sprintf("%s\n (%s/%s दिन)", text, dateconv.EnglishToNepaliNumber(nDate.YearDay()),
		dateconv.EnglishToNepaliNumber(nDate.CurrentYearDays()))
	return text
}

func printProgressBar(progress float64) string {
	blocks := int(progress * float64(barWidth))
	bar := "[" + repeat("▓", blocks) + repeat("░", barWidth-blocks) + "]"
	percentage := int(progress * total)
	return fmt.Sprintf("\r%s %d%%", bar, percentage)
}

func repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return fmt.Sprintf("%s%s", s, repeat(s, count-1))
}
