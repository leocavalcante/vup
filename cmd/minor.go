package cmd

import (
	"fmt"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the minor portion of a semantic version",
	Long:  `Increase or decrease the minor portion of a semantic version.`,
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
			v.Minor.Inc(val)
		}

		if rb {
			err := v.Minor.Dec(val)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Println(v)
		return err
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)
	minorCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the minor portion by the given value")
	minorCmd.Flags().BoolP("downgrade", "d", false, "Downgrade the minor portion by the given value")
	minorCmd.Flags().IntP("value", "v", 1, "Value to increment or decrement the minor portion")
}
