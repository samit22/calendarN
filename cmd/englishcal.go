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
	"time"

	"github.com/samit22/calendarN/data"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

var log = logger.Logger{}

// englishcalCmd represents the englishcal command
var englishcalCmd = &cobra.Command{
	Use:   "eng",
	Short: "Generates english calendar",
	Long: `It generated english calendar.
By default it generated calendar for current month.
You can also pass arugment in format yyyy-mm`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		checkArgsAndGenerateEngCalendar(args)
	},
}

func checkArgsAndGenerateEngCalendar(args []string) (c Calendar) {
	var argument bool
	if len(args) > 0 && args[0] != "" {
		argument = true
	}
	now := time.Now()
Default:
	if argument {
		newD := args[0] + "-01"
		d, err := time.Parse(IsoDate, newD)
		if err != nil {
			log.Errorf("Invalid date should be valid date in format yyyy-mm\n")
			log.Infof("Using current date for calendar\n")
			argument = false
			goto Default
		}
		c = generateCalendar(d.Year(), int(d.Month()), now)
		return
	}
	c = generateCalendar(now.Year(), int(now.Month()), now)
	return
}

func generateCalendar(year, month int, now time.Time) (c Calendar) {
	var today int
	c.Year = year
	c.Month = month

	currentLocation := now.Location()
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	c.Days = lastOfMonth.Day()

	finalRow := [][]Row{}
	rows := []Row{}

	log.PrintColorf(logger.Magentacolor, "\t   %d %s\n", year, time.Month(month).String())
	if year == now.Year() && month == int(now.Month()) {
		today = now.Day()
	}
	for i, d := range data.EngDayShort {
		if i == 0 || i == 6 {
			log.PrintColorf(logger.Red, "%s", d)
		} else {
			log.Printf("%s", d)
		}
		if i == len(data.EngDayShort)-1 {
			fmt.Println()
		}
	}
	counter := 0
	for i := 1; i <= lastOfMonth.Day(); i++ {
		if i == 1 {
			noOfTab := int(firstOfMonth.Weekday())
			counter += noOfTab
			for j := 1; j <= noOfTab; j++ {
				if j == noOfTab {
					log.Printf("%s", "  ")
				} else {
					log.Printf("%s", "   ")
				}
				rows = append(rows, Row{Blank: true})
			}
		}
		if today == i {
			log.PrintBackgroundf(logger.BackgroundGreen, "%s", generateEngDay(i))
			rows = append(rows, Row{
				Day:   i,
				Today: true,
			})
		} else {
			if counter == 0 || counter == 6 {
				log.PrintColorf(logger.Red, "%s", generateEngDay(i))
			} else {
				log.Printf("%s", generateEngDay(i))
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
		if i == lastOfMonth.Day() {
			fmt.Println()
		}
	}
	// Add the last row if it has any days
	if len(rows) > 0 {
		finalRow = append(finalRow, rows)
	}
	c.Rows = finalRow
	fmt.Println()
	return
}

func generateEngDay(inp int) string {
	if inp < 10 {
		return fmt.Sprintf(" %d ", inp)
	} else {
		return fmt.Sprintf("%d ", inp)
	}
}

func init() {
	rootCmd.AddCommand(englishcalCmd)
}
