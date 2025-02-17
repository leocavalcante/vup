package cmd

import (
	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the minor portion of a semantic version",
	Long:  `Increase or decrease the minor portion of a semantic version.`,
	RunE: handle(func(v *vup.Version) vup.Part {
		return v.Minor
	}),
}

func init() {
	rootCmd.AddCommand(minorCmd)
}
