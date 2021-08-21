package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hupe1980/gopwn"
	"github.com/spf13/cobra"
)

func newCaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cave [file] [size]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires a file and a size argument")
			}
			if _, err := strconv.Atoi(args[1]); err != nil {
				return err
			}
			return nil
		},
		Short:         "Search for code caves",
		SilenceUsage:  true,
		SilenceErrors: true,
		Example:       "gopwn cave /usr/bin/ping 200",
		RunE: func(cmd *cobra.Command, args []string) error {
			size, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			fh, bt, err := gopwn.OpenFile(args[0])
			if err != nil {
				return err
			}
			var caves []gopwn.Cave
			switch bt {
			case gopwn.BINTYPE_ELF:
				elf, err := gopwn.NewELFFromReader(fh)
				if err != nil {
					return err
				}
				defer elf.Close()
				caves = elf.Caves(size)
			case gopwn.BINTYPE_PE:
				pe, err := gopwn.NewPEFromReader(fh)
				if err != nil {
					return err
				}
				defer pe.Close()
				caves = pe.Caves(size)
			case gopwn.BINTYPE_MACHO:
				macho, err := gopwn.NewMACHOFromReader(fh)
				if err != nil {
					return err
				}
				defer macho.Close()
				caves = macho.Caves(size)
			}

			if len(caves) == 0 {
				fmt.Println("\n[-] NO CAVE DETECTED!")
				return nil
			}

			for _, cave := range caves {
				fmt.Println("\n[+] CAVE DETECTED!")
				fmt.Printf("[!] Section Name: %s\n", cave.SectionName)
				fmt.Printf("[!] Section Flags: %s\n", cave.Infos)
				fmt.Printf("[!] Virtual Address: %#x\n", cave.Addr)
				fmt.Printf("[!] Cave Begin: %#x\n", cave.Begin)
				fmt.Printf("[!] Cave End: %#x\n", cave.End)
				fmt.Printf("[!] Cave Size: %#x (%d bytes)\n", cave.Size, cave.Size)
			}

			return nil
		},
	}

	return cmd
}
