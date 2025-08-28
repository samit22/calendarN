package cmd

import (
	"fmt"

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

func dateConvert(args []string) error {
	argLength := len(args)

	if argLength != 2 {
		return fmt.Errorf("argument does not include `etn` or `nte` and date")
	}

	switch args[0] {
	case "etn":
		_, err := converEtoN(args[1])
		return err
	case "nte":
		log.PrintColor(logger.Red, "To be implemented!!\n")
		return fmt.Errorf("nte is not implemented")
	default:
		log.Errorf("invalid argument use etn | nte")
		return fmt.Errorf("argument is neither `etn` nor `nte`")
	}
}

func converEtoN(d string) (string, error) {
	dc := dateconv.Converter{}
	nDate, err := dc.EtoN(d)
	if err != nil {
		return "", err
	}
	log.Successf("Eng: %s => %s  || %s\n", d, nDate.RomanFullDate(), nDate.DevanagariFullDate())
	return nDate.RomanFullDate(), nil
}

func init() {
	rootCmd.AddCommand(convert)
}
