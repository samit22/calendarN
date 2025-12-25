package cmd

import (
	"os"
	"testing"

	"github.com/samit22/calendarN/countdown"
)

func Test_listCountdowns(t *testing.T) {
	dir := t.TempDir()
	filePath = dir + "/" + fileName

	tests := []struct {
		name       string
		createData bool
		data       string
		want       map[string]countdown.Response
	}{
		{
			name: "When file does not exist",
			want: map[string]countdown.Response{},
		},
		{
			name:       "When file exists it reads the file",
			createData: true,
			data:       `sam :: 3030-01-02`,
			want: map[string]countdown.Response{
				"sam": {},
			},
		},
		{
			name:       "when date is invalid it ignores the line",
			createData: true,
			data: `sam :: 3030-01-02
			bad :: 030-01-02`,
			want: map[string]countdown.Response{
				"sam": {},
			},
		},
		{
			name:       "when date is in past it ignores the line",
			createData: true,
			data: `sam :: 3030-02-02
			bad :: 1930-01-02
			good :: 3030-01-02 02:00:00`,
			want: map[string]countdown.Response{
				"sam":  {},
				"good": {},
			},
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

			got := listCountdowns()
			if len(got) != len(tt.want) {
				t.Errorf("listCountdowns() = %v, want %v", got, tt.want)
			}
		})
	}
}
