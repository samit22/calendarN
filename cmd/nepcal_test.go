package cmd

import (
	"testing"

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
