package dateconv

import (
	"reflect"
	"testing"
)

func TestNepaliMonth(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		{
			name: "date",
			args: 1,
			want: []string{"Baisakh", "बैशाख"},
		},
		{
			name: "date",
			args: 13,
			want: []string{"Unknown", "थाछैन ब्रो"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NepaliMonth(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NepaliMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNepaliWeekDay(t *testing.T) {

	tests := []struct {
		name string
		args int
		want []string
	}{
		{
			name: "invalid week",
			args: 7,
			want: []string{"Unknown", "थाछैन ब्रो"},
		},
		{
			name: "valid week",
			args: 2,
			want: []string{"Mangalbaar", "मंगलवार", "मं"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NepaliWeekDay(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NepaliWeekDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nepaliNum(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nepaliNum(tt.args.num); got != tt.want {
				t.Errorf("nepaliNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
