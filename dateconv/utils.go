package dateconv

import (
	"fmt"
	"strconv"
	"strings"
)

func GetDaysForMonth(year int, month int) (int, error) {
	d := Date{
		year:  year,
		month: month,
		day:   1,
	}
	err := d.Validate()
	if err != nil {
		return 0, err
	}
	return nepaliDates[year][month-1], nil
}

func EnglishToNepaliNumber(num int) string {
	return englishNumberToNepali(num)
}

func englishNumberToNepali(num int) (res string) {
	min2Num := fmt.Sprintf("%02d", num)
	stringN := strings.Split(min2Num, "")
	for _, n := range stringN {
		res += nepaliNum(n)
	}
	return
}
func sToI(inp string) int {
	i, _ := strconv.ParseInt(inp, 10, 0)
	return int(i)
}

func parseToNepaliDate(date string) (*Date, error) {
	if date == "" {
		return nil, fmt.Errorf("date cannot be empty")
	}
	spt := strings.Split(date, "-")
	if len(spt) != 3 {
		return nil, fmt.Errorf("invalid date format, should be 2079-11-02")
	}
	d := Date{
		year:  sToI(spt[0]),
		month: sToI(spt[1]),
		day:   sToI(spt[2]),
	}
	err := d.Validate()
	if err != nil {
		return nil, err
	}
	return &d, nil
}
