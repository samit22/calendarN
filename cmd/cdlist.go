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
	"os"
	"strings"
	"time"

	"github.com/samit22/calendarN/countdown"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	countDown.AddCommand(list)
	list.Flags().StringVarP(&name, "name", "n", randCharcater(5), "Name of the countdown to")
	list.Flags().Int64VarP(&run, "run", "r", 5, "Run countdown for n seconds, use -1 for infinite.")

}

// todayCmd represents the today command
var list = &cobra.Command{
	Use:     "all",
	Short:   "Gives list of all saved countdown.",
	Long:    `It will return all the saved countdowns.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		listCountdowns(args)
	},
}

func listCountdowns(args []string) map[string]countdown.Response {
	response := make(map[string]countdown.Response)
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		log.Errorf("no countdown saved, to save one use `calendarN countdown -n 'AUS' -s 2023-11-05`")
		return response
	}
	var existingData = formatSavedData(data)
	ct := countdown.NewCountdown()
	var cleanData bool
	var newData string
	for name, dateTime := range existingData {
		d := strings.Split(dateTime, " ")
		date := d[0]
		hour := "00:00:00"
		if len(d) > 1 {
			hour = strings.TrimSpace(d[1])
		}
		datetime := date + " " + hour

		parsedTime, err := time.Parse("2006-01-02 15:04:05", datetime)
		if err != nil {
			log.Errorf("failed to parse saved data: err: %v", err)
			cleanData = true
			continue
		}
		if parsedTime.Before(time.Now()) {
			log.Infof("time in past t: %s\n", parsedTime.String())
			cleanData = true
			continue
		}
		ec, err := ct.GetEnglishCountdown(date, hour, "")
		if err != nil {
			log.Errorf("failed to parse saved data: err: %v", err)
			cleanData = true
			continue
		}
		log.PrintColor(logger.Yellow, fmt.Sprintf("%s -> %s\n", name, dateTime))
		log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %d seconds\n\n", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))
		newData += fmt.Sprintf("%s :: %s\n", name, dateTime)
		response[name] = *ec
	}
	if cleanData {
		log.Infof("cleaning up old data")
		if newData != "" {
			err = os.WriteFile(filePath, []byte(newData), 0644)
			if err != nil {
				log.Errorf("failed to write new data: err: %v", err)
			}
		}
	}
	return response
}
