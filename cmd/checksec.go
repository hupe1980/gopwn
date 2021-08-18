package main

import (
	"errors"
	"fmt"

	"github.com/hupe1980/gopwn"
	"github.com/spf13/cobra"
)

func newChecksecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "checksec [elf]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a elf argument")
			}
			return nil
		},
		Short:         "Check binary security settings",
		SilenceUsage:  true,
		SilenceErrors: true,
		Example:       "gopwn checksec /usr/bin/ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			elf, err := gopwn.NewELF(args[0])
			if err != nil {
				return err
			}
			fmt.Println(elf.Checksec())
			return nil
		},
	}

	return cmd
}
