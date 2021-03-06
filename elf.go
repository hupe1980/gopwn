package gopwn

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

type ELF struct {
	file    *elf.File
	hdr     interface{} // elf.Header32 or elf.Header64
	arch    Arch
	endian  Endian
	symbols []elf.Symbol
	fileBytes
}

func NewELF(path string) (*ELF, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewELFFromReader(fh)
}

func NewELFFromBytes(b []byte) (*ELF, error) {
	r := bytes.NewReader(b)
	return NewELFFromReader(r)
}

func NewELFFromReader(r BinaryReader) (*ELF, error) {
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}

	symbols, _ := f.Symbols()

	var arch Arch
	switch f.Machine {
	case elf.EM_386:
		arch = ARCH_I386
	case elf.EM_X86_64:
		arch = ARCH_AMD64
	case elf.EM_ARM:
		arch = ARCH_ARM
	case elf.EM_AARCH64:
		arch = ARCH_AARCH64
	default:
		return nil, fmt.Errorf("unsupported machine type %x", f.Machine)
	}

	var endian Endian
	switch f.Data {
	case elf.ELFDATA2LSB:
		endian = LITTLE_ENDIAN
	case elf.ELFDATA2MSB:
		endian = BIG_ENDIAN
	default:
		return nil, fmt.Errorf("unknown endianness %x", f.Data)
	}

	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	var hdr interface{}
	switch f.Class {
	case elf.ELFCLASS32:
		hdr = new(elf.Header32)
		if err := binary.Read(r, f.ByteOrder, hdr); err != nil {
			return nil, err
		}
	case elf.ELFCLASS64:
		hdr = new(elf.Header64)
		if err := binary.Read(r, f.ByteOrder, hdr); err != nil {
			return nil, err
		}
	}

	rawData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return &ELF{
		file:    f,
		hdr:     hdr,
		arch:    arch,
		endian:  endian,
		symbols: symbols,
		fileBytes: fileBytes{
			raw: rawData,
			addrToOffset: func(addr uint64) (uint64, error) {
				for _, p := range f.Progs {
					start := p.Vaddr
					end := p.Vaddr + p.Filesz

					if addr >= start && addr < end {
						return addr - p.Vaddr + p.Off, nil
					}
				}
				return 0, fmt.Errorf("address %x is not in range of an ELF segment", addr)
			},
		},
	}, nil
}

// Address determines the virtual address for the specified file offset
func (e *ELF) Address(offset uint64) (uint64, error) {
	for _, p := range e.file.Progs {
		start := p.Off
		end := p.Off + p.Filesz

		if offset >= start && offset < end {
			return offset - p.Off + p.Vaddr, nil
		}
	}
	return 0, fmt.Errorf("Offset %x is not in range of an ELF segment", offset)
}

// Offset determines the offset for the specified virtual address
func (e *ELF) Offset(addr uint64) (uint64, error) {
	return e.addrToOffset(addr)
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
		if prog.Type == elf.PT_GNU_STACK {
			if prog.Flags&elf.PF_X == 0 {
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

func (e *ELF) DumpHeader() {
	fmt.Println("-------------------------- Elf Header ------------------------")
	switch e.file.Class {
	case elf.ELFCLASS64:
		h := e.hdr.(*elf.Header64)
		fmt.Printf("Magic: % x\n", h.Ident)
		fmt.Printf("Class: %s\n", elf.Class(h.Ident[elf.EI_CLASS]))
		fmt.Printf("Data: %s\n", elf.Data(h.Ident[elf.EI_DATA]))
		fmt.Printf("Version: %s\n", elf.Version(h.Version))
		fmt.Printf("OS/ABI: %s\n", elf.OSABI(h.Ident[elf.EI_OSABI]))
		fmt.Printf("ABI Version: %d\n", h.Ident[elf.EI_ABIVERSION])
		fmt.Printf("Elf Type: %s\n", elf.Type(h.Type))
		fmt.Printf("Machine: %s\n", elf.Machine(h.Machine))
		fmt.Printf("Entry: 0x%x\n", h.Entry)
		fmt.Printf("Program Header Offset: 0x%x\n", h.Phoff)
		fmt.Printf("Section Header Offset: 0x%x\n", h.Shoff)
		fmt.Printf("Flags: 0x%x\n", h.Flags)
		fmt.Printf("Elf Header Size (bytes): %d\n", h.Ehsize)
		fmt.Printf("Program Header Entry Size (bytes): %d\n", h.Phentsize)
		fmt.Printf("Number of Program Header Entries: %d\n", h.Phnum)
		fmt.Printf("Size of Section Header Entry: %d\n", h.Shentsize)
		fmt.Printf("Number of Section Header Entries: %d\n", h.Shnum)
		fmt.Printf("Index of Section Header string table: %d\n", h.Shstrndx)
	case elf.ELFCLASS32:
		h := e.hdr.(*elf.Header32)
		fmt.Printf("Magic: % x\n", h.Ident)
		fmt.Printf("Class: %s\n", elf.Class(h.Ident[elf.EI_CLASS]))
		fmt.Printf("Data: %s\n", elf.Data(h.Ident[elf.EI_DATA]))
		fmt.Printf("Version: %s\n", elf.Version(h.Version))
		fmt.Printf("OS/ABI: %s\n", elf.OSABI(h.Ident[elf.EI_OSABI]))
		fmt.Printf("ABI Version: %d\n", h.Ident[elf.EI_ABIVERSION])
		fmt.Printf("Elf Type: %s\n", elf.Type(h.Type))
		fmt.Printf("Machine: %s\n", elf.Machine(h.Machine))
		fmt.Printf("Entry: 0x%x\n", h.Entry)
		fmt.Printf("Program Header Offset: 0x%x\n", h.Phoff)
		fmt.Printf("Section Header Offset: 0x%x\n", h.Shoff)
		fmt.Printf("Flags: 0x%x\n", h.Flags)
		fmt.Printf("Elf Header Size (bytes): %d\n", h.Ehsize)
		fmt.Printf("Program Header Entry Size (bytes): %d\n", h.Phentsize)
		fmt.Printf("Number of Program Header Entries: %d\n", h.Phnum)
		fmt.Printf("Size of Section Header Entry: %d\n", h.Shentsize)
		fmt.Printf("Number of Section Header Entries: %d\n", h.Shnum)
		fmt.Printf("Index of Section Header string table: %d\n", h.Shstrndx)
	}
}

func (e *ELF) Caves(caveSize int) []Cave {
	var caves []Cave
	for _, s := range e.file.Sections {
		data, _ := s.Data()
		caves = append(caves, searchCaves(s.Name, data, s.Offset, s.Addr, s.Size, s.Flags.String(), caveSize)...)
	}
	return caves
}

func (e *ELF) Strings(optFns ...func(o *StringsOptions)) []string {
	options := StringsOptions{}
	for _, fn := range optFns {
		fn(&options)
	}

	var sections []dataReader
	if len(options.Sections) > 0 {
		for _, name := range options.Sections {
			sections = append(sections, e.file.Section(name))
		}
	} else {
		for _, s := range e.file.Sections {
			sections = append(sections, s)
		}
	}
	return parseStrings(sections)
}
