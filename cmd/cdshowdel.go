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
	"os"
	"strings"
	"time"

	"github.com/samit22/calendarN/countdown"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

func init() {
	countDown.AddCommand(show)
	show.Flags().StringVarP(&name, "name", "n", "", "Name of the countdown to	show")
	show.Flags().Int64VarP(&run, "run", "r", 5, "Run countdown for n seconds, use -1 for infinite.")

	deleteCD.Flags().StringVarP(&name, "name", "n", "", "Name of the countdown to	show")

}

var show = &cobra.Command{
	Use:     "show",
	Short:   "Gives the countdown for specific name.",
	Long:    `You can use this command to show the countdown for specific countdown.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		showCountdown(args)
	},
}

var deleteCD = &cobra.Command{
	Use:     "delete",
	Short:   "Deletes the countdown for specific name.",
	Long:    `You can use this command to delete the countdown for specific name.`,
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		deleteCountdown(args)
	},
}

func showCountdown(args []string) (string, error) {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		log.Errorf("no countdown saved, to save one use `calendarN countdown -n 'AUS' -s 2023-11-05`")
		return "", err
	}
	var existingData = formatSavedData(data)
	ct := countdown.NewCountdown()

	dateTime, exists := existingData[name]
	if !exists {
		log.Errorf("no countdown saved with name: %s", name)
		return "", fmt.Errorf("no countdown saved with name: %s", name)
	}
	d := strings.Split(dateTime, " ")
	date := d[0]
	hour := "00:00:00"
	if len(d) > 1 {
		hour = d[1]
	}
	datetime := date + " " + hour

	parsedTime, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		log.Errorf("failed to parse saved data: err: %v", err)
		return "", err
	}
	if parsedTime.Before(time.Now()) {
		log.Infof("time in past t: %s\n", parsedTime.String())
		return "", fmt.Errorf("time in past t: %s\n", parsedTime.String())
	}

	// this error is already handled in the above if block
	ec, _ := ct.GetEnglishCountdown(date, hour, "")

	log.PrintColor(logger.Yellow, fmt.Sprintf("%s -> %s\n", name, dateTime))
	log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %d seconds\n\n", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))

	if run != -1 && run < 2 {
		return name, nil
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

				log.PrintColor(logger.Yellow, fmt.Sprintf("%d days %d hours %d minutes %02d seconds\r ", ec.Days, ec.Hours, ec.Minutes, ec.Seconds))
			}
		}
	}()
	var infinite bool
	if run == -1 {
		infinite = true
	}
	if infinite {
		log.Infof("Running for infinite loop use Ctrl + C to exit\n")
	}
	if !infinite {
		time.Sleep(time.Duration(run-1) * time.Second)
		ticker.Stop()
		done <- true
	} else {
		<-done
	}
	return name, nil
}

func deleteCountdown(args []string) error {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		log.Errorf("no countdown saved, to save one use `calendarN countdown -n 'AUS' -s 2023-11-05`")
		return err
	}
	existingData := formatSavedData(data)

	_, exists := existingData[name]
	if !exists {
		log.Errorf("no countdown saved with name: %s", name)
		return fmt.Errorf("no countdown saved with name: %s", name)
	}
	for fileKey, _ := range existingData {
		if name == fileKey {
			delete(existingData, name)
		}
	}
	newData := ""
	for key, value := range existingData {
		newData += key + " :: " + value + "\n"
	}
	err = os.WriteFile(filePath, []byte(newData), 0644)
	if err != nil {
		log.Errorf("failed to save countdown err: %v", err)
		return err
	}
	log.Successf("Countdown deleted for: %s\n", name)
	return nil
}
