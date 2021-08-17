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

type cyclicOptions struct {
	alphabet         string
	distSubseqLength int
}

func newCyclicCreateCmd() *cobra.Command {
	opts := &cyclicOptions{}
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
			pattern := gopwn.Cyclic(length, func(o *gopwn.CyclicOptions) {
				o.Alphabet = opts.alphabet
				o.DistSubseqLength = opts.distSubseqLength
			})

			fmt.Print(pattern)
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.alphabet, "alphabet", "a", "abcdefghijklmnopqrstuvwxyz", "alphabet to use in the cyclic pattern (optional)")
	cmd.Flags().IntVarP(&opts.distSubseqLength, "dslength", "n", 4, "size of the unique subsequences (optional)")

	return cmd
}

func newCyclicFinderCmd() *cobra.Command {
	opts := &cyclicOptions{}
	cmd := &cobra.Command{
		Use: "find [subseq]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a subseq argument")
			}
			return nil
		},
		Short:         "Cyclic finder",
		SilenceUsage:  true,
		SilenceErrors: true,
		Example:       "gopwn cyclic find",
		RunE: func(cmd *cobra.Command, args []string) error {
			subseq := args[0]
			offset := gopwn.CyclicFind([]byte(subseq), func(o *gopwn.CyclicOptions) {
				o.Alphabet = opts.alphabet
				o.DistSubseqLength = opts.distSubseqLength
			})

			fmt.Print(offset)
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.alphabet, "alphabet", "a", "abcdefghijklmnopqrstuvwxyz", "alphabet to use in the cyclic pattern (optional)")
	cmd.Flags().IntVarP(&opts.distSubseqLength, "dslength", "n", 4, "size of the unique subsequences (optional)")

	return cmd
}
