package cmd

import (
	"time"

	"github.com/samit22/calendarN/dateconv"
	"github.com/samit22/calendarN/logger"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var convert = &cobra.Command{
	Use:   "convert [etn|nte]",
	Short: "Convert",
	Long: `Convert has two commands etn and nte
	Usages calendarN convert etn 2022-08-18`,
	ValidArgs: []string{"etn", "nte"},
	PostRun:   PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		dateConvert(args)
	},
}

func dateConvert(args []string) {
	argLength := len(args)
	if argLength > 0 {
		switch args[0] {
		case "etn":
			if argLength > 1 {
				converEtoN(args[1])
			} else {
				log.Errorf("Date is required for conversion. Usage calendarN etn '2022-08-18'\n")
			}

		case "nte":
			log.PrintColor(logger.Red, "To be implemented!!\n")
		default:
			log.Errorf("invalid argument use etn | nte")
		}

	}
}

func converEtoN(d string) (string, error) {
	_, err := time.Parse(IsoDate, d)
	if err != nil {
		return "", err
	}
	dc := dateconv.Converter{}
	nDate, err := dc.EtoN(d)
	log.Successf("Eng: %s => %s  || %s\n", d, nDate.RomanFullDate(), nDate.DevanagariFullDate())
	return nDate.RomanFullDate(), nil
}

func init() {
	rootCmd.AddCommand(convert)
}
