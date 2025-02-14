package cmd

import (
	"fmt"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch [flags] [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Handles the patch portion of a semantic version",
	Long:  `Increase or decrease the patch portion of a semantic version.`,
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
			v.Patch.Inc(val)
		}

		if rb {
			err := v.Patch.Dec(val)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Println(v)
		return err
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
	patchCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the patch portion by the given value")
	patchCmd.Flags().BoolP("downgrade", "d", false, "Downgrade the patch portion by the given value")
	patchCmd.Flags().IntP("value", "v", 1, "Value to increment or decrement the patch portion")
}
