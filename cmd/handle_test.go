package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func resetFlags(t *testing.T) {
	t.Helper()
	rootCmd.SetOut(nil)
	rootCmd.SetErr(nil)
	rootCmd.SetArgs(nil)
	for _, set := range []struct {
		flag  string
		value string
	}{
		{flag: "downgrade", value: "false"},
		{flag: "upgrade", value: "true"},
		{flag: "value", value: "1"},
	} {
		if err := rootCmd.PersistentFlags().Set(set.flag, set.value); err != nil {
			t.Fatal(err)
		}
	}
	if err := majorCmd.PersistentFlags().Set("rc", "false"); err != nil {
		t.Fatal(err)
	}
	if err := minorCmd.PersistentFlags().Set("rc", "false"); err != nil {
		t.Fatal(err)
	}
	if err := patchCmd.PersistentFlags().Set("rc", "false"); err != nil {
		t.Fatal(err)
	}
	if err := rcCmd.PersistentFlags().Set("promote", "false"); err != nil {
		t.Fatal(err)
	}
}

func runCmd(t *testing.T, args ...string) (string, error) {
	t.Helper()
	resetFlags(t)
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestHandleRCRules(t *testing.T) {
	t.Cleanup(func() { resetFlags(t) })

	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr string
	}{
		{
			name:    "rc upgrade on a final release",
			args:    []string{"rc", "1.2.3"},
			wantErr: "existing rc suffix",
		},
		{
			name: "decrement rc1 to rc0",
			args: []string{"rc", "-d", "1.2.3-rc1"},
			want: "1.2.3-rc0",
		},
		{
			name: "patch upgrade drops an existing rc",
			args: []string{"patch", "1.2.3-rc2"},
			want: "1.2.4",
		},
		{
			name: "patch --rc starts rc1",
			args: []string{"patch", "--rc", "1.2.3"},
			want: "1.2.4-rc1",
		},
		{
			name:    "downgrade rejects --rc",
			args:    []string{"patch", "--downgrade", "--rc", "1.2.3"},
			wantErr: "--rc cannot be combined with --downgrade",
		},
		{
			name: "promote removes the rc suffix",
			args: []string{"rc", "--promote", "1.2.3-rc2"},
			want: "1.2.3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runCmd(t, tt.args...)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want substring %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("output = %q, want %q", got, tt.want)
			}
		})
	}
}
