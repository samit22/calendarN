package dateconv

import (
	"reflect"
	"testing"
	"time"
)

func TestDateMethods(t *testing.T) {
	d := &Date{
		year:    2010,
		month:   12,
		day:     1,
		weekDay: time.Saturday,
	}

	t.Log("Year() returns year of the date")
	{
		yr := d.Year()
		if yr != 2010 {
			t.Errorf("expect 2010 got %d", yr)
		}
	}
	t.Log("Month() returns month of the date")
	{
		mnth := d.Month()
		if mnth != 12 {
			t.Errorf("expect 2010 got %d", mnth)
		}
	}
	t.Log("Day() returns day of the date")
	{
		day := d.Day()
		if day != 1 {
			t.Errorf("expect 1 got %d", day)
		}
	}
	t.Log("WeekDay() returns day of the date")
	{
		wd := d.WeekDay()
		if wd != time.Saturday {
			t.Errorf("expect Saturday got %v", wd)
		}
	}
	t.Log("RomanFullDate() returns full date in yyy-mm-dd format")
	{
		rd := d.RomanFullDate()
		if rd != "2010-12-01" {
			t.Errorf("expect 2010-12-01 got %s", rd)
		}
	}

	t.Log("RomanMonth() returns moth in roman foramt")
	{
		rm := d.RomanMonth()
		if rm != "Chaitra" {
			t.Errorf("expect 2010-12-01 got %s", rm)
		}
	}

	t.Log("RomanWeekDay() returns week day in roman foramt")
	{
		rm := d.RomanWeekDay()
		if rm != "Sanibaar" {
			t.Errorf("expect Sanibaar got %s", rm)
		}
	}

	t.Log("DevanagariFullDate() returns week day in roman foramt")
	{
		rm := d.DevanagariFullDate()
		if rm != "२०१०-१२-०१" {
			t.Errorf("expect २०१०-१२-०१ got %s", rm)
		}
	}
	t.Log("DevanagariMonth() returns devanagari month")
	{
		rm := d.DevanagariMonth()
		if rm != "चैत्र" {
			t.Errorf("expect चैत्र got %s", rm)
		}
	}

	t.Log("DevanagariWeekDay() returns devanagari week day")
	{
		rm := d.DevanagariWeekDay()
		if rm != "शनिवार" {
			t.Errorf("expect शनिवार got %s", rm)
		}
	}

}

func TestNewDate(t *testing.T) {
	type args struct {
		y int
		m int
		d int
	}
	// Adjusting time zone issue of 15 min
	et := time.Date(2022, time.August, 18, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		{
			name: "error date",
			args: args{
				y: 1991,
				m: 1,
				d: 1,
			},
			wantErr: true,
		},

		{
			name: "error date",
			args: args{
				y: 2079,
				m: 5,
				d: 2,
			},
			want: &Date{
				year:    2079,
				month:   5,
				day:     2,
				weekDay: time.Thursday,
				engTime: &et,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDate(tt.args.y, tt.args.m, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotEngTime := got.engTime.Format(IsoDate)
				wantEngTime := tt.want.engTime.Format(IsoDate)
				if gotEngTime != wantEngTime {
					t.Errorf("NewDate() got eng time = %v, want %v", gotEngTime, wantEngTime)
					return
				}

				got.engTime = nil
				tt.want.engTime = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewDate() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDate_YearDay(t *testing.T) {
	type fields struct {
		year    int
		month   int
		day     int
		weekDay time.Weekday
		engTime *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "number of days",
			fields: fields{
				year:  2079,
				month: 1,
				day:   1,
			},
			want: 1,
		},
		{
			name: "number of days",
			fields: fields{
				year:  2079,
				month: 2,
				day:   2,
			},
			want: 33,
		},
		{
			name: "end of year",
			fields: fields{
				year:  2079,
				month: 12,
				day:   30,
			},
			want: 365,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				year:    tt.fields.year,
				month:   tt.fields.month,
				day:     tt.fields.day,
				weekDay: tt.fields.weekDay,
				engTime: tt.fields.engTime,
			}
			if got := d.YearDay(); got != tt.want {
				t.Errorf("YearDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test for Validate error messages fix
func TestDate_Validate_ErrorMessages(t *testing.T) {
	tests := []struct {
		name         string
		date         Date
		wantErr      bool
		errContains  string
	}{
		{
			name: "valid date",
			date: Date{
				year:  2079,
				month: 5,
				day:   15,
			},
			wantErr: false,
		},
		{
			name: "year below range",
			date: Date{
				year:  1999,
				month: 5,
				day:   15,
			},
			wantErr:     true,
			errContains: "invalid year 1999",
		},
		{
			name: "year above range",
			date: Date{
				year:  2091,
				month: 5,
				day:   15,
			},
			wantErr:     true,
			errContains: "invalid year 2091",
		},
		{
			name: "month below range",
			date: Date{
				year:  2079,
				month: 0,
				day:   15,
			},
			wantErr:     true,
			errContains: "invalid month 0",
		},
		{
			name: "month above range",
			date: Date{
				year:  2079,
				month: 13,
				day:   15,
			},
			wantErr:     true,
			errContains: "invalid month 13",
		},
		{
			name: "day below range",
			date: Date{
				year:  2079,
				month: 5,
				day:   0,
			},
			wantErr:     true,
			errContains: "invalid day 0 should be between 1 to 32",
		},
		{
			name: "day above range - error message should say 32 not 12",
			date: Date{
				year:  2079,
				month: 5,
				day:   33,
			},
			wantErr:     true,
			errContains: "invalid day 33 should be between 1 to 32",
		},
		{
			name: "day exceeds month days",
			date: Date{
				year:  2079,
				month: 1, // month 1 of 2079 has 31 days
				day:   32,
			},
			wantErr:     true,
			errContains: "does not have 32 days",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.date.Validate()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Validate() expected error but got none")
					return
				}
				if tt.errContains != "" && !containsString(err.Error(), tt.errContains) {
					t.Errorf("Validate() error = %q, want error containing %q", err.Error(), tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("Validate() unexpected error: %v", err)
				}
			}
		})
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > 0 && (containsString(s[1:], substr) || s[:len(substr)] == substr)))
}
