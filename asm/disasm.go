package asm

import (
	"fmt"

	"github.com/hupe1980/gopwn/bins"
	"github.com/knightsc/gapstone"
)

type DISASM struct {
	engine gapstone.Engine
}

func NewDISASM(arch bins.Arch) (*DISASM, error) {
	archs := map[bins.Arch]int{
		bins.ARCH_X86_64:  gapstone.CS_ARCH_X86,
		bins.ARCH_I386:    gapstone.CS_ARCH_X86,
		bins.ARCH_ARM:     gapstone.CS_ARCH_ARM,
		bins.ARCH_AARCH64: gapstone.CS_ARCH_ARM64,
	}
	modes := map[bins.Arch]int{
		bins.ARCH_X86_64:  gapstone.CS_MODE_64,
		bins.ARCH_I386:    gapstone.CS_MODE_32,
		bins.ARCH_ARM:     gapstone.CS_MODE_ARM,
		bins.ARCH_AARCH64: gapstone.CS_MODE_ARM,
	}

	engine, err := gapstone.New(archs[arch], modes[arch])
	if err != nil {
		return nil, err
	}

	if err := engine.SetOption(gapstone.CS_OPT_SYNTAX, gapstone.CS_OPT_SYNTAX_INTEL); err != nil {
		return nil, err
	}

	return &DISASM{engine: engine}, nil
}

func (d *DISASM) Disam(data []byte, vma uint64) (string, error) {
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

func (d *DISASM) Close() error {
	return d.engine.Close()
}
