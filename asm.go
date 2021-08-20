package gopwn

import (
	"github.com/hupe1980/gopwn/asm"
	"github.com/hupe1980/gopwn/bins"
)

func ASM(assembly string, arch bins.Arch) ([]byte, error) {
	ks, err := asm.NewASM(arch)
	if err != nil {
		return nil, err
	}
	defer ks.Close()

	return ks.Assemble(assembly)
}

func ASM_X86_64(assembly string) ([]byte, error) {
	return ASM(assembly, bins.ARCH_X86_64)
}

func ASM_I386(assembly string) ([]byte, error) {
	return ASM(assembly, bins.ARCH_I386)
}

func DISASM(data []byte, vma uint64, arch bins.Arch) (string, error) {
	engine, err := asm.NewDISASM(arch)
	if err != nil {
		return "", err
	}
	defer engine.Close()

	return engine.Disam(data, vma)
}

func DISASM_X86_64(data []byte, vma uint64) (string, error) {
	return DISASM(data, vma, bins.ARCH_X86_64)
}

func DISASM_I386(data []byte, vma uint64) (string, error) {
	return DISASM(data, vma, bins.ARCH_I386)
}
