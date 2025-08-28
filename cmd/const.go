package cmd

import "github.com/spf13/cobra"

var Version string

const (
	IsoDate = "2006-01-02"
)

var PostRunMsg = func(cmd *cobra.Command, args []string) {
	log.Infof("\n\n \t\t  Follow @samit_gh in twitter\n")
}
