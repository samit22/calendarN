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
	"math/rand"
	"time"

	"github.com/samit22/calendarN/countdown"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

var (
	run      int64
	name     string
	cal      string
	timezone string
	save     bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(countDown)
	countDown.Flags().StringVarP(&name, "name", "n", randCharcater(5), "Name of the countdown to")
	countDown.Flags().StringVarP(&cal, "cal", "c", "eng", "Which calendar eng/nep (only eng supported now)")
	countDown.Flags().StringVarP(&timezone, "timezone", "t", "local", "Timezone to be used (only local supported now)")
	countDown.Flags().Int64VarP(&run, "run", "r", 5, "Run coutdown for n seconds, use -1 for infinite.")
	countDown.Flags().BoolVarP(&save, "save", "s", false, "Save the coutdown if flag is true")
}

// todayCmd represents the today command
var countDown = &cobra.Command{
	Use:     "countdown",
	Short:   "Gives countdown to a date.",
	Long:    `It should provide countdown to the the given date, in days, hours, min and second.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Errorf("empty date for countdown use nepcalN countdown 2022-08-18 or 2022-08-18 00:00:00")
			return
		}
		date := args[0]
		if date == "" {
			log.Errorf("empty date for countdown use nepcalN countdown 2022-08-18 or 2022-08-18 00:00:00")
			return
		}
		var tm, tz string
		if len(args) > 1 {
			tm = args[1]
		}

		ec, err := getEnglishCountdown(date, tm, tz)
		if err != nil {
			log.Errorf("failed to generate countdown err: %v", err)
			return
		}
		var infinite bool
		if run == -1 {
			infinite = true
		}
		if infinite {
			log.Infof("Running for infinite loop use Ctrl + C to exit\n")
		}
		log.Successf("Countdown for: %s\n", name)
		log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %d seconds\r", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))

		ticker := time.NewTicker(1 * time.Second)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					ec, _ = ec.Next()
					log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %d seconds\r", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))
				}
			}
		}()
		if !infinite {
			time.Sleep(time.Duration(run) * time.Second)
			ticker.Stop()
			done <- true
		} else {
			<-done
		}
	},
}

func getEnglishCountdown(date, time, timezone string) (*countdown.Response, error) {
	ct := countdown.NewCountdown()
	return ct.GetEnglishCountdown(date, time, timezone)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randCharcater(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
