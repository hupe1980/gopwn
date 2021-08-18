package gopwn

import (
	"debug/elf"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

type ELF struct {
	path    string // Path to the file
	file    *elf.File
	symbols []elf.Symbol
}

func NewELF(path string) (*ELF, error) {
	f, err := elf.Open(path)
	if err != nil {
		return nil, err
	}

	symbols, _ := f.Symbols()

	return &ELF{
		path:    path,
		file:    f,
		symbols: symbols,
	}, nil
}

// Canary checks whether the current binary is using stack canaries
func (e *ELF) Canary() bool {
	for _, symbol := range e.symbols {
		if symbol.Name == "__stack_chk_fail" {
			return true
		}
	}
	return false
}

// NX checks whether the current binary uses NX protections
func (e *ELF) NX() bool {
	for _, prog := range e.file.Progs {
		if uint32(prog.Type) == uint32(0x6474e551) { // PT_GNU_STACK
			if (uint32(prog.Flags) & uint32(elf.PF_X)) == 0 {
				return true
			}
		}
	}
	return false
}

func (e *ELF) Checksec() string {
	nx := map[bool]string{
		true:  color.GreenString("NX enabled"),
		false: color.RedString("NX disabled"),
	}
	stack := map[bool]string{
		true:  color.GreenString("Canary found"),
		false: color.RedString("No canary found"),
	}

	var builder strings.Builder
	writer := tabwriter.NewWriter(&builder, 0, 0, 3, ' ', 0)

	fmt.Fprintf(writer, "NX\t%s\n", nx[e.NX()])
	fmt.Fprintf(writer, "Stack\t%s\n", stack[e.Canary()])

	writer.Flush()

	return builder.String()
}

func (e *ELF) Close() error {
	return e.file.Close()
}
