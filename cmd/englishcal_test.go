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
