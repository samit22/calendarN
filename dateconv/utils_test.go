package dateconv

import (
	"reflect"
	"testing"
)

func TestGetDaysForMonth(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "invalid date",
			args: args{
				year:  1990,
				month: 1,
			},
			wantErr: true,
		},
		{
			name: "valid date",
			args: args{
				year:  2035,
				month: 4,
			},
			want: 32,
		},
		{
			name: "valid date 29",
			args: args{
				year:  2008,
				month: 7,
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDaysForMonth(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDaysForMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDaysForMonth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_englishNumberToNepali(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name: "One number",
			args: args{
				num: 4,
			},
			wantRes: "०४",
		},
		{
			name: "all numbers",
			args: args{
				num: 1234567890,
			},
			wantRes: "१२३४५६७८९०",
		},
		{
			name: "all numbers",
			args: args{
				num: 1,
			},
			wantRes: "०१",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := englishNumberToNepali(tt.args.num); gotRes != tt.wantRes {
				t.Errorf("englishNumberToNepali() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_parseToNepaliDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseToNepaliDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToNepaliDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseToNepaliDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sToI(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sToI(tt.args.inp); got != tt.want {
				t.Errorf("sToI() = %v, want %v", got, tt.want)
			}
		})
	}
}
