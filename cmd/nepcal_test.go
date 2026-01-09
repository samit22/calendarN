package cmd

import (
	"testing"

	"github.com/samit22/calendarN/dateconv"
	"github.com/stretchr/testify/assert"
)

func Test_parseArgsAndGenerate(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantC    Calendar
		wantDeep bool
	}{
		{
			name: "set date time",
			args: args{
				args: []string{"2079-05"},
			},
			wantC: Calendar{
				Year:  2079,
				Month: 05,
				Days:  31,
			},
			wantDeep: true,
		},
		{
			name: "invalid date time",
			args: args{
				args: []string{"abcd"},
			},
			wantC:    Calendar{},
			wantDeep: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC := parseArgsAndGenerate(tt.args.args)
			if tt.wantDeep {
				assert.Equal(t, tt.wantC.Year, gotC.Year)
				assert.Equal(t, tt.wantC.Month, gotC.Month)
				assert.Equal(t, tt.wantC.Days, gotC.Days)
			} else {
				if gotC.Year < 2000 {
					t.Errorf("should have received year > 2000 got %+v", gotC.Year)
				}
				if gotC.Month < 1 {
					t.Errorf("should have received month > 0 got %+v", gotC.Year)
				}
				if gotC.Days < 27 {
					t.Errorf("should have received days > 27 got %+v", gotC)
				}
			}
		})
	}
}

// Test for Nepali calendar last row fix - ensures all days are captured in Rows
func Test_generateNepCalendar_LastRowIncluded(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		expectedDays int
	}{
		{
			name:         "Baisakh 2079 - 31 days",
			year:         2079,
			month:        1,
			expectedDays: 31,
		},
		{
			name:         "Jestha 2079 - 31 days",
			year:         2079,
			month:        2,
			expectedDays: 31,
		},
		{
			name:         "Bhadra 2079 - 31 days",
			year:         2079,
			month:        5,
			expectedDays: 31,
		},
		{
			name:         "Chaitra 2079 - 30 days",
			year:         2079,
			month:        12,
			expectedDays: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thisNep, err := dateconv.NewDate(tt.year, tt.month, 15)
			if err != nil {
				t.Fatalf("Failed to create date: %v", err)
			}

			c := generateNepCalendar(tt.year, tt.month, thisNep)

			if c.Days != tt.expectedDays {
				t.Errorf("generateNepCalendar() Days = %v, want %v", c.Days, tt.expectedDays)
			}

			// Count total days in all rows
			totalDaysInRows := 0
			for _, row := range c.Rows {
				for _, cell := range row {
					if !cell.Blank {
						totalDaysInRows++
					}
				}
			}

			if totalDaysInRows != tt.expectedDays {
				t.Errorf("generateNepCalendar() total days in Rows = %v, want %v (last row may be missing)",
					totalDaysInRows, tt.expectedDays)
			}

			// Verify last row exists and contains days
			if len(c.Rows) > 0 {
				lastRow := c.Rows[len(c.Rows)-1]
				hasNonBlankDay := false
				for _, cell := range lastRow {
					if !cell.Blank {
						hasNonBlankDay = true
						break
					}
				}
				if !hasNonBlankDay {
					t.Errorf("generateNepCalendar() last row should contain at least one non-blank day")
				}
			}
		})
	}
}

// Test that Nepali calendar rows contain correct day sequence
func Test_generateNepCalendar_DaySequence(t *testing.T) {
	year := 2079
	month := 5
	thisNep, err := dateconv.NewDate(year, month, 15)
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	c := generateNepCalendar(year, month, thisNep)

	// Collect all non-blank days
	var days []int
	for _, row := range c.Rows {
		for _, cell := range row {
			if !cell.Blank {
				days = append(days, cell.Day)
			}
		}
	}

	// Verify days are sequential from 1 to Days
	for i, day := range days {
		expected := i + 1
		if day != expected {
			t.Errorf("Day sequence broken: position %d has day %d, want %d", i, day, expected)
		}
	}

	if len(days) != c.Days {
		t.Errorf("Total days collected = %d, want %d", len(days), c.Days)
	}
}
