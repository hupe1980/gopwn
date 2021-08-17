package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hupe1980/gopwn"
	"github.com/spf13/cobra"
)

func newCyclicCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "cyclic",
		Short:        "Generation of unique sequences",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newCyclicCreateCmd(),
		newCyclicFinderCmd(),
	)

	return cmd
}

type cyclicCreateOptions struct {
	length       int
	alphabet     string
	subseqLength int
}

func newCyclicCreateCmd() *cobra.Command {
	opts := &cyclicCreateOptions{}
	cmd := &cobra.Command{
		Use: "create [length]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a length argument")
			}
			if _, err := strconv.Atoi(args[0]); err != nil {
				return err
			}
			return nil
		},
		Short:         "Cyclic creator",
		SilenceUsage:  true,
		SilenceErrors: true,
		Example:       "gopwn cyclic create 200 > pattern.txt",
		RunE: func(cmd *cobra.Command, args []string) error {
			length, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			pattern := gopwn.AdvancedCylic(gopwn.AdvancedCyclicParams{
				Alphabet:     opts.alphabet,
				SubseqLength: opts.subseqLength,
				Length:       length,
			})
			fmt.Print(pattern)
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.alphabet, "alphabet", "a", "abcdefghijklmnopqrstuvwxyz", "alphabet to use in the cyclic pattern (optional)")
	cmd.Flags().IntVarP(&opts.subseqLength, "subseq", "n", 4, "size of the unique subsequences (optional)")

	return cmd
}

func newCyclicFinderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "find",
		Short:         "Cyclic finder",
		SilenceUsage:  true,
		SilenceErrors: true,
		Example:       "gopwn cyclic find",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
