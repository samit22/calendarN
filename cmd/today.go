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
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/samit22/calendarN/dateconv"
	"github.com/samit22/calendarN/logger"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var jsonOP, minifiedDates bool

type NepJson struct {
	YearEng    int    `json:"year_english"`
	MonthEng   int    `json:"month_english"`
	DayEng     int    `json:"day_english"`
	WeekDayEng int    `json:"week_day_english"`
	YearDays   int    `json:"year_days"`
	YearNep    string `json:"year"`
	MonthNep   string `json:"month"`
	DayNep     string `json:"day"`
	WeekDayNep string `json:"week_day"`
	FullDate   string `json:"full_date"`
}
type EngJson struct {
	Year        int     `json:"year"`
	Month       int     `json:"month"`
	MonthString string  `json:"month_string"`
	Day         int     `json:"day"`
	WeekDay     int     `json:"week_day"`
	FullDate    string  `json:"full_date"`
	YearDays    int     `json:"year_days"`
	Week        int     `json:"week"`
	Progess     float64 `json:"progress"`
}
type TodayJSON struct {
	English EngJson `json:"english"`
	Nepali  NepJson `json:"nepali"`
}

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Gives English and Nepali today's detail",
	Long: `Today gives the english and nepali dates for today.
It also has the number of days passed for the both the years
with the percentage.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		if minifiedDates {
			getMinifiedDate()
		} else {
			properResponse(getNepToday(), getToday())
		}
	},
}

func properResponse(n *dateconv.Date, e EngJson) TodayJSON {

	nep := NepJson{
		YearEng:    n.Year(),
		MonthEng:   n.Month(),
		DayEng:     n.Day(),
		WeekDayEng: int(n.WeekDay()),
		YearDays:   n.YearDay(),
		YearNep:    n.DevanagariYear(),
		MonthNep:   n.DevanagariMonth(),
		DayNep:     n.DevanagariDay(),
		WeekDayNep: n.DevanagariWeekDay(),
		FullDate:   fmt.Sprintf("%s, %s %s, %s", n.DevanagariWeekDay(), n.DevanagariDay(), n.DevanagariMonth(), n.DevanagariYear()),
	}
	tj := TodayJSON{
		English: e,
		Nepali:  nep,
	}
	if !jsonOP {
		return tj
	}
	op, _ := json.MarshalIndent(tj, "", " ")
	log.PrintColorf(logger.Cyan, "%s\n", op)

	return tj
}

func getNepToday() *dateconv.Date {
	dc := dateconv.Converter{}
	now := time.Now().Local().Format("2006-01-02")

	nDate, _ := dc.EtoN(now)
	fmt.Println()
	if jsonOP {
		return nDate
	}
	log.PrintColorf(logger.Green, "***************************\n")
	log.PrintColorf(logger.Cyan, "|         नेपाली आ ज       |\n")
	log.PrintColorf(logger.Green, "|-------------------------|\n")

	log.PrintColorf(logger.Cyan, "|  %s, %s %s, %s |\n", nDate.DevanagariWeekDay(), nDate.DevanagariDay(), nDate.DevanagariMonth(), nDate.DevanagariYear())
	log.PrintColorf(logger.Green, "|                         |\n")
	log.PrintColorf(logger.Cyan, "| यो वर्षको दिन: %s       |\n", dateconv.EnglishToNepaliNumber(nDate.YearDay()))
	log.PrintColorf(logger.Green, "***************************\n\n")
	return nDate
}

func getToday() EngJson {
	now := time.Now().Local()
	currentLocation := now.Location()
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, currentLocation)

	yearEnd := time.Date(now.Year(), 12, 31, 0, 0, 0, 0, currentLocation)
	elpDays := math.Floor(time.Since(yearStart).Seconds() / (24 * 60 * 60))
	totalDays := yearEnd.YearDay()

	percenntage := math.Floor(elpDays / float64(totalDays) * 100)
	_, week := now.ISOWeek()
	op := EngJson{
		Year:        now.Year(),
		Month:       int(now.Month()),
		MonthString: now.Month().String(),
		Day:         now.Day(),
		WeekDay:     int(now.Weekday()),
		FullDate:    now.Format("Monday, 02 January, 2006"),
		YearDays:    int(elpDays),
		Week:        week,
		Progess:     elpDays / float64(totalDays) * 100,
	}
	if jsonOP {
		return op
	}
	log.PrintColorf(logger.Green, "***************************\n")
	log.PrintColorf(logger.Cyan, "|       English Today     |\n")
	log.PrintColorf(logger.Green, "|-------------------------|\n")
	log.PrintColorf(logger.Cyan, "| %s |\n", now.Format("Monday, 02 January, 2006"))
	log.PrintColorf(logger.Green, "|                         |\n")
	log.PrintColorf(logger.Cyan, "| Day of the year: %d    |\n", int(elpDays))

	log.PrintColorf(logger.Cyan, "| Week of the year: %d    |\n", week)
	log.PrintColorf(logger.Green, "***************************\n\n")
	progressText := fmt.Sprintf(" %s%d progress: %s", logger.Cyan, now.Year(), logger.Reset)

	bar := progressbar.NewOptions(1000,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetDescription(progressText),
		progressbar.OptionSetRenderBlankState(false),
		progressbar.OptionSpinnerType(0),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]▉[reset]",
			SaucerHead:    "[green][reset]",
			SaucerPadding: " ",
			BarStart:      "▉",
			BarEnd:        "▉",
		}))

	for i := 0; i <= int(percenntage*10); i++ {
		bar.Add(1)
		time.Sleep(2 * time.Millisecond)
	}
	return op
}

func getMinifiedDate() {
	dc := dateconv.Converter{}
	now := time.Now().Local()

	nDate, _ := dc.EtoN(now.Format("2006-01-02"))
	log.PrintColorf(logger.Cyan, "%s, %s %s, %s | ", nDate.DevanagariWeekDay(), nDate.DevanagariDay(), nDate.DevanagariMonth(), nDate.DevanagariYear())
	log.PrintColorf(logger.Cyan, "%s\n", now.Format("Monday, 02 January, 2006"))

}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().BoolVarP(&jsonOP, "json", "j", jsonOP, "JSON output.")
	todayCmd.Flags().BoolVarP(&minifiedDates, "minified", "m", minifiedDates, "Minified today's date.")
}
