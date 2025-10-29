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

		if cmd.Name() == "rc" {
			p, err := cmd.Flags().GetBool("promote")
			if err != nil {
				return err
			}

			if p {
				v.RC.Clear()
				_, err := fmt.Println(v)
				return err
			}
		}

		if up && !rb {
			f(v).Inc(val)
			if cmd.Name() != "rc" {
				rc, err := cmd.Flags().GetBool("rc")
				if err != nil {
					return err
				}
				if rc {
					switch cmd.Name() {
					case "major":
						v.Minor.Clear()
						v.Patch.Clear()
					case "minor":
						v.Patch.Clear()
					}
					v.RC.Set(1)
				}
			}
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
