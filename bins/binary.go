package bins

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
