package gopwn

import (
	"bytes"
	"debug/elf"
	"fmt"
	"os"
)

type ELF struct {
	path  string // Path to the file
	file  *os.File
	elf   *elf.File
	bits  int
	bytes int
}

func NewELF(path string) (*ELF, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	ef, err := elf.NewFile(f)
	if err != nil {
		return nil, err
	}

	e := &ELF{path: path, file: f, elf: ef}

	var ident [elf.EI_NIDENT]byte
	if _, err := e.file.Read(ident[:4]); err != nil {
		return nil, err
	}

	if !isElf(ident[:4]) {
		return nil, fmt.Errorf("Bad magic number at %d\n", ident[0:4])
	}

	switch ef.Class {
	case elf.ELFCLASS64:
		e.bits = 64
		e.bytes = e.bits / 8
	case elf.ELFCLASS32:
		e.bits = 32
		e.bytes = e.bits / 8
	}

	return e, nil
}

func (e *ELF) GOT() {

}

func (e *ELF) Close() error {
	return e.file.Close()
}

func isElf(magic []byte) bool {
	return bytes.HasPrefix(magic, []byte{'\x7f', 'E', 'L', 'F'})
}
