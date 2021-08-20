package gopwn

import (
	"errors"
	"fmt"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/knightsc/gapstone"
)

type Assembler struct {
	ks *keystone.Keystone
}

func NewAssembler(arch Arch) (*Assembler, error) {
	var ks *keystone.Keystone
	var err error

	switch arch {
	case ARCH_X86_64:
		ks, err = keystone.New(keystone.ARCH_X86, keystone.MODE_64)
	case ARCH_I386:
		ks, err = keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	case ARCH_AARCH64:
		ks, err = keystone.New(keystone.ARCH_ARM, keystone.MODE_64)
	case ARCH_ARM:
		ks, err = keystone.New(keystone.ARCH_ARM, keystone.MODE_32)
	default:
		return nil, errors.New("Unsupported machine type.")
	}

	if err != nil {
		return nil, err
	}

	if err := ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_INTEL); err != nil {
		return nil, err
	}

	return &Assembler{ks: ks}, nil
}

func (a *Assembler) Assemble(assembly string) ([]byte, error) {
	insn, _, ok := a.ks.Assemble(assembly, 0)
	if !ok {
		return nil, errors.New("Could not assemble instruction")
	}
	return insn, nil
}

func (a *Assembler) Close() error {
	return a.ks.Close()
}

func Assemble(assembly string, arch Arch) ([]byte, error) {
	ks, err := NewAssembler(arch)
	if err != nil {
		return nil, err
	}
	defer ks.Close()

	return ks.Assemble(assembly)
}

func Assemble_X86_64(assembly string) ([]byte, error) {
	return Assemble(assembly, ARCH_X86_64)
}

func Assemble_I386(assembly string) ([]byte, error) {
	return Assemble(assembly, ARCH_I386)
}

type Disassembler struct {
	engine gapstone.Engine
}

func NewDisassembler(arch Arch) (*Disassembler, error) {
	archs := map[Arch]int{
		ARCH_X86_64:  gapstone.CS_ARCH_X86,
		ARCH_I386:    gapstone.CS_ARCH_X86,
		ARCH_ARM:     gapstone.CS_ARCH_ARM,
		ARCH_AARCH64: gapstone.CS_ARCH_ARM64,
	}
	modes := map[Arch]int{
		ARCH_X86_64:  gapstone.CS_MODE_64,
		ARCH_I386:    gapstone.CS_MODE_32,
		ARCH_ARM:     gapstone.CS_MODE_ARM,
		ARCH_AARCH64: gapstone.CS_MODE_ARM,
	}

	engine, err := gapstone.New(archs[arch], modes[arch])
	if err != nil {
		return nil, err
	}

	if err := engine.SetOption(gapstone.CS_OPT_SYNTAX, gapstone.CS_OPT_SYNTAX_INTEL); err != nil {
		return nil, err
	}

	return &Disassembler{engine: engine}, nil
}

func (d *Disassembler) Disam(data []byte, vma uint64) (string, error) {
	insns, err := d.engine.Disasm(data, vma, 0)
	if err != nil {
		return "", err
	}
	var output string
	for _, insn := range insns {
		output += fmt.Sprintf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
	}
	return output, nil
}

func (d *Disassembler) Close() error {
	return d.engine.Close()
}

func Disam(data []byte, vma uint64, arch Arch) (string, error) {
	engine, err := NewDisassembler(arch)
	if err != nil {
		return "", err
	}
	defer engine.Close()

	return engine.Disam(data, vma)
}

func Disam_X86_64(data []byte, vma uint64) (string, error) {
	return Disam(data, vma, ARCH_X86_64)
}

func Disam_I386(data []byte, vma uint64) (string, error) {
	return Disam(data, vma, ARCH_I386)
}
