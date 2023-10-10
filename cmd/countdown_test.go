package cmd

import (
	"testing"
	"time"
)

func Test_runCountdown(t *testing.T) {
	type args struct {
		args []string
		run  int64
	}
	newT := time.Now().Add(50 * time.Hour)

	tests := []struct {
		name        string
		args        args
		wantCdTimes int
		wantErr     bool
	}{
		{
			name: "empty arguments",
			args: args{
				args: nil,
			},
			wantErr: true,
		},
		{
			name: "empty date",
			args: args{
				args: []string{""},
			},
			wantErr: true,
		},
		{
			name: "failed date",
			args: args{
				args: []string{"20"},
			},
			wantErr: true,
		},
		{
			name: "empty time",
			args: args{
				args: []string{newT.Format("2006-01-02")},
				run:  1,
			},
			wantErr: false,
		},
		{
			name: "Date and time given",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  1,
			},
			wantErr: false,
		}, {
			name: "Date and time and run default time",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  5,
			},
			wantErr: false,
		},
		{
			name: "Date and time and run given",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			run = tt.args.run

			gotCdTimes, err := runCountdown(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runCountdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotCdTimes != int(tt.args.run) {
				t.Errorf("expected to have printed countdown %d times got %d", tt.args.run, gotCdTimes)
			}

		})
	}
}
