/*
Copyright © 2022 Samit info@samitghimire.com.np

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
	"fmt"
	"math"
	"time"

	"github.com/samit22/calendarN/dateconv"
	"github.com/samit22/calendarN/logger"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Gives English and Nepali today's detail",
	Long: `Today gives the english and nepali dates for today.
It also has the number of days passed for the both the years
with the percentage.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		getNepToday()
		getToday()

	},
}

func getNepToday() {
	dc := dateconv.Converter{}
	now := time.Now().Local().Format("2006-01-02")

	nDate, _ := dc.EtoN(now)
	fmt.Println()
	log.PrintColorf(logger.Green, "***************************\n")
	log.PrintColorf(logger.Cyan, "|         नेपाली आ ज       |\n")
	log.PrintColorf(logger.Green, "|-------------------------|\n")

	log.PrintColorf(logger.Cyan, "|  %s, %s %s, %s |\n", nDate.DevanagariWeekDay(), nDate.DevanagariDay(), nDate.DevanagariMonth(), nDate.DevanagariYear())
	log.PrintColorf(logger.Green, "|                         |\n")
	log.PrintColorf(logger.Cyan, "| यो वर्षको दिन: %s       |\n", dateconv.EnglishToNepaliNumber(nDate.YearDay()))
	log.PrintColorf(logger.Green, "***************************\n\n")
}

func getToday() {
	defer func() {
		fmt.Println()
	}()
	now := time.Now().Local()
	currentLocation := now.Location()
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, currentLocation)

	yearEnd := time.Date(now.Year(), 12, 31, 0, 0, 0, 0, currentLocation)
	elpDays := time.Since(yearStart).Hours() / 24
	totalDays := yearEnd.YearDay()

	percenntage := math.Round(elpDays / float64(totalDays) * 100)
	log.PrintColorf(logger.Green, "***************************\n")
	log.PrintColorf(logger.Cyan, "|       English Today     |\n")
	log.PrintColorf(logger.Green, "|-------------------------|\n")
	log.PrintColorf(logger.Cyan, "| %s |\n", now.Format("Monday, 02 January, 2006"))
	log.PrintColorf(logger.Green, "|                         |\n")
	log.PrintColorf(logger.Cyan, "| Day of the year: %d    |\n", int(elpDays))
	_, week := now.ISOWeek()
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

}

func init() {
	rootCmd.AddCommand(todayCmd)

}
