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
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/samit22/calendarN/dateconv"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

// englishcalCmd represents the englishcal command
var nepCalCmd = &cobra.Command{
	Use:   "nep",
	Short: "Shows the nepali calendar",
	Long: `Return current nepali calendar supports arguments 2000-2090 BS.
	Should be provided in the forma 2079-01 if specific month is required`,

	Run: func(cmd *cobra.Command, args []string) {
		parseArgsAndGenerate(args)
	},
	PostRun: PostRunMsg,
}

func parseArgsAndGenerate(args []string) (c Calendar) {
	var argsTry bool
	if len(args) > 0 && args[0] != "" {
		argsTry = true
	}
TryAgain:
	if argsTry {
		y, m := checkInputNepaliDate(args[0])
		conv := dateconv.Converter{}
		np, _ := conv.EtoN(time.Now().Format(IsoDate))
		d, err := dateconv.NewDate(y, m, np.Day())
		if err != nil {
			log.Errorf("invalid input error: %v\n", err)
			log.Info("Showing this month's calendar\n")
			argsTry = false
			goto TryAgain
		}
		c = generateNepCalendar(d.Year(), d.Month(), d)
		return
	}
	now := time.Now()
	conv := dateconv.Converter{}
	todayNep, _ := conv.EtoN(now.Format(IsoDate))
	c = generateNepCalendar(todayNep.Year(), todayNep.Month(), todayNep)
	return
}

func checkInputNepaliDate(inp string) (y, m int) {
	spt := strings.Split(inp, "-")
	if len(spt) != 2 {
		return
	}
	y64, _ := strconv.ParseInt(spt[0], 10, 0)
	m64, _ := strconv.ParseInt(spt[1], 10, 0)
	return int(y64), int(m64)
}

func generateNepCalendar(year, month int, thisNep *dateconv.Date) (c Calendar) {
	c.Year = year
	c.Month = month
	var now = time.Now()
	var today int
	conv := dateconv.Converter{}
	nepFirstDayOfMonth := fmt.Sprintf("%d-%d-%d", year, month, 1)
	d, err := conv.NtoE(nepFirstDayOfMonth)
	if err != nil {
		log.Errorf("failed to get first day of the month err: %v", err)
		return
	}

	log.PrintColorf(logger.Magentacolor, "\t   %s %s\n", thisNep.DevanagariYear(), thisNep.DevanagariMonth())

	if now.Year() == thisNep.GetEnglishDate().Year() && now.Month() == thisNep.GetEnglishDate().Month() {
		today = thisNep.Day()
	}
	for i := range []int{0, 1, 2, 3, 4, 5, 6} {
		if i == 0 || i == 6 {
			log.PrintColorf(logger.Red, "%s ", dateconv.NepaliWeekDay(i)[2])
		} else if i == 4 {
			log.Printf("%s", dateconv.NepaliWeekDay(i)[2])
		} else {
			log.Printf("%s ", dateconv.NepaliWeekDay(i)[2])
		}
		if i == 6 {
			fmt.Println()
		}
	}
	counter := 0
	daysInThisMonth, _ := dateconv.GetDaysForMonth(year, month)
	c.Days = daysInThisMonth
	finalRow := [][]Row{}
	rows := []Row{}
	for i := 1; i <= daysInThisMonth; i++ {
		if i == 1 {
			noOfTab := int(d.Weekday())
			counter += noOfTab
			for j := 1; j <= noOfTab; j++ {
				if j == noOfTab {
					log.Printf("%s", " ")
				} else {
					log.Printf("%s", "  ")
				}
				rows = append(rows, Row{Blank: true})
			}
		}
		if today == i {
			log.PrintBackgroundf(logger.BackgroundGreen, "%s", dateconv.EnglishToNepaliNumber(i))
			rows = append(rows, Row{
				Day:   i,
				Today: true,
			})
		} else {
			if counter == 0 || counter == 6 {
				log.PrintColorf(logger.Red, "%s", dateconv.EnglishToNepaliNumber(i))
			} else {
				log.Printf("%s", dateconv.EnglishToNepaliNumber(i))
			}
			rows = append(rows, Row{
				Day: i,
			})
		}
		if counter != 6 {
			counter++

		} else {
			fmt.Print("\n")
			finalRow = append(finalRow, rows)
			rows = []Row{}
			counter = 0
		}
		if i == daysInThisMonth {
			fmt.Println()
		}

	}
	c.Rows = finalRow
	fmt.Println()
	return
}

func init() {
	rootCmd.AddCommand(nepCalCmd)
}
