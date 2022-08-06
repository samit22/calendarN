package cmd

import (
	"testing"
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
				if gotC.Year != tt.wantC.Year && gotC.Month != tt.wantC.Month && gotC.Days != tt.wantC.Days {
					t.Errorf("parseArgsAndGenerate() = %v, want %v", gotC, tt.wantC)
				}
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
