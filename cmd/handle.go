package cmd

import (
	"fmt"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/spf13/cobra"
)

func handle(f func(*vup.Version) vup.Part) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		version := args[0]

		v, err := vup.NewVersion(version)
		if err != nil {
			return err
		}

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

		if up && !rb {
			f(v).Inc(val)
		}

		if rb {
			err := f(v).Dec(val)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Println(v)
		return err
	}
}
