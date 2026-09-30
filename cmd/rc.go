package cmd

import (
	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

var rcCmd = &cobra.Command{
	Use:   "rc [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the rc portion of a semantic version",
	Long:  `Increase or decrease the rc portion of a semantic version.`,
	RunE: handle(func(v *vup.Version) vup.Part {
		return v.RC
	}),
}

func init() {
	rootCmd.AddCommand(rcCmd)
	rcCmd.PersistentFlags().BoolP("promote", "p", false, "Promote a RC version to a final release")
}
