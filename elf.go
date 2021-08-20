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
	arch    Arch
	endian  Endian
	symbols []elf.Symbol
}

func NewELF(path string) (*ELF, error) {
	f, err := elf.Open(path)
	if err != nil {
		return nil, err
	}

	symbols, _ := f.Symbols()

	var arch Arch
	switch f.Machine {
	case elf.EM_386:
		arch = ARCH_I386
	case elf.EM_X86_64:
		arch = ARCH_X86_64
	case elf.EM_ARM:
		arch = ARCH_ARM
	case elf.EM_AARCH64:
		arch = ARCH_AARCH64
	default:
		return nil, fmt.Errorf("Unsupported machine type %x.", f.Machine)
	}

	var endian Endian
	switch f.Data {
	case elf.ELFDATA2LSB:
		endian = LITTLE_ENDIAN
	case elf.ELFDATA2MSB:
		endian = BIG_ENDIAN
	default:
		return nil, fmt.Errorf("Unknown endianness %x.", f.Data)
	}

	return &ELF{
		path:    path,
		file:    f,
		arch:    arch,
		endian:  endian,
		symbols: symbols,
	}, nil
}

func (e *ELF) Architecture() Arch {
	return e.arch
}

func (e *ELF) Endianness() Endian {
	return e.endian
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

// PIE checks whether the current binary is position-independent
func (e *ELF) PIE() bool {
	return e.file.Type == elf.ET_DYN
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
	pie := map[bool]string{
		true:  color.GreenString("PIE enabled"),
		false: color.RedString("No PIE"),
	}

	var builder strings.Builder
	writer := tabwriter.NewWriter(&builder, 0, 0, 3, ' ', 0)

	fmt.Fprintf(writer, "NX:\t%s\n", nx[e.NX()])
	fmt.Fprintf(writer, "Stack:\t%s\n", stack[e.Canary()])
	fmt.Fprintf(writer, "PIE:\t%s\n", pie[e.PIE()])

	symbols := color.GreenString("No Symbols")
	sl := len(e.symbols)
	if sl > 0 {
		symbols = color.RedString(fmt.Sprintf("%d Symbols", sl))
	}
	fmt.Fprintf(writer, "Symbols:\t%s\n", symbols)

	writer.Flush()

	return builder.String()
}

func (e *ELF) Close() error {
	return e.file.Close()
}
