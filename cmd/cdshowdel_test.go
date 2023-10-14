package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_showCountdown(t *testing.T) {
	dir := t.TempDir()
	filePath = dir + "/" + fileName
	type args struct {
		args []string
		name string
	}
	tests := []struct {
		createData bool
		data       string
		name       string
		args       args
		wantName   string
		wantErr    string
		err        bool
	}{
		{
			name:       "If the countdown with name exists it returns the countdown",
			createData: true,
			data:       `sam :: 3030-01-02`,
			args: args{
				args: []string{},
				name: "sam",
			},
			wantName: "sam",
		},
		{
			name: "If the countdown with name does not exists it returns empty",

			args: args{
				args: []string{},
				name: "sam",
			},
			wantName: "",
			wantErr:  "no such file or directory",
			err:      true,
		},
		{
			name:       "When time is in past it returns error",
			createData: true,
			data:       `sam :: 2020-01-02`,
			args: args{
				args: []string{},
				name: "sam",
			},
			wantName: "",
			wantErr:  "time in past t: 2020-01-02 00:00:00 +0000 UTC",
			err:      true,
		},
		{
			name:       "When time is in invalid format it returs error",
			createData: true,
			data:       `sam :: 220-01-02`,
			args: args{
				args: []string{},
				name: "sam",
			},
			wantName: "",
			err:      true,
			wantErr:  `parsing time "220-01-02 00:00:00" as "2006-01-02 15:04:05": cannot parse "01-02 00:00:00" as "2006"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name = tt.args.name
			if tt.createData {
				os.WriteFile(filePath, []byte(tt.data), 0644)
				t.Cleanup(func() {
					os.Remove(filePath)
				})
			}
			gotName, gotErr := showCountdown(tt.args.args)

			if !tt.err && gotName != tt.wantName {
				t.Errorf("showCountdown() = %v, want %v", gotName, tt.wantName)
			}

			if tt.err && !strings.Contains(gotErr.Error(), tt.wantErr) {
				t.Errorf("Err showCountdown() = %v, want %v", gotErr.Error(), tt.wantErr)
			}

		})
	}
}

func Test_deleteCountdown(t *testing.T) {
	dir := t.TempDir()
	filePath = dir + "/" + fileName
	type args struct {
		args []string
		name string
	}
	tests := []struct {
		createData bool
		data       string
		name       string
		args       args
		wantErr    bool
	}{
		{

			name:       "If the countdown with name exists it deletes the countdown",
			createData: true,
			data:       `sam :: 3030-01-02`,
			args: args{
				args: []string{},
				name: "sam",
			},
			wantErr: false,
		},
		{
			name: "If the countdown with name does not exists it returns nil",

			args: args{
				args: []string{},
				name: "sam",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.createData {
				os.WriteFile(filePath, []byte(tt.data), 0644)
				t.Cleanup(func() {
					os.Remove(filePath)
				})
			}
			if err := deleteCountdown(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("deleteCountdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
