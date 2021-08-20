package gopwn

type ENVOptions struct {
	CyclicOptions
}

type env struct {
	arch          Arch
	endian        Endian
	cyclicOptions CyclicOptions

	assembler    *Assembler
	disassembler *Disassembler
}

func NewEnv(arch Arch, endian Endian, optFns ...func(o *ENVOptions)) (*env, error) {
	options := ENVOptions{}
	for _, fn := range optFns {
		fn(&options)
	}

	assembler, err := NewAssembler(arch)
	if err != nil {
		return nil, err
	}

	disassembler, err := NewDisassembler(arch)
	if err != nil {
		return nil, err
	}

	return &env{
		arch:         arch,
		endian:       endian,
		assembler:    assembler,
		disassembler: disassembler,
	}, nil
}

func NewEnvFromBinary(path string, optFns ...func(o *ENVOptions)) (*env, error) {
	elf, err := NewELF(path)
	if err != nil {
		return nil, err
	}
	defer elf.Close()
	return NewEnv(elf.Architecture(), elf.Endianness(), optFns...)
}

func (e *env) Arch() Arch {
	return e.arch
}

func (e *env) Endianness() Endian {
	return e.endian
}

func (e *env) Close() error {
	if err := e.assembler.Close(); err != nil {
		return err
	}
	return e.disassembler.Close()
}

func (e *env) Cyclic(length int, optFns ...func(o *CyclicOptions)) string {
	var envOpsFuncs []func(o *CyclicOptions)
	envOpsFuncs = append(envOpsFuncs, func(o *CyclicOptions) { o = &e.cyclicOptions }) //nolint
	envOpsFuncs = append(envOpsFuncs, optFns...)
	return Cyclic(length, envOpsFuncs...)
}

func (e *env) CyclicFind(subseq []byte, optFns ...func(o *CyclicOptions)) int {
	var envOpsFuncs []func(o *CyclicOptions)
	envOpsFuncs = append(envOpsFuncs, func(o *CyclicOptions) { o = &e.cyclicOptions }) //nolint
	envOpsFuncs = append(envOpsFuncs, optFns...)
	return CyclicFind(subseq, envOpsFuncs...)
}

func (e *env) Assemble(assembly string) ([]byte, error) {
	return e.assembler.Assemble(assembly)
}

func (e *env) Disam(data []byte, vma uint64) (string, error) {
	return e.disassembler.Disam(data, vma)
}
