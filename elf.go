package gopwn

import (
	"bytes"
	"debug/elf"
	"errors"
	"os"
)

type ELF struct {
	path  string // Path to the file
	file  *os.File
	ident [elf.EI_NIDENT]byte
	hdr   interface{}
	bits  int
	bytes int
}

func NewELF(path string) (*ELF, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	e := &ELF{file: file}

	if _, err := e.file.Read(e.ident[:4]); err != nil {
		return nil, err
	}

	if isElf(e.ident[:4]) == false {
		return nil, errors.New("Not an ELF binary")
	}

	switch elf.Class(e.ident[elf.EI_CLASS]) {
	case elf.ELFCLASS64:
		e.hdr = new(elf.Header64)
		e.bits = 64
		e.bytes = e.bits / 8
	case elf.ELFCLASS32:
		e.hdr = new(elf.Header32)
		e.bits = 32
		e.bytes = e.bits / 8
	default:
		return nil, errors.New("Invalid ELF Arch Class")
	}

	return e, nil
}

func (e *ELF) Close() error {
	return e.file.Close()
}

func isElf(magic []byte) bool {
	return bytes.HasPrefix(magic, []byte{'\x7f', 'E', 'L', 'F'})
}
