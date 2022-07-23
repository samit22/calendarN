package cmd

import "github.com/spf13/cobra"

const (
	Version = "0.0.2"
	IsoDate = "2006-01-02"
)

var PostRunMsg = func(cmd *cobra.Command, args []string) {
	log.Infof("\n \t\t ©Samit Ghimire 2022 \n")
}
