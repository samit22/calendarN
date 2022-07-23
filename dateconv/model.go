package dateconv

import (
	"fmt"
	"time"
)

type Date struct {
	year    int
	month   int
	day     int
	weekDay time.Weekday
	engTime *time.Time
}

func NewDate(y, m, d int) (*Date, error) {
	npDate := &Date{
		year:  y,
		month: m,
		day:   d,
	}
	err := npDate.Validate()
	if err != nil {
		return nil, err
	}
	npDate.GetEnglishDate()
	npDate.weekDay = npDate.engTime.Weekday()
	return npDate, nil
}

func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
func (d *Date) WeekDay() time.Weekday {
	return d.weekDay
}

func (d *Date) RomanFullDate() string {
	return fmt.Sprintf("%d-%02d-%02d", d.year, d.month, d.day)
}
func (d *Date) RomanMonth() string {
	return NepaliMonth(d.month)[0]
}

func (d *Date) RomanWeekDay() string {
	return NepaliWeekDay(int(d.weekDay))[0]
}

func (d *Date) DevanagariFullDate() string {
	return fmt.Sprintf("%s-%s-%s", d.DevanagariYear(), englishNumberToNepali(d.month), d.DevanagariDay())
}

func (d *Date) DevanagariYear() string {
	return englishNumberToNepali(d.year)
}

func (d *Date) DevanagariMonth() string {
	return NepaliMonth(int(d.month))[1]
}
func (d *Date) DevanagariDay() string {
	return englishNumberToNepali(d.day)
}
func (d *Date) DevanagariWeekDay() string {
	return NepaliWeekDay(int(d.weekDay))[1]
}
func (d *Date) YearDay() int {
	var days = d.day
	m := d.month - 2
	if m < 0 {
		return days
	}
	yrs := nepaliDates[d.year]
	for month := m; month >= 0; month-- {
		days += yrs[month]
	}
	return days

}

func (d *Date) GetEnglishDate() (t *time.Time) {
	if d.engTime != nil {
		return d.engTime
	}
	totalDays := findTotalDays(*d)

	engStartDate := time.Date(engDateB1, engStartMonth, engStartDay, 0, 0, 0, 0, time.Local)
	engDate := engStartDate.Add(time.Duration(totalDays) * 24 * time.Hour)
	d.engTime = &engDate
	return &engDate
}

func (d *Date) Validate() error {

	if d.year < nepStartYear || d.year > nepEndYear {
		return fmt.Errorf("invalid year %d, range supports %d to %d", d.year, nepStartYear, nepEndYear)
	}
	if d.month < 1 || d.month > 12 {
		return fmt.Errorf("invalid month %d, should be between %d to %d", d.month, 1, 12)
	}
	if d.day < 1 || d.day > 32 {
		return fmt.Errorf("invalid day %d should be between %d to %d", d.day, 1, 12)

	}
	// d.month - 1 since index starts with 0
	if d.day > nepaliDates[d.year][d.month-1] {
		return fmt.Errorf("year %d does not have %d days for month %s", d.year, d.day, NepaliMonth(d.month)[0])
	}
	return nil
}
