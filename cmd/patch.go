package cmd

import (
	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the patch portion of a semantic version",
	Long:  `Increase or decrease the patch portion of a semantic version.`,
	RunE: handle(func(v *vup.Version) vup.Part {
		return v.Patch
	}),
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
