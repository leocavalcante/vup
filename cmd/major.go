package cmd

import (
	"fmt"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the major portion of a semantic version",
	Long:  `Increase or decrease the major portion of a semantic version.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		version := args[0]
		up, err := cmd.Flags().GetBool("upgrade")
		if err != nil {
			return err
		}

		rb, err := cmd.Flags().GetBool("downgrade")
		if err != nil {
			return err
		}

		val, err := cmd.Flags().GetInt("value")
		if err != nil {
			return err
		}

		v, err := vup.NewVersion(version)
		if err != nil {
			return err
		}

		if up {
			v.Major.Inc(val)
		}

		if rb {
			err := v.Major.Dec(val)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Println(v)
		return err
	},
}

func init() {
	rootCmd.AddCommand(majorCmd)
	majorCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the major portion by the given value")
	majorCmd.Flags().BoolP("downgrade", "d", false, "Downgrade the major portion by the given value")
	majorCmd.Flags().IntP("value", "v", 1, "Value to increment or decrement the major portion")
}
