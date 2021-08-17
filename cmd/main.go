package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	version = "dev"
)

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "gopwn",
		Version:       version,
		Short:         "",
		Long:          "",
		SilenceErrors: true,
	}
	cmd.AddCommand(
		newCyclicCmd(),
	)
	return cmd
}

func main() {
	rootCmd := newRootCmd(version)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, promptui.IconBad, err)
		os.Exit(1)
	}
}
