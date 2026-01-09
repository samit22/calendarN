package dateconv

import (
	"fmt"
	"time"
)

type Converter struct {
}

func (c *Converter) EtoN(date string) (*Date, error) {
	resp := &Date{}
	givenD, err := time.Parse(IsoDate, date)
	if err != nil {
		return resp, err
	}

	engStartDate := time.Date(engStartYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	if givenD.Before(engStartDate) {
		return resp, fmt.Errorf("date not supported before 1944 Jan")
	}
	durDiff := givenD.Sub(engStartDate)
	dayDiff := int(durDiff.Hours()) / 24
	npDay := nepStartDay + dayDiff
	npMonth := nepStartMonth
	npYear := nepStartYear

	for npDay > nepaliDates[npYear][npMonth-1] {
		npDay -= nepaliDates[npYear][npMonth-1]
		if npMonth == 12 {
			npMonth = 1
			npYear++
			if npYear > nepEndYear {
				return resp, fmt.Errorf("date not supported after year %d BS", nepEndYear)
			}
		} else {
			npMonth++
		}
	}
	return &Date{npYear, npMonth, npDay, givenD.Weekday(), &givenD}, nil
}

func (c *Converter) NtoE(date string) (t *time.Time, err error) {

	d, err := parseToNepaliDate(date)
	if err != nil {
		return
	}
	return d.GetEnglishDate(), nil
}

func findTotalDays(d Date) int {
	var totalDays = d.day - 1

	m := d.month - 2
	y := d.year
	if m < 0 {
		y = y - 1
		m = 11
	}
	for year := y; year >= nepStartYear; year-- {
		for month := m; month >= 0; month-- {
			totalDays += nepaliDates[year][month]
		}
		m = 11
	}
	return totalDays
}
