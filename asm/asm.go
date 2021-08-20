package asm

import (
	"errors"

	"github.com/hupe1980/gopwn/bins"
	"github.com/keystone-engine/keystone/bindings/go/keystone"
)

type ASM struct {
	ks *keystone.Keystone
}

func NewASM(arch bins.Arch) (*ASM, error) {
	var ks *keystone.Keystone
	var err error

	switch arch {
	case bins.ARCH_X86_64:
		ks, err = keystone.New(keystone.ARCH_X86, keystone.MODE_64)
	case bins.ARCH_I386:
		ks, err = keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	case bins.ARCH_AARCH64:
		ks, err = keystone.New(keystone.ARCH_ARM, keystone.MODE_64)
	case bins.ARCH_ARM:
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

	return &ASM{ks: ks}, nil
}

func (a *ASM) Assemble(assembly string) ([]byte, error) {
	insn, _, ok := a.ks.Assemble(assembly, 0)
	if !ok {
		return nil, errors.New("Could not assemble instruction")
	}
	return insn, nil
}

func (a *ASM) Close() error {
	return a.ks.Close()
}
