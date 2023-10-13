/*
Copyright © calendarN

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"testing"

	"github.com/samit22/calendarN/countdown"
)

func Test_listCountdowns(t *testing.T) {
	dir := t.TempDir()
	filePath = dir + "/" + fileName

	type args struct {
		args []string
	}
	tests := []struct {
		createData bool
		data       string
		name       string
		args       args
		want       map[string]countdown.Response
	}{
		{
			name: "When file does not exist",
			args: args{
				args: []string{},
			},
			want: map[string]countdown.Response{},
		},
		{
			name:       "When file exists it reads the file",
			createData: true,
			data:       `sam :: 3030-01-02`,
			args: args{
				args: []string{},
			},
			want: map[string]countdown.Response{
				"sam": {},
			},
		},
		{
			name:       "when date is invalid it ignores the line",
			createData: true,
			data: `sam :: 3030-01-02
			bad :: 030-01-02`,
			args: args{
				args: []string{},
			},
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
			args: args{
				args: []string{},
			},
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

			got := listCountdowns(tt.args.args)
			if len(got) != len(tt.want) {
				t.Errorf("listCountdowns() = %v, want %v", got, tt.want)
			}
		})
	}
}
