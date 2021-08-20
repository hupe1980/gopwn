package gopwn

import (
	"github.com/hupe1980/gopwn/tubes"
)

type Arch int

const (
	ARCH_X86_64 Arch = iota
	ARCH_I386
	ARCH_AARCH64
	ARCH_ARM
)

func (a Arch) String() string {
	archString := map[Arch]string{
		0: "x86_64",
		1: "i386",
		2: "arm",
		3: "arm_64",
	}
	return archString[a]
}

type Endian int

const (
	LITTLE_ENDIAN Endian = iota
	BIG_ENDIAN
)

func (a Endian) String() string {
	endianString := map[Endian]string{
		0: "little-endian (LE)",
		1: "big-endian (BE)",
	}
	return endianString[a]
}

func NewProcess(argv []string, optFns ...func(o *tubes.ProcessOptions)) (*tubes.Process, error) {
	p, err := tubes.NewProcess(argv, optFns...)
	if err != nil {
		return nil, err
	}
	if err := p.Start(); err != nil {
		return nil, err
	}
	return p, nil
}

func NewRemote(network, addr string) (*tubes.Remote, error) {
	return tubes.NewRemote(network, addr)
}

func NewListener(addr string) (*tubes.Listener, error) {
	return tubes.NewListener(addr)
}
