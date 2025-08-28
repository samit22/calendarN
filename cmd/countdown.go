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
	"os/user"
	"strings"
	"time"

	"github.com/samit22/calendarN/countdown"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

var (
	run       int64
	name      string
	cal       string
	timezone  string
	save      bool
	overwrite bool
)
var (
	currentUser, _ = user.Current()
	homeDir        = currentUser.HomeDir
	folderPath     = homeDir + "/.calendarN"
	fileName       = ".calendar"
	filePath       = folderPath + "/" + fileName
)

func init() {
	rootCmd.AddCommand(countDown)
	countDown.Flags().StringVarP(&name, "name", "n", randCharcater(5), "Name of the countdown to")
	countDown.Flags().StringVarP(&cal, "cal", "c", "eng", "Which calendar eng/nep (only eng supported now)")
	countDown.Flags().StringVarP(&timezone, "timezone", "t", "local", "Timezone to be used (only local supported now)")
	countDown.Flags().Int64VarP(&run, "run", "r", 5, "Run countdown for n seconds, use -1 for infinite.")
	countDown.Flags().BoolVarP(&save, "save", "s", false, "Save the countdown if flag is true")
	countDown.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "Overwrite existing name")

}

// todayCmd represents the today command
var countDown = &cobra.Command{
	Use:     "countdown",
	Short:   "Gives countdown to a date.",
	Long:    `It should provide countdown to the the given date, in days, hours, min and second.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		runCountdown(args)
	},
}

func runCountdown(args []string) (cdTimes int, err error) {
	if len(args) == 0 {
		log.Errorf("empty date for countdown use calendarN countdown 2022-08-18 or 2022-08-18 00:00:00")
		err = fmt.Errorf("empty date")
		return
	}
	date := args[0]
	if date == "" {
		log.Errorf("empty date for countdown use calendarN countdown 2022-08-18 or 2022-08-18 00:00:00")
		err = fmt.Errorf("empty date")
		return
	}
	var tm, tz string
	if len(args) > 1 {
		tm = args[1]
	}

	ec, err := getEnglishCountdown(date, tm, tz)
	if err != nil {
		log.Errorf("failed to generate countdown err: %v", err)
		err = fmt.Errorf("countdown generation failed")
		return
	}
	if save {
		loadDataToFile(name, date)
	}
	var infinite bool
	if run == -1 {
		infinite = true
	}
	if infinite {
		log.Infof("Running for infinite loop use Ctrl + C to exit\n")
	}
	log.Successf("Countdown for: %s\n", name)
	log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %02d seconds\r ", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))
	cdTimes++
	if run != -1 && run < 2 {
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				ec, _ = ec.Next()
				cdTimes++
				log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %02d seconds\r ", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))
			}
		}
	}()
	if !infinite {
		time.Sleep(time.Duration(run-1) * time.Second)
		ticker.Stop()
		done <- true
	} else {
		<-done
	}
	return
}

func loadDataToFile(name, date string) error {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			err := os.MkdirAll(folderPath, os.ModePerm)
			if err != nil {
				log.Errorf("Error creating folder: %v", err)
				return err
			}
		}

	}
	var existingData = formatSavedData(data)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorf("Error opening or creating file: %v", err)
		return err
	}
	defer file.Close()

	if !overwrite {
		if _, ok := existingData[name]; ok {
			log.Errorf("Can't save, same name already exists.\n")
			return nil
		}
	}
	_, err = file.Write([]byte(fmt.Sprintf("%s :: %s\n", name, date)))
	if err != nil {
		log.Errorf("Error writing to file: %v", err)
	}
	return err
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

func formatSavedData(data []byte) map[string]string {
	var existingData = make(map[string]string)
	if len(data) > 0 {
		readData := strings.Split(string(data), "\n")
		for _, row := range readData {
			split := strings.Split(row, " :: ")
			if len(split) > 1 {
				existingData[split[0]] = split[1]
			}
		}
	}
	return existingData
}
