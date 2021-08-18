package main

import "github.com/spf13/cobra"

func newCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "completion [bash|zsh|fish]",
		Short:                 "Prints shell autocompletion scripts for gopwn",
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish"},
		Args:                  cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			switch args[0] {
			case "bash":
				err = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				err = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				err = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			}

			return err
		},
	}

	return cmd
}
