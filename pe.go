package gopwn

import (
	"bytes"
	"debug/pe"
	"fmt"
	"os"
)

type PE struct {
	file *pe.File
	arch Arch
}

func NewPE(path string) (*PE, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewPEFromReader(fh)
}

func NewPEFromBytes(b []byte) (*PE, error) {
	r := bytes.NewReader(b)
	return NewPEFromReader(r)
}

func NewPEFromReader(r BinaryReader) (*PE, error) {
	f, err := pe.NewFile(r)
	if err != nil {
		return nil, err
	}

	var arch Arch
	switch f.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		arch = ARCH_I386
	case pe.IMAGE_FILE_MACHINE_AMD64:
		arch = ARCH_AMD64
	case pe.IMAGE_FILE_MACHINE_ARM:
		arch = ARCH_ARM
	case pe.IMAGE_FILE_MACHINE_ARM64:
		arch = ARCH_AARCH64
	default:
		return nil, fmt.Errorf("unsupported machine type %x", f.Machine)
	}

	return &PE{
		file: f,
		arch: arch,
	}, nil
}

func (p *PE) Close() error {
	return p.file.Close()
}

func (p *PE) Caves(caveSize int) []Cave {
	var caves []Cave
	for _, s := range p.file.Sections {
		data, _ := s.Data()
		// If the Size is greater than the VirtualSize the difference will
		// be filled with Zeros, so this space is an code cave
		if s.Size > s.VirtualSize {
			data = append(data, bytes.Repeat([]byte("\x00"), int(s.Size-s.VirtualSize))...)
		}
		caves = append(caves, searchCaves(s.Name, data, uint64(s.Offset), uint64(s.VirtualAddress), uint64(s.Size), parsePeCharacteristics(s.Characteristics), caveSize)...)
	}
	return caves
}

func parsePeCharacteristics(c uint32) string {
	return "TODO"
}

func (p *PE) Strings(optFns ...func(o *StringsOptions)) []string {
	options := StringsOptions{}
	for _, fn := range optFns {
		fn(&options)
	}

	var sections []dataReader
	if len(options.Sections) > 0 {
		for _, name := range options.Sections {
			sections = append(sections, p.file.Section(name))
		}
	} else {
		for _, s := range p.file.Sections {
			sections = append(sections, s)
		}
	}
	return parseStrings(sections)
}
