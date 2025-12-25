package cmd

import (
	"testing"
	"time"
)

func Test_checkArgsAndGenerateEngCalendar(t *testing.T) {
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
				args: []string{"2022-06"},
			},
			wantC: Calendar{
				Year:  2022,
				Month: 6,
				Days:  30,
			},
			wantDeep: true,
		},
		{
			name: "invalid date time returns current calendar",
			args: args{
				args: []string{"abcd"},
			},
			wantC: Calendar{
				Year:  time.Now().Year(),
				Month: int(time.Now().Month()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC := checkArgsAndGenerateEngCalendar(tt.args.args)
			if tt.wantDeep {
				if gotC.Year != tt.wantC.Year && gotC.Month != tt.wantC.Month && gotC.Days != tt.wantC.Days {
					t.Errorf("checkArgsAndGenerateEngCalendar() = %v, want %v", gotC, tt.wantC)
				}
			} else {
				if gotC.Year < 1900 || gotC.Month < 0 || gotC.Days < 27 {
					t.Errorf("should have received calendar got %+v", gotC)
				}
			}
		})
	}
}

// Test for calendar last row fix - ensures all days are captured in Rows
func Test_generateCalendar_LastRowIncluded(t *testing.T) {
	tests := []struct {
		name          string
		year          int
		month         int
		expectedDays  int
		checkLastRow  bool
	}{
		{
			name:          "June 2022 - month ends on Thursday",
			year:          2022,
			month:         6,
			expectedDays:  30,
			checkLastRow:  true,
		},
		{
			name:          "February 2022 - month ends on Monday",
			year:          2022,
			month:         2,
			expectedDays:  28,
			checkLastRow:  true,
		},
		{
			name:          "July 2022 - month ends on Sunday (Saturday)",
			year:          2022,
			month:         7,
			expectedDays:  31,
			checkLastRow:  true,
		},
		{
			name:          "October 2022 - month ends on Monday",
			year:          2022,
			month:         10,
			expectedDays:  31,
			checkLastRow:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Date(tt.year, time.Month(tt.month), 15, 0, 0, 0, 0, time.Local)
			c := generateCalendar(tt.year, tt.month, now)

			if c.Days != tt.expectedDays {
				t.Errorf("generateCalendar() Days = %v, want %v", c.Days, tt.expectedDays)
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
				t.Errorf("generateCalendar() total days in Rows = %v, want %v (last row may be missing)", 
					totalDaysInRows, tt.expectedDays)
			}

			// Verify last row exists and contains days
			if tt.checkLastRow && len(c.Rows) > 0 {
				lastRow := c.Rows[len(c.Rows)-1]
				hasNonBlankDay := false
				for _, cell := range lastRow {
					if !cell.Blank {
						hasNonBlankDay = true
						break
					}
				}
				if !hasNonBlankDay {
					t.Errorf("generateCalendar() last row should contain at least one non-blank day")
				}
			}
		})
	}
}

// Test that calendar rows contain correct day sequence
func Test_generateCalendar_DaySequence(t *testing.T) {
	year := 2022
	month := 6
	now := time.Date(year, time.Month(month), 15, 0, 0, 0, 0, time.Local)
	c := generateCalendar(year, month, now)

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
