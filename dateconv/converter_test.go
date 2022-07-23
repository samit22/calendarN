package dateconv

import (
	"testing"
)

func TestConverter_EtoN(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		wantT   string
		wantErr bool
	}{
		{
			name: "Invalid date format",
			args: struct {
				date string
			}{
				date: "440414",
			},
			wantT:   "",
			wantErr: true,
		},
		{
			name: "Day before range",
			args: struct {
				date string
			}{
				date: "1920-04-14",
			},
			wantT:   "2001-01-02",
			wantErr: true,
		},

		{
			name: "start of the seq",
			args: struct {
				date string
			}{
				date: "1944-04-14",
			},
			wantT: "2001-01-02",
		},
		{
			name: "new date",
			args: struct {
				date string
			}{
				date: "1988-01-15",
			},
			wantT: "2044-10-01",
		},
	}
	for _, tt := range tests {
		t.Log(tt.name)
		{
			t.Run(tt.name, func(t *testing.T) {
				c := &Converter{}
				got, err := c.EtoN(tt.args.date)

				if tt.wantErr {
					if err == nil {
						t.Errorf("EtoN() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}
				strDate := got.RomanFullDate()
				if strDate != tt.wantT {
					t.Errorf("EtoN() got = %v, want %v", strDate, tt.wantT)
				}
			})
		}
	}
}

func TestConverter_NtoE(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		wantT   string
		wantErr bool
	}{
		{
			name: "error empty date",
			args: struct {
				date string
			}{
				date: "",
			},
			wantErr: true,
		},
		{
			name: "error missing day",
			args: struct {
				date string
			}{
				date: "2000-01",
			},
			wantErr: true,
		},
		{
			name: "error bad date",
			args: struct {
				date string
			}{
				date: "1992-01-01",
			},
			wantErr: true,
		},
		{
			name: "error bad date month",
			args: struct {
				date string
			}{
				date: "2000-01-31",
			},
			wantErr: true,
		},
		{
			name: "start of the seq",
			args: struct {
				date string
			}{
				date: "2000-01-01",
			},
			wantT: "1943-04-14",
		},
		{
			name: "new date",
			args: struct {
				date string
			}{
				date: "2045-10-01",
			},
			wantT: "1989-01-14",
		},
		{
			name: "Check curernt date",
			args: struct {
				date string
			}{
				date: "2079-05-01",
			},
			wantT: "2022-08-17",
		},
	}
	for _, tt := range tests {
		t.Log(tt.name)
		{
			t.Run(tt.name, func(t *testing.T) {
				c := &Converter{}
				gotT, err := c.NtoE(tt.args.date)
				if tt.wantErr {
					if err == nil {
						t.Errorf("EtoN() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}
				if gotT == nil {
					t.Errorf("NtoE() unexpected error = %v", err)
					return
				}
				gotTStr := gotT.Format(IsoDate)
				if gotTStr != tt.wantT {
					t.Errorf("NtoE() gotT = %v, want %v", gotTStr, tt.wantT)
				}
			})
		}
	}
}

func Test_findTotalDays(t *testing.T) {
	type args struct {
		d Date
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Check date or same day",
			args: struct {
				d Date
			}{
				d: Date{
					year:  2000,
					month: 1,
					day:   1,
				},
			},
			want: 0,
		},
		{
			name: "Check date or next year",
			args: struct {
				d Date
			}{
				d: Date{
					year:  2001,
					month: 1,
					day:   5,
				},
			},
			want: 369,
		},
		{
			name: "Check date or next month",
			args: struct {
				d Date
			}{
				d: Date{
					year:  2001,
					month: 2,
					day:   5,
				},
			},
			want: 369,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTotalDays(tt.args.d); got != tt.want {
				t.Errorf("findTotalDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
