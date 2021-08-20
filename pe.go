package gopwn

import (
	"debug/pe"
	"fmt"
)

type PE struct {
	path string // Path to the file
	file *pe.File
	arch Arch
}

func NewPE(path string) (*PE, error) {
	f, err := pe.Open(path)
	if err != nil {
		return nil, err
	}

	var arch Arch
	switch f.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		arch = ARCH_I386
	case pe.IMAGE_FILE_MACHINE_AMD64:
		arch = ARCH_X86_64
	case pe.IMAGE_FILE_MACHINE_ARM:
		arch = ARCH_ARM
	case pe.IMAGE_FILE_MACHINE_ARM64:
		arch = ARCH_AARCH64
	default:
		return nil, fmt.Errorf("Unsupported machine type %x.", f.Machine)
	}

	return &PE{
		path: path,
		file: f,
		arch: arch,
	}, nil
}

func (p *PE) Close() error {
	return p.file.Close()
}
