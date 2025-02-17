package cmd

import (
	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the major portion of a semantic version",
	Long:  `Increase or decrease the major portion of a semantic version.`,
	RunE: handle(func(v *vup.Version) vup.Part {
		return v.Major
	}),
}

func init() {
	rootCmd.AddCommand(majorCmd)
}
