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

		wantRC := false
		if cmd.Name() != "rc" {
			wantRC, err = cmd.Flags().GetBool("rc")
			if err != nil {
				return err
			}
			if wantRC && rb {
				return fmt.Errorf("--rc cannot be combined with --downgrade")
			}
			if wantRC && !up {
				return fmt.Errorf("--rc requires an upgrade")
			}
			if wantRC && val == 0 {
				return fmt.Errorf("--rc requires a positive version step")
			}
		}

		if cmd.Name() == "rc" {
			p, err := cmd.Flags().GetBool("promote")
			if err != nil {
				return err
			}

			if p {
				if rb {
					return fmt.Errorf("--promote cannot be combined with --downgrade")
				}
				if !rcPresent(v.RC) {
					return fmt.Errorf("rc command requires an existing rc suffix")
				}
				v.RC.Clear()
				_, err = fmt.Fprintln(cmd.OutOrStdout(), v)
				return err
			}

			if !rcPresent(v.RC) {
				return fmt.Errorf("rc command requires an existing rc suffix")
			}
		}

		if val < 0 {
			return fmt.Errorf("value must not be negative")
		}

		if up && !rb {
			if err := f(v).Inc(val); err != nil {
				return err
			}
			if cmd.Name() != "rc" && val > 0 {
				if wantRC {
					switch cmd.Name() {
					case "major":
						v.Minor.Clear()
						v.Patch.Clear()
					case "minor":
						v.Patch.Clear()
					}
					v.RC.Set(1)
				} else {
					v.RC.Clear()
				}
			}
		}

		if rb {
			err := f(v).Dec(val)
			if err != nil {
				return err
			}
			if cmd.Name() != "rc" && val > 0 {
				v.RC.Clear()
			}
		}

		_, err = fmt.Fprintln(cmd.OutOrStdout(), v)
		return err
	}
}

func rcPresent(p vup.Part) bool {
	rc, ok := p.(interface{ Present() bool })
	return ok && rc.Present()
}
