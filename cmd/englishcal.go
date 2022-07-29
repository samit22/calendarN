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
			generateCalendar(d.Year(), int(d.Month()), now)
			return
		}
		generateCalendar(now.Year(), int(now.Month()), now)
	},
}

func generateCalendar(year, month int, now time.Time) {
	var today int

	currentLocation := now.Location()
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	log.PrintColorf(logger.Magentacolor, "\t\t  %d %s\n", year, time.Month(month).String())
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
			}
		}
		if today == i {
			log.PrintBackgroundf(logger.BackgroundGreen, "%s", generateEngDay(i))
		} else {
			if counter == 0 || counter == 6 {
				log.PrintColorf(logger.Red, "%s", generateEngDay(i))
			} else {
				log.Printf("%s", generateEngDay(i))
			}

		}
		if counter != 6 {
			counter++

		} else {
			fmt.Print("\n")
			counter = 0
		}
		if i == lastOfMonth.Day() {
			fmt.Println()
		}

	}
	fmt.Println()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// englishcalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// englishcalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
