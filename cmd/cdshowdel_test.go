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
		want       map[string]countdown.Response
	}{
		{
			name:       "If the countdown with name exists it returns the countdown",
			createData: true,
			data:       `sam :: 3030-01-02`,
			args: args{
				args: []string{},
				name: "sam",
			},
			want: map[string]countdown.Response{
				"sam": {},
			},
		},
		{
			name: "If the countdown with name does not exists it returns nil",

			args: args{
				args: []string{},
				name: "sam",
			},
			want: map[string]countdown.Response{},
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
			got := showCountdown(tt.args.args)
			if len(got) != len(tt.want) {
				t.Errorf("showCountdown() = %v, want %v", got, tt.want)
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
