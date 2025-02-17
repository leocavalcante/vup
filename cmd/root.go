package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vup",
	Short: "A semantic version manager",
	Long:  `Manipulates semantic version based string values.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("upgrade", "u", true, "Upgrade the major portion by the given value")
	rootCmd.PersistentFlags().BoolP("downgrade", "d", false, "Downgrade the major portion by the given value")
	rootCmd.PersistentFlags().IntP("value", "v", 1, "Value to increment or decrement the major portion")
}
