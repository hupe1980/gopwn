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
		return nil, fmt.Errorf("Unsupported machine type %x.", f.Machine)
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
		body, _ := s.Data()
		// If the Size is greater than the VirtualSize the difference will
		// be filled with Zeros, so this space is an code cave
		if s.Size > s.VirtualSize {
			body = append(body, bytes.Repeat([]byte("\x00"), int(s.Size-s.VirtualSize))...)
		}
		caves = append(caves, searchCaves(s.Name, body, uint64(s.Offset), uint64(s.VirtualAddress), parsePeCharacteristics(s.Characteristics), caveSize)...)
	}
	return caves
}

func parsePeCharacteristics(c uint32) string {
	return "TODO"
}
