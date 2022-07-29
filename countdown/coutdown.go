package countdown

import (
	"fmt"
	"time"

	"github.com/samit22/calendarN/calerr"
)

const (
	acceptFormatDate     = "2006-01-02"
	acceptFormatDateTime = acceptFormatDate + " 15:04:05"
	durationDays         = 24 * 60 * 60
	durationHours        = 60 * 60
	durationMinutes      = 60
)

type CurrentTimeIface interface {
	GetCurrentTime() time.Time
	SetLocation(loc *time.Location)
}
type tm struct {
	location *time.Location
}

func (t *tm) SetLocation(loc *time.Location) {
	t.location = loc
}

func (t *tm) GetCurrentTime() time.Time {
	return time.Now().In(t.location)
}

// Response returns the countdown for the given date
//  Days = days for the countdown
//  Hours = hours remaining
//  Minutes = minutes remaining
//  Seconds = seconds remaining
type Response struct {
	Days           int `json:"days"`
	Hours          int `json:"hours"`
	Minutes        int `json:"minutes"`
	Seconds        int `json:"seconds"`
	originaEngDate *time.Time
	ct             CurrentTimeIface
}

type new struct {
	currentTime CurrentTimeIface
}

// NewCountdown initializes the countdown generation
func NewCountdown() *new {
	return &new{
		currentTime: &tm{
			location: time.Local,
		},
	}
}

// Next return the next duration of the countdown
// Can be used to generate the ever running timer using ticker
func (c *Response) Next() (*Response, error) {
	ot := c.originaEngDate
	if ot == nil {
		return nil, fmt.Errorf("time is not defined")
	}
	t := *ot
	diff := t.Sub(c.ct.GetCurrentTime()).Seconds()
	c.Days = int(diff / durationDays)
	rem := int(diff) % durationDays
	c.Hours = int(rem / durationHours)
	rem = rem % durationHours
	c.Minutes = int(rem / durationMinutes)
	c.Seconds = rem % durationMinutes
	return c, nil
}

// GetEnglishCountdown returns the countdown for the given date
func (n *new) GetEnglishCountdown(date, tm, timezone string) (*Response, error) {
	timeLoc := time.Local
	if timezone != "" {
		tl, err := time.LoadLocation(timezone)
		if err != nil {
			return nil, calerr.New(err, "Invalid timezone sent, please use standard timezone, eg: Asia/Kathmandu.", 422)
		}
		timeLoc = tl
		n.currentTime.SetLocation(tl)
	}
	t, err := time.ParseInLocation(acceptFormatDateTime, date+" "+tm, timeLoc)
	if err != nil {
		t, err = time.ParseInLocation(acceptFormatDate, date, timeLoc)
		if err != nil {
			return nil, calerr.New(err, "invalid date format, should follow iso8601 time stamp'2022-11-02 04:05:00' or date 2022-01-02 format", 422)
		}
	}
	tn := n.currentTime.GetCurrentTime()
	if t.Before(tn) {
		return nil, calerr.New(err, "invalid date, should be in the future", 422)
	}
	cr := &Response{
		originaEngDate: &t,
		ct:             n.currentTime,
	}
	return cr.Next()
}
