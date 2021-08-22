package gopwn

import "github.com/hupe1980/gopwn/tube"

type Arch int

const (
	ARCH_AMD64 Arch = iota
	ARCH_I386
	ARCH_AARCH64
	ARCH_ARM
)

func (a Arch) String() string {
	archString := map[Arch]string{
		0: "amd64",
		1: "i386",
		2: "arm",
		3: "aarch64",
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

type Bintype int

const (
	BINTYPE_UNKNOWN Bintype = iota
	BINTYPE_ELF
	BINTYPE_PE
	BINTYPE_MACHO
)

func (b Bintype) String() string {
	bintypeString := map[Bintype]string{
		0: "Unknown",
		1: "ELF",
		2: "PE",
		3: "MACH-O",
	}
	return bintypeString[b]
}

func NewProcess(argv []string, optFns ...func(o *tube.ProcessOptions)) (*tube.Process, error) {
	return tube.NewProcess(argv, optFns...)
}

func NewRemote(network, addr string) (*tube.Remote, error) {
	return tube.NewRemote(network, addr)
}

func NewListener(addr string) (*tube.Listener, error) {
	return tube.NewListener(addr)
}
